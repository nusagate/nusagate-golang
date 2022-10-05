package nusagate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Nusagate interface {
	CreateInvoice(request *ReqCreateInvoice) (*ResCreateInvoice, *Error)
	GetInvoiceById(id string) (*ResGetInvoiceDetail, *Error)
	GetInvoices(query *ReqGetInvoiceList) (*ResGetInvoiceList, *Error)
	VoidInvoice(id string) (*ResVoidInvoice, *Error)

	CreateTransfer(request *ReqCreateTransfer) (*ResCreateTransfer, *Error)
	GetTransferById(id string) (*ResGetTransferDetail, *Error)
	GetTransfers(query *ReqGetTransferList) (*ResGetTransferList, *Error)
	CalculateTransfer(request *ReqTransfer) (*ResCalculateTransfer, *Error)
}

type Impl struct {
	Config      *ConfigOption
	HttpRequest *HttpRequestImpl
}

func New(isProduction bool, apiKey string, secretKey string, config ...ConfigOption) Nusagate {
	defaultConfig := DefaultConfig(isProduction, apiKey, secretKey)

	if len(config) > 0 {
		defaultConfig = &config[0]
		defaultConfig.IsProduction = isProduction
		defaultConfig.ApiKey = apiKey
		defaultConfig.SecretKey = secretKey
	}

	fmt.Println(defaultConfig)

	return &Impl{
		Config: defaultConfig,
	}
}
