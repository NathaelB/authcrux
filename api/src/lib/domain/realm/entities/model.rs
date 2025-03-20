use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use sqlx::prelude::FromRow;
use uuid::{NoContext, Timestamp, Uuid};

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize, Ord, PartialOrd, FromRow)]
pub struct Realm {
    pub id: Uuid,
    pub name: String,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
}

impl Realm {
    pub fn new(name: String) -> Self {
        let now = Utc::now();
        let seconds = match now.timestamp().try_into() {
            Ok(s) => s,
            Err(_) => 0, // Fallback to 0 if conversion fails
        };

        let timestamp = Timestamp::from_unix(NoContext, seconds, 0);
        Self {
            id: Uuid::new_v7(timestamp),
            name,
            created_at: now,
            updated_at: now,
        }
    }
}
