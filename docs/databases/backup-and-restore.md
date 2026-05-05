# Database Backup and Restore

We have two scripts for managing database backups: `scripts/backup.sh` and `scripts/restore.sh`.

All 6 databases are backed up: `collection-db`, `device-db`, `auth-db`, `context-db`, `erp-db`, and `chirpstack-db`.

## Backup

Run the backup script from the project root:

```bash
./scripts/backup.sh
```

This dumps all 6 databases into `backups/` as compressed files, one per database:

```
backups/
  collection-db_2026-05-04_15-08-12.dump.gz
  device-db_2026-05-04_15-08-12.dump.gz
  auth-db_2026-05-04_15-08-12.dump.gz
  ...
```

- Backups older than 7 days are deleted automatically
- All services can stay running during backup — `pg_dump` is non-disruptive
- `collection-db` (TimescaleDB) is backed up with `--no-tablespaces` as required by TimescaleDB

### Scheduling (production)

Add a cron job to run the backup daily at 2am:

```bash
crontab -e
```

```
0 2 * * * /path/to/scripts/backup.sh >> /path/to/backups/backup.log 2>&1
```

### Credentials

The script reads credentials from `.env` in the project root (production). In development, it falls back to the hardcoded defaults in `docker-compose.yml`.

## Restore

To restore a single database:

```bash
./scripts/restore.sh <db-service> <backup-file>
```

Example:

```bash
./scripts/restore.sh device-db backups/device-db_2026-05-04_15-08-12.dump.gz
```

The script handles everything automatically:
1. Stops the dependent application service
2. Restores the data
3. Restarts the service

For `collection-db` specifically (TimescaleDB), the script also tears down and recreates the container and volume before restoring, since TimescaleDB requires a clean empty database to restore into.

### Available services

| Service | Contains |
|---|---|
| `collection-db` | Sensor measurements (TimescaleDB hypertable) |
| `device-db` | Sensors, gateways, measurement types |
| `auth-db` | Companies, factories, factory areas, users |
| `context-db` | Aggregation rules and context data |
| `erp-db` | Production resources, orders, operations |
| `chirpstack-db` | LoRaWAN network server data |

### Restoring multiple databases

Run the script once per database you need to restore:

```bash
./scripts/restore.sh device-db backups/device-db_2026-05-04_15-08-12.dump.gz
./scripts/restore.sh auth-db   backups/auth-db_2026-05-04_15-08-12.dump.gz
```

## Notes

- `backups/` is excluded from git and Tilt file watching
- Backups are stored locally — for production, copy them off-site (S3, external drive, etc.) so they survive a server failure
- `context-db` holds computed/aggregated data and can be skipped if you are comfortable letting the context service recalculate it
