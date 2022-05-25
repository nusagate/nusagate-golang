package nusagate

import "time"

type ReqCreateInvoice struct {
	ExternalId  string    `json:"externalId"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	DueDate     time.Time `json:"dueDate"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
}

type ReqGetInvoiceList struct {
	Page     string `json:"page"`
	PerPage  string `json:"perPage"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
	OrderBy  string `json:"orderBy"`
	SortBy   string `json:"sortBy"`
	Status   string `json:"status"`
	Search   string `json:"search"`
}

type ReqWithdrawal struct {
	Address      string  `json:"address"`
	Amount       float64 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
}


type ReqGetWithdrawalList struct {
	Page     string `json:"page"`
	PerPage  string `json:"perPage"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
	Status   string `json:"status"`
}
