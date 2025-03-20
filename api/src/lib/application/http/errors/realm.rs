use crate::domain::realm::entities::error::RealmError;

use super::ApiError;

impl From<RealmError> for ApiError {
    fn from(value: RealmError) -> Self {
        match value {
            RealmError::NotFound => ApiError::NotFound("Realm not found".to_string()),
            RealmError::AlreadyExists => ApiError::validation_error("Realm already exists", "name"),
            RealmError::Invalid => ApiError::validation_error("Realm format is invalid", "name"),
            RealmError::InternalServerError => {
                ApiError::InternalServerError("Internal server error".to_string())
            }
        }
    }
}
