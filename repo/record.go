package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/mhbhuiyan99/Finance-Dashboard-System/domain"
	record "github.com/mhbhuiyan99/Finance-Dashboard-System/record"
)

type RecordRepo interface {
	record.RecordRepo
}

type recordRepo struct {
	db *sqlx.DB
}

func NewRecordRepo(db *sqlx.DB) RecordRepo {
	return &recordRepo{
		db: db,
	}
}

func (r *recordRepo) Create(record domain.FinancialRecord) (*domain.FinancialRecord, error) {
	return nil, nil
}

func (r *recordRepo) GetByID(id string) (*domain.FinancialRecord, error) {
	return nil, nil
}

func (r *recordRepo) Update(record domain.FinancialRecord) (*domain.FinancialRecord, error) {
	return nil, nil
}

func (r *recordRepo) Delete(id string) error {
	return nil
}

func (r *recordRepo) List() ([]domain.FinancialRecord, error) {
	return nil, nil
}

func (r *recordRepo) Transactions(TID string, filter domain.RecordFilter) ([]domain.FinancialRecord, error) {
	return nil, nil
}