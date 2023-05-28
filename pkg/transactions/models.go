package transaction

import "time"

type TransactionType struct {
	ID   string `json:"id"`
	Type string `json:"type"` // Credit | Debit | Transfer
	Name string `json:"name"`
}

type Category struct {
	ID              string `json:"id"`
	CategoryName    string `json:"category_name"`
	SubcategoryName string `json:"subcategory_name"`
}

type Currency struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Account struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"` // Cash | Bank Account | Investment | Savings | Cripto Account
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

type Transaction struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Type           string    `json:"type"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	Account        string    `json:"account"`
	OriginalAmount float64   `json:"original_amount"`
	Currency       string    `json:"currency"`
	BaseAmount     float64   `json:"base_amount"`
}

type Transfer struct {
	ID                 string    `json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	SourceAccount      string    `json:"source_account"`
	SourceAmount       float64   `json:"source_amount"`
	SourceFee          float64   `json:"source_fee"`
	DestinationAccount string    `json:"destination_account"`
	DestinationAmount  float64   `json:"destination_amount"`
	DestinationFee     float64   `json:"destination_fee"`
}
