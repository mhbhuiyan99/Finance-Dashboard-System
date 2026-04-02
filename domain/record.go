package domain

import "time"

type RecordType string

const (
    RecordIncome  RecordType = "income"
    RecordExpense RecordType = "expense"
)

type FinancialRecord struct {
    ID        string     `db:"id" json:"id"`
    UserID    string     `db:"user_id" json:"user_id"`
    Amount    float64    `db:"amount" json:"amount"`
    Type      RecordType `db:"type" json:"type"`
    Category  string     `db:"category" json:"category"`
    Date      time.Time  `db:"date" json:"date"`
    Notes     string     `db:"notes" json:"notes"`
    CreatedAt time.Time  `db:"created_at" json:"created_at"`
    UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
    DeletedAt *time.Time `db:"deleted_at" json:"-"`
}

type RecordFilter struct {
    Type      string
    Category  string
    DateFrom  string
    DateTo    string
    Page      int
    PageSize  int
}