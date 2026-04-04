package report

import "github.com/mhbhuiyan99/Finance-Dashboard-System/domain"

type service struct {
	rcdRepo RecordRepo
}

func NewService(rcdRepo RecordRepo) Service {
	return &service{
		rcdRepo: rcdRepo,
	}
}

func (svc *service) Create(r domain.FinancialRecord) (*domain.FinancialRecord, error) {
	return svc.rcdRepo.Create(r)
}

func (svc *service) GetByID(id string) (*domain.FinancialRecord, error) {
	return svc.rcdRepo.GetByID(id)
}

func (svc *service) List() ([]domain.FinancialRecord, error) {
	return svc.rcdRepo.List()
}

func (svc *service) Delete(id string) error {
	return svc.rcdRepo.Delete(id)
}

func (svc *service) Update(r domain.FinancialRecord) (*domain.FinancialRecord, error) {
	return svc.rcdRepo.Update(r)
}

func (svc *service) Transactions(TID string, filter domain.RecordFilter) ([]domain.FinancialRecord, error) {
	return svc.rcdRepo.Transactions(TID, filter)
}