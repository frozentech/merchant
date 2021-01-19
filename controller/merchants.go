package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/frozentech/logs"
	merchantError "github.com/frozentech/merchant/error"
	"github.com/frozentech/merchant/model"
)

// Merchants ...
func Merchants(w http.ResponseWriter, r *http.Request) {
	resp := NewResponse(w)
	Log = logs.New()
	Log.Print(Strpad("METHOD", 15, " ", "RIGHT"), r.Method)
	Log.Print(Strpad("ENDPOINT", 15, " ", "RIGHT"), r.URL.EscapedPath())
	if r.Method == http.MethodGet {
		Log.Print(Strpad("REQUEST", 15, " ", "RIGHT"), "")
	} else {
		Log.Print(Strpad("REQUEST", 15, " ", "RIGHT"), string(ReadBody(r)))
	}

	defer func() {
		Log.Print(Strpad("STATUS", 15, " ", "RIGHT"), resp.status)
		Log.Print(Strpad("RESPONSE", 15, " ", "RIGHT"), resp.body)
		Log.Dump(true)
	}()

	switch r.Method {
	case http.MethodPost:
		MerchantsPost(resp, r)
		return
	case http.MethodGet:
		MerchantsGet(resp, r)
		return
	case http.MethodOptions:
		Options(resp, r)
		return
	default:
		MethodNotAllowed(resp, r)
		return
	}
}

// MerchantsPost create a merchant record
func MerchantsPost(w *Response, r *http.Request) {
	var (
		merchant = model.NewMerchant()
		resp     = model.Response{}
		err      error
	)

	if err = RequestBody(r, merchant, w); err != nil {
		return
	}

	merchant.ID = model.GenerateUUID()
	merchant.CreateAt = time.Now().Format(model.MYSQLTimestampFormat)
	if err = merchant.Create(); err != nil {
		resp = merchantError.StatusRecord(merchantError.DatabaseOperationFailure)
		Output(w, resp.Error.Status, resp)
		return
	}

	resp = merchantError.StatusRecord(merchantError.NoError)
	resp.SetResult(merchant)
	Output(w, http.StatusCreated, resp)
	return
}

// MerchantsGet list all merchant record
func MerchantsGet(w *Response, r *http.Request) {
	var (
		merchant          = model.NewMerchant()
		merchants         []model.Merchant
		resp              = model.Response{}
		err               error
		page              = r.FormValue("page")
		limit             = r.FormValue("limit")
		intPage, intLimit int
	)

	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "10"
	}

	if intPage, err = strconv.Atoi(page); err != nil {
		resp = merchantError.StatusRecord(merchantError.InvalidPageValue)
		Output(w, resp.Error.Status, resp)
		return
	}

	if intLimit, err = strconv.Atoi(limit); err != nil {
		resp = merchantError.StatusRecord(merchantError.InvalidLimitValue)
		Output(w, resp.Error.Status, resp)
		return
	}

	merchants, _ = merchant.FindAll(intPage, intLimit)
	resp = merchantError.StatusRecord(merchantError.NoError)
	resp.SetResult(
		model.Records{
			Page:   intPage,
			Count:  len(merchants),
			Record: merchants,
		})
	Output(w, http.StatusOK, resp)
	return
}
