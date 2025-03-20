use std::sync::Arc;

use axum::{Extension, http::StatusCode};
use axum_macros::TypedPath;
use axum_valid::{garde::Validate, query::Query, ValidQuery};
use serde::{Deserialize, Serialize};

use crate::application::http::{
    errors::ApiError,
    handlers::ApiSuccess,
    validation::ValidJsonExt,
};
use crate::domain::realm::{entities::model::Realm, ports::RealmService};

#[derive(TypedPath)]
#[typed_path("/realms")]
pub struct ListRealmsRoute;

/// Paramètres de pagination et de filtrage pour la liste des realms
#[derive(Debug, Serialize, Deserialize, Validate)]
pub struct ListRealmsParams {
    #[garde(skip_if_none, range(min = 1, message = "La page doit être au moins 1"))]
    pub page: Option<u32>,
    
    #[garde(skip_if_none, range(min = 1, max = 100, message = "La limite doit être entre 1 et 100"))]
    pub limit: Option<u32>,
    
    #[garde(skip)]
    pub search: Option<String>,
}

pub async fn list_realms<R: RealmService>(
    _: ListRealmsRoute,
    Extension(realm_service): Extension<Arc<R>>,
    ValidQuery(params): ValidQuery<ListRealmsParams>,
) -> Result<ApiSuccess<Vec<Realm>>, ApiError> {
    // Utiliser les valeurs par défaut si non fournies
    let page = params.page.unwrap_or(1);
    let limit = params.limit.unwrap_or(20);
    
    // Récupérer les realms avec pagination
    let realms = realm_service
        .list_realms(page, limit, params.search)
        .await
        .map_err(ApiError::from)?;
    
    Ok(ApiSuccess::new(StatusCode::OK, realms))
} 