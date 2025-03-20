use serde::{Deserialize, Serialize};
use validator::Validate;

#[derive(Debug, Serialize, Deserialize, Validate)]
pub struct CreateRealmValidator {
    #[validate(length(min = 1, message = "name is required"))]
    #[serde(default)]
    pub name: String,

    #[validate(length(min = 1, message = "description is required"))]
    #[serde(default)]
    pub description: String,
}

#[derive(Debug, Serialize, Deserialize, Validate)]
pub struct UpdateRealmValidator {
    #[validate(length(min = 1, message = "name is required"))]
    pub name: String,

    pub description: Option<String>,
}
