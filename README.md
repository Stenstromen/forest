# Forest

## Runs MySQL Schema

```sql
CREATE TABLE users (
  id           BINARY(16) NOT NULL PRIMARY KEY,
  email        VARCHAR(255) NOT NULL UNIQUE,
  created_at   DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at   DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)
) ENGINE=InnoDB;
```

```sql
CREATE TABLE runs (
  id                 BINARY(16) NOT NULL PRIMARY KEY,

  user_id            BINARY(16) NOT NULL,

  occurred_at        DATETIME(3) NOT NULL,

  distance_m         INT UNSIGNED NOT NULL,
  duration_s         INT UNSIGNED NOT NULL,
  calories           INT UNSIGNED NULL,

  created_at         DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  updated_at         DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),

  CONSTRAINT fk_runs_user
    FOREIGN KEY (user_id) REFERENCES users(id)
    ON DELETE CASCADE,

  CONSTRAINT chk_runs_distance_positive CHECK (distance_m > 0),
  CONSTRAINT chk_runs_duration_positive CHECK (duration_s > 0)
) ENGINE=InnoDB;

CREATE INDEX idx_runs_user_occurred ON runs(user_id, occurred_at);
CREATE INDEX idx_runs_occurred ON runs(occurred_at);
```
