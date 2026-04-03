-- +migrate Down
DROP INDEX IF EXISTS idx_records_deleted_at;
DROP INDEX IF EXISTS idx_records_date;
DROP INDEX IF EXISTS idx_records_user_id;
DROP TABLE IF EXISTS financial_records;