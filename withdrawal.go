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

func (nusagate *Impl) CreateWithdrawal(request *ReqWithdrawal) (*ResCreateWithdrawal, *Error) {

	resp := &ResCreateWithdrawal{}
	var url = fmt.Sprintf("%s/v1/withdrawals/", nusagate.Config.BaseUrl)
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


func (nusagate *Impl) GetWithdrawalById(withdrawalId string) (*ResGetWithdrawalDetail, *Error) {

	resp := &ResGetWithdrawalDetail{}
	var url = fmt.Sprintf("%s/v1/withdrawals/%s", nusagate.Config.BaseUrl, withdrawalId)

	errRequest := nusagate.HttpRequest.Call(http.MethodGet, url,nusagate.Config.ApiKey, nusagate.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}

func (nusagate *Impl) GetWithdrawals(query *ReqGetWithdrawalList) (*ResGetWithdrawalList, *Error) {

	resp := &ResGetWithdrawalList{}
	var url = fmt.Sprintf("%s/v1/withdrawals?page=%s&perPage=%s&fromDate=%s&toDate=%s&status=%s", nusagate.Config.BaseUrl, query.Page, query.PerPage, query.FromDate, query.ToDate, query.Status)

	errRequest := nusagate.HttpRequest.Call(http.MethodGet, url,nusagate.Config.ApiKey, nusagate.Config.SecretKey, nil, resp)
	if errRequest != nil {
		return resp, errRequest
	}

	return resp, nil
}


func (nusagate *Impl) CalculateWithdrawal(request *ReqWithdrawal) (*ResCalculateWithdrawal, *Error) {

	resp := &ResCalculateWithdrawal{}
	var url = fmt.Sprintf("%s/v1/withdrawals/calculate", nusagate.Config.BaseUrl)
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