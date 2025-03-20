use super::{
    entities::{error::RealmError, model::Realm},
    ports::{RealmRepository, RealmService},
};

#[derive(Debug, Clone)]
pub struct RealmServiceImpl<R>
where
    R: RealmRepository,
{
    pub realm_repository: R,
}

impl<R> RealmServiceImpl<R>
where
    R: RealmRepository,
{
    pub fn new(realm_repository: R) -> Self {
        Self { realm_repository }
    }
}

impl<R> RealmService for RealmServiceImpl<R>
where
    R: RealmRepository,
{
    async fn create_realm(&self, name: String) -> Result<Realm, RealmError> {
        self.realm_repository.create_realm(name).await
    }

    async fn get_by_name(&self, name: String) -> Result<Realm, RealmError> {
        self.realm_repository
            .get_by_name(name)
            .await?
            .ok_or(RealmError::NotFound)
    }
}
