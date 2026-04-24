-- +goose Up
CREATE TABLE "auth"."erp_agent_credential" (
  "company_id" uuid PRIMARY KEY,
  "key_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "secret_hash" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "rotated_at" timestamptz,
  "revoked_at" timestamptz,

  CONSTRAINT fk_erp_agent_credential_company
    FOREIGN KEY (company_id)
    REFERENCES auth.company(company_id)
    ON DELETE CASCADE,

  CONSTRAINT uq_erp_agent_credential_key_id UNIQUE (key_id)
);

CREATE INDEX idx_erp_agent_credential_key_id
  ON auth.erp_agent_credential(key_id);

CREATE INDEX idx_erp_agent_credential_revoked_at
  ON auth.erp_agent_credential(revoked_at);

-- +goose Down
DROP TABLE "auth"."erp_agent_credential";
