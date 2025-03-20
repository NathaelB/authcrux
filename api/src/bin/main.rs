use std::sync::Arc;

use authcrux::{application::http::{HttpServer, HttpServerConfig}, env::{AppEnv, Env}};
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

    let server_config = HttpServerConfig::new(env.port.clone());

    let http_server = HttpServer::new(server_config).await?;
    http_server.run().await?;
    println!("Hello AuthCrux");

    Ok(())
}
