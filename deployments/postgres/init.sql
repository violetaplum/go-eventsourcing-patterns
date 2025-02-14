-- deployments/postgres/init.sql
CREATE TABLE IF NOT EXISTS accounts (
                                        id        VARCHAR(100) PRIMARY KEY,
                                        user_name VARCHAR(255) NOT NULL,
                                        balance BIGINT NOT NULL,
                                        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                        updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS events (
                                      id         VARCHAR(100) PRIMARY KEY,
                                      account_id VARCHAR(100) NOT NULL REFERENCES accounts(id),
                                      event_type VARCHAR(255) NOT NULL,
                                      event_data JSONB NOT NULL,
                                      created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
                             );


CREATE INDEX idx_events_account_id ON events(account_id);

