use thiserror::Error;

#[derive(Debug, Clone, Error)]
pub enum RealmError {
    #[error("Realm not found")]
    NotFound,
    #[error("Realm already exists")]
    AlreadyExists,
    #[error("Invalid realm")]
    Invalid,
}
