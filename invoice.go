package nusagate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"reflect"
)

func (nusagate *Impl) CreateInvoice(request *ReqCreateInvoice) (*ResCreateInvoice, *Error) {

	resp := &ResCreateInvoice{}
	var url = fmt.Sprintf("%s/v1/invoices/", nusagate.Config.BaseUrl)
	var errMarshal error
	jsonRequest := []byte("")

	validate = validator.New()
	errValidate := validate.Struct(request)
	if errValidate != nil {
		return resp, ErrorRequestParam(errValidate)
	}

	isParamsNil := reflect.ValueOf(request).Kind() == reflect.Ptr && reflect.ValueOf(request).IsNil()

	if !isParamsNil {
		jsonRequest, errMarshal = json.Marshal(request)
		if errMarshal != nil {
			log.Println(errMarshal)
			return resp, ErrorGo(errMarshal)
		}
	}

	errRequest := nusagate.HttpRequest.Call(http.MethodPost, url, nusagate.Config.ApiKey, nusagate.Config.SecretKey, bytes.NewBuffer(jsonRequest), resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}


func (nusagate *Impl) GetInvoiceById(invoiceId string) (*ResGetInvoiceDetail, *Error) {

	resp := &ResGetInvoiceDetail{}
	var url = fmt.Sprintf("%s/v1/invoices/%s", nusagate.Config.BaseUrl, invoiceId)

	errRequest := nusagate.HttpRequest.Call(http.MethodGet, url,nusagate.Config.ApiKey, nusagate.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (nusagate *Impl) GetInvoices(query *ReqGetInvoiceList) (*ResGetInvoiceList, *Error) {

	resp := &ResGetInvoiceList{}
	var url = fmt.Sprintf("%s/v1/invoices?page=%s&perPage=%s&fromDate=%s&toDate=%s&orderBy=%s&sortBy=%s&status=%s&search=%s", nusagate.Config.BaseUrl, query.Page, query.PerPage, query.FromDate, query.ToDate, query.OrderBy, query.SortBy, query.Status, query.Search)

	errRequest := nusagate.HttpRequest.Call(http.MethodGet, url,nusagate.Config.ApiKey, nusagate.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	fmt.Println(resp)
	return resp, nil
}


func (nusagate *Impl) VoidInvoice(invoiceId string) (*ResVoidInvoice, *Error) {

	resp := &ResVoidInvoice{}
	var url = fmt.Sprintf("%s/v1/invoices/%s/void", nusagate.Config.BaseUrl, invoiceId)

	errRequest := nusagate.HttpRequest.Call(http.MethodPost, url,nusagate.Config.ApiKey, nusagate.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}
