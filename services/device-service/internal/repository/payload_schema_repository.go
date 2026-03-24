package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/dbutil"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	upsertDraftPayloadSchemaQuery = `
		INSERT INTO device.payload_schema (chirpstack_profile_id, payload_key)
		SELECT $1, unnest($2::text[])
		ON CONFLICT (chirpstack_profile_id, payload_key) DO NOTHING
	`

	savePayloadSchemaLabelsQuery = `
		UPDATE device.payload_schema AS ps
		SET measurement_type = v.measurement_type,
		    unit              = v.unit
		FROM (
		    SELECT
		        unnest($1::text[]) AS chirpstack_profile_id,
		        unnest($2::text[]) AS payload_key,
		        unnest($3::text[]) AS measurement_type,
		        unnest($4::text[]) AS unit
		) AS v
		WHERE ps.chirpstack_profile_id = v.chirpstack_profile_id
		  AND ps.payload_key           = v.payload_key
	`

	findPayloadSchemaByProfileIDQuery = `
		SELECT id, chirpstack_profile_id, payload_key, measurement_type, unit
		FROM device.payload_schema
		WHERE chirpstack_profile_id = $1
		ORDER BY payload_key ASC
	`

	findDraftProfilesQuery = `
		SELECT DISTINCT chirpstack_profile_id
		FROM device.payload_schema
		WHERE measurement_type IS NULL
		ORDER BY chirpstack_profile_id ASC
	`
)

// PayloadSchemaRepository handles persistence of payload schema rows in the database.
type PayloadSchemaRepository struct {
	db *dbutil.DB
}

// NewPayloadSchemaRepository creates a new PayloadSchemaRepository with the given database.
func NewPayloadSchemaRepository(db *dbutil.DB) *PayloadSchemaRepository {
	return &PayloadSchemaRepository{db: db}
}

// UpsertDrafts inserts payload keys as draft rows (measurement_type = NULL) for a profile.
// Skips keys that already exist so existing labels are never overwritten.
func (r *PayloadSchemaRepository) UpsertDrafts(ctx context.Context, chirpstackProfileID string, payloadKeys []string) error {
	_, err := r.db.Pool.Exec(ctx, upsertDraftPayloadSchemaQuery, chirpstackProfileID, payloadKeys)
	if err != nil {
		return fmt.Errorf("upsert draft payload schema: %w", err)
	}

	return nil
}

// SaveLabels updates measurement_type and unit for the given payload schema rows.
// All rows are updated in a single query — atomic by default, no transaction needed.
// Returns domain.ErrInvalidMeasurementType if any slug does not exist in the vocabulary.
func (r *PayloadSchemaRepository) SaveLabels(ctx context.Context, schemas []domain.PayloadSchema) error {
	profileIDs := make([]string, len(schemas))
	payloadKeys := make([]string, len(schemas))
	measurementTypes := make([]string, len(schemas))
	units := make([]*string, len(schemas))

	for i, s := range schemas {
		profileIDs[i] = s.ChirpstackProfileID
		payloadKeys[i] = s.PayloadKey
		measurementTypes[i] = *s.MeasurementType
		units[i] = s.Unit
	}

	_, err := r.db.Pool.Exec(ctx, savePayloadSchemaLabelsQuery, profileIDs, payloadKeys, measurementTypes, units)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			slog.Warn("foreign key violation on payload schema save", "constraint", pgErr.ConstraintName, "detail", pgErr.Detail)
			return domain.ErrInvalidMeasurementType
		}
		return fmt.Errorf("save payload schema labels: %w", err)
	}

	return nil
}

// FindByProfileID returns all payload schema rows for a given ChirpStack profile ID,
// ordered by payload key. Includes both draft (measurement_type = nil) and labeled rows.
func (r *PayloadSchemaRepository) FindByProfileID(ctx context.Context, chirpstackProfileID string) ([]domain.PayloadSchema, error) {
	rows, err := r.db.Pool.Query(ctx, findPayloadSchemaByProfileIDQuery, chirpstackProfileID)
	if err != nil {
		return nil, fmt.Errorf("find payload schema by profile id: %w", err)
	}
	defer rows.Close()

	out := []domain.PayloadSchema{}

	for rows.Next() {
		var s domain.PayloadSchema
		if err := rows.Scan(&s.ID, &s.ChirpstackProfileID, &s.PayloadKey, &s.MeasurementType, &s.Unit); err != nil {
			return nil, fmt.Errorf("scan payload schema: %w", err)
		}
		out = append(out, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return out, nil
}

// FindDraftProfiles returns distinct ChirpStack profile IDs that have at least one
// unlabeled row. Used to populate the admin dashboard badge.
func (r *PayloadSchemaRepository) FindDraftProfiles(ctx context.Context) ([]string, error) {
	rows, err := r.db.Pool.Query(ctx, findDraftProfilesQuery)
	if err != nil {
		return nil, fmt.Errorf("find draft profiles: %w", err)
	}
	defer rows.Close()

	out := []string{}

	for rows.Next() {
		var profileID string
		if err := rows.Scan(&profileID); err != nil {
			return nil, fmt.Errorf("scan draft profile: %w", err)
		}
		out = append(out, profileID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return out, nil
}
