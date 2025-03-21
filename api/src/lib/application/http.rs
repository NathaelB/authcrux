use std::sync::Arc;

use anyhow::Context;
use axum::Extension;
use handlers::realm::realm_routes;
use tracing::{info, info_span};

use crate::domain::realm::ports::RealmService;

use super::state::AppState;

pub mod errors;
pub mod handlers;
pub mod responses;
pub mod validation;

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct HttpServerConfig {
    pub port: String,
}

impl HttpServerConfig {
    pub fn new(port: String) -> Self {
        Self { port }
    }
}

pub struct HttpServer {
    router: axum::Router,
    listener: tokio::net::TcpListener,
}

impl HttpServer {
    pub async fn new<R>(
        config: HttpServerConfig,
        realm_service: Arc<R>,
    ) -> Result<Self, anyhow::Error>
    where
        R: RealmService,
    {
        let trace_layer = tower_http::trace::TraceLayer::new_for_http().make_span_with(
            |request: &axum::extract::Request| {
                let uri: String = request.uri().to_string();
                info_span!("http_request", method = ?request.method(), uri)
            },
        );

        let state = AppState::new(realm_service);

        let router = axum::Router::new()
            .merge(realm_routes::<R>())
            .layer(trace_layer)
            .layer(Extension(Arc::clone(&state.realm_service)));

        let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{}", config.port))
            .await
            .with_context(|| format!("Failed to bind to port {}", config.port))?;

        Ok(Self { router, listener })
    }

    pub async fn run(self) -> Result<(), anyhow::Error> {
        info!(
            "Server is running on http://{}",
            self.listener.local_addr()?
        );
        axum::serve(self.listener, self.router)
            .await
            .context("received error while running server")?;

        Ok(())
    }
}
