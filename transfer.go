package nusagate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func (nusagate *Impl) CreateTransfer(request *ReqCreateTransfer) (*ResCreateTransfer, *Error) {

	resp := &ResCreateTransfer{}
	var url = fmt.Sprintf("%s/v1/merchant-tranfers/", nusagate.Config.BaseUrl)
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

func (nusagate *Impl) GetTransferById(withdrawalId string) (*ResGetTransferDetail, *Error) {

	resp := &ResGetTransferDetail{}
	var url = fmt.Sprintf("%s/v1/merchant-tranfers/%s", nusagate.Config.BaseUrl, withdrawalId)

	errRequest := nusagate.HttpRequest.Call(http.MethodGet, url, nusagate.Config.ApiKey, nusagate.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (nusagate *Impl) GetTransfers(query *ReqGetTransferList) (*ResGetTransferList, *Error) {

	resp := &ResGetTransferList{}
	var url = fmt.Sprintf("%s/v1/merchant-tranfers?page=%s&perPage=%s&fromDate=%s&toDate=%s&status=%s", nusagate.Config.BaseUrl, query.Page, query.PerPage, query.FromDate, query.ToDate, query.Status)

	errRequest := nusagate.HttpRequest.Call(http.MethodGet, url, nusagate.Config.ApiKey, nusagate.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (nusagate *Impl) CalculateTransfer(request *ReqTransfer) (*ResCalculateTransfer, *Error) {

	resp := &ResCalculateTransfer{}
	var url = fmt.Sprintf("%s/v1/merchant-tranfers/calculate", nusagate.Config.BaseUrl)
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
