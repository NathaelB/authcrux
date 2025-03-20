use crate::{application::http::handlers::ApiError, domain::realm::entities::error::RealmError};

impl From<RealmError> for ApiError {
    fn from(value: RealmError) -> Self {
        match value {
            RealmError::NotFound => ApiError::NotFound("Realm not found".to_string()),
            RealmError::AlreadyExists => {
                ApiError::UnProcessableEntity("Realm already exists".to_string())
            }
            RealmError::Invalid => {
                ApiError::UnProcessableEntity("Realm format is invalid".to_string())
            }
            RealmError::InternalServerError => {
                ApiError::InternalServerError("Internal server error".to_string())
            }
        }
    }
}
