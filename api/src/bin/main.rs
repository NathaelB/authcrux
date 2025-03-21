use std::sync::Arc;

use authcrux::{
    application::http::{HttpServer, HttpServerConfig},
    domain::realm::service::RealmServiceImpl,
    env::{AppEnv, Env},
    infrastructure::{
        db::postgres::Postgres, repositories::realm_repository::PostgresRealmRepository,
    },
};
use clap::Parser;

fn init_logger(env: Arc<Env>) {
    match env.env {
        AppEnv::Development => {
            tracing_subscriber::fmt::init();
        }
        AppEnv::Production => {
            tracing_subscriber::fmt()
                .json()
                .with_max_level(tracing::Level::INFO)
                .init();
        }
    }
}

#[tokio::main]
async fn main() -> Result<(), anyhow::Error> {
    dotenv::dotenv().ok();

    let env = Arc::new(Env::parse());
    init_logger(Arc::clone(&env));

    let postgres = Arc::new(Postgres::new(Arc::clone(&env)).await?);

    let realm_repository = PostgresRealmRepository::new(Arc::clone(&postgres));
    let realm_service = Arc::new(RealmServiceImpl::new(realm_repository));

    let server_config = HttpServerConfig::new(env.port.clone());

    let http_server = HttpServer::new(server_config, realm_service).await?;
    http_server.run().await?;
    println!("Hello AuthCrux");

    Ok(())
}
