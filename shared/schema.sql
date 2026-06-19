-- Shared database schema used by both C# (Register) and Go (Login) backends
CREATE TABLE IF NOT EXISTS users (
  id         INTEGER  PRIMARY KEY AUTOINCREMENT,
  username   TEXT     NOT NULL UNIQUE,
  password   TEXT     NOT NULL,   -- bcrypt hash
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
