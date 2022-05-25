package nusagate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)


var validate *validator.Validate


type Nusagate interface {
	CreateInvoice(request *ReqCreateInvoice) (*ResCreateInvoice, *Error)
	GetInvoiceById(invoiceId string) (*ResGetInvoiceDetail, *Error)
	GetInvoices(query *ReqGetInvoiceList) (*ResGetInvoiceList, *Error)
	VoidInvoice(invoiceId string) (*ResVoidInvoice, *Error)

	CreateWithdrawal(request *ReqWithdrawal) (*ResCreateWithdrawal, *Error)
	GetWithdrawalById(withdrawalId string) (*ResGetWithdrawalDetail, *Error)
	GetWithdrawals(query *ReqGetWithdrawalList) (*ResGetWithdrawalList, *Error)
	CalculateWithdrawal(request *ReqWithdrawal) (*ResCalculateWithdrawal, *Error)
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
