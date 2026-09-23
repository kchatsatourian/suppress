package state

const schema = `
CREATE TABLE IF NOT EXISTS state (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS deduplication (
    link TEXT PRIMARY KEY,
    expires_at INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_deduplication_expires_at
ON deduplication (expires_at);

CREATE TABLE IF NOT EXISTS channels (
    channel INTEGER PRIMARY KEY,
    url TEXT NOT NULL,
    type TEXT
);

CREATE TABLE IF NOT EXISTS subscriptions (
    subscription INTEGER PRIMARY KEY,
    type TEXT NOT NULL,
    url TEXT NOT NULL,
    pattern TEXT
);

CREATE TABLE IF NOT EXISTS subscription_channels (
    subscription INTEGER NOT NULL,
    channel INTEGER NOT NULL,
    PRIMARY KEY (subscription, channel),
    FOREIGN KEY (subscription) REFERENCES subscriptions(subscription),
    FOREIGN KEY (channel) REFERENCES channels(channel)
);
`
