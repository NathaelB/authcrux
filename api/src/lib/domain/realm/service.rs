use super::{entities::{error::RealmError, model::Realm}, ports::{RealmRepository, RealmService}};

#[derive(Debug, Clone)]
pub struct RealmServiceImpl<R> 
where
  R: RealmRepository
{
  pub realm_repository: R
}

impl<R> RealmService for RealmServiceImpl<R>
where
  R: RealmRepository
{
  async fn create_realm(&self, name: String) -> Result<Realm, RealmError> {
    self.realm_repository.create_realm(name).await
  }
}
