use axum::Router;
use axum_extra::routing::RouterExt;
use create_realm::create_realm;

use crate::domain::realm::ports::RealmService;

pub mod create_realm;

pub fn realm_routes<R: RealmService>() -> Router {
    Router::new().typed_post(create_realm::<R>)
}
