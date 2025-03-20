use std::sync::Arc;

use crate::domain::realm::ports::RealmService;

#[derive(Debug, Clone)]
pub struct AppState<R>
where
    R: RealmService,
{
    pub realm_service: Arc<R>,
}

impl<R> AppState<R>
where
    R: RealmService,
{
    pub fn new(realm_service: Arc<R>) -> Self {
        Self { realm_service }
    }
}
