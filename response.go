package nusagate

import "time"

type ResCreateInvoice struct {
	Data struct {
		Id              string      `json:"id"`
		Slug            string      `json:"slug"`
		Type            string      `json:"type"`
		ExternalId      string      `json:"externalId"`
		Description     string      `json:"description"`
		DueDate         time.Time   `json:"dueDate"`
		ExpiredAmountAt *time.Time  `json:"expiredAmountAt"`
		Email           string      `json:"email"`
		PhoneNumber     string      `json:"phoneNumber"`
		PayCurrency     string      `json:"payCurrency"`
		PaymentLink     string      `json:"paymentLink"`
		BaseCurrency    string      `json:"baseCurrency"`
		Price           string      `json:"price"`
		Amount          float64     `json:"amount"`
		PaymentType     string      `json:"paymentType"`
		Fee             string      `json:"fee"`
		PaidAt          *time.Time  `json:"paidAt"`
		CompletedAt     *time.Time  `json:"completedAt"`
		Details         interface{} `json:"details"`
		CreatedAt       time.Time   `json:"createdAt"`
		UpdatedAt       time.Time   `json:"updatedAt"`
		Status          string      `json:"status"`
	} `json:"data"`
}

type ResGetInvoiceList struct {
	Data []struct {
		Id          string    `json:"id"`
		Slug        string    `json:"slug"`
		Type        string    `json:"type"`
		ExternalId  string    `json:"externalId"`
		DueDate     time.Time `json:"dueDate"`
		Email       string    `json:"email"`
		PayCurrency string    `json:"payCurrency"`
		Price       float64   `json:"price"`
		Amount      float64   `json:"amount"`
		CreatedAt   time.Time `json:"createdAt"`
		Status      string    `json:"status"`
	} `json:"data"`
}

type ResGetInvoiceDetail struct {
	Data struct {
		Id              string      `json:"id"`
		Slug            string      `json:"slug"`
		Type            string      `json:"type"`
		ExternalId      string      `json:"externalId"`
		Description     string      `json:"description"`
		DueDate         time.Time   `json:"dueDate"`
		ExpiredAmountAt interface{} `json:"expiredAmountAt"`
		Email           interface{} `json:"email"`
		PayCurrency     string      `json:"payCurrency"`
		PaymentLink     string      `json:"paymentLink"`
		BaseCurrency    string      `json:"baseCurrency"`
		Amount          float64     `json:"amount"`
		PaymentType     string      `json:"paymentType"`
		Fee             int         `json:"fee"`
		PaidAt          interface{} `json:"paidAt"`
		CompletedAt     interface{} `json:"completedAt"`
		CreatedAt       time.Time   `json:"createdAt"`
		Details         interface{} `json:"details"`
		UpdatedAt       time.Time   `json:"updatedAt"`
		Status          string      `json:"status"`
		PhoneNumber     string      `json:"phoneNumber"`
		Price           int         `json:"price"`
		QrCode          string      `json:"qrCode"`
		PaymentAddress  struct {
			Id             string `json:"id"`
			Address        string `json:"address"`
			BlockchainCode string `json:"blockchainCode"`
			IsStableCoin   bool   `json:"isStableCoin"`
		} `json:"paymentAddress"`
	} `json:"data"`
}

type ResVoidInvoice struct {
	Data struct {
		ReferenceId string `json:"referenceId"`
	} `json:"data"`
}

type ResCreateTransfer struct {
	Data struct {
		ReferenceId    string  `json:"referenceId"`
		FeeTransfer    float64 `json:"feeTransfer"`
		ReceivedAmount float64 `json:"receivedAmount"`
	} `json:"data"`
}

type ResCalculateTransfer struct {
	Data struct {
		FeeTransfer    int     `json:"feeTransfer"`
		ReceivedAmount float64 `json:"receivedAmount"`
	} `json:"data"`
}

type ResGetTransferList struct {
	Data []struct {
		Id             string     `json:"id"`
		Slug           string     `json:"slug"`
		ExternalID     string     `json:"extenalId"`
		Status         string     `json:"status"`
		Amount         float64    `json:"amount"`
		Fee            float64    `json:"fee"`
		ReceivedAmount float64    `json:"receivedAmount"`
		ConfirmedAt    *time.Time `json:"confirmedAt"`
		CurrencyCode   string     `json:"currencyCode"`
		PaymentType    string     `json:"paymentType"`
		CreatedAt      time.Time  `json:"createdAt"`
		UpdatedAt      time.Time  `json:"updatedAt"`
	} `json:"data"`
}

type ResGetTransferDetail struct {
	Data struct {
		Id             string    `json:"id"`
		Slug           string    `json:"slug"`
		ExternalID     string    `json:"extenalId"`
		Status         string    `json:"status"`
		Amount         float64   `json:"amount"`
		Fee            float64   `json:"fee"`
		ReceivedAmount float64   `json:"receivedAmount"`
		ConfirmedAt    time.Time `json:"confirmedAt"`
		CurrencyCode   string    `json:"currencyCode"`
		PaymentType    string    `json:"paymentType"`
		Transaction    struct {
			Id             string      `json:"id"`
			BlockchainCode string      `json:"blockchainCode"`
			Type           string      `json:"type"`
			TxId           string      `json:"txId"`
			PayCurrency    string      `json:"payCurrency"`
			Reference      string      `json:"reference"`
			Amount         string      `json:"amount"`
			NetAmount      string      `json:"netAmount"`
			Status         string      `json:"status"`
			FromAddress    []string    `json:"fromAddress"`
			ToAddress      string      `json:"toAddress"`
			Details        interface{} `json:"details"`
			ConfirmedAt    *time.Time  `json:"confirmedAt"`
			CreatedAt      time.Time   `json:"createdAt"`
			UpdatedAt      time.Time   `json:"updatedAt"`
		} `json:"transaction"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	} `json:"data"`
}
