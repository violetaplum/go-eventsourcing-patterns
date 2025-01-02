-- init.sql
CREATE TABLE events (
                        id BIGSERIAL PRIMARY KEY,
                        aggregate_id VARCHAR(36) NOT NULL,
                        event_type VARCHAR(50) NOT NULL,
                        event_data JSONB NOT NULL,
                        version INT NOT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

                        UNIQUE(aggregate_id, version)
);

CREATE INDEX idx_events_aggregate_id ON events(aggregate_id);
CREATE INDEX idx_events_created_at ON events(created_at);
CREATE INDEX idx_events_type ON events(event_type);

CREATE TABLE account_view (
                              id VARCHAR(36) PRIMARY KEY,
                              balance BIGINT NOT NULL,
                              version INT NOT NULL,
                              updated_at TIMESTAMP NOT NULL
);