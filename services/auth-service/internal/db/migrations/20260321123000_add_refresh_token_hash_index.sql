-- +goose Up
CREATE INDEX idx_refresh_token_hash
  ON auth.refresh_token(token_hash);

-- +goose Down
DROP INDEX IF EXISTS auth.idx_refresh_token_hash;
