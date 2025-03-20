use std::sync::Arc;

use axum::{Extension, Json, http::StatusCode};
use axum_macros::TypedPath;
use serde::Deserialize;

use crate::{
    application::http::handlers::{ApiError, ApiSuccess},
    domain::realm::{entities::model::Realm, ports::RealmService},
};

#[derive(Deserialize, TypedPath)]
#[typed_path("/realms")]
pub struct CreateRealmRoute;

#[derive(Debug, Deserialize)]
pub struct CreateRealmRequest {
    pub name: String,
}

pub async fn create_realm<R: RealmService>(
    _: CreateRealmRoute,
    Extension(realm_service): Extension<Arc<R>>,
    Json(payload): Json<CreateRealmRequest>,
) -> Result<ApiSuccess<Realm>, ApiError> {
    realm_service
        .create_realm(payload.name)
        .await
        .map_err(ApiError::from)
        .map(|realm| ApiSuccess::new(StatusCode::CREATED, realm))
}
