package controller

import (
	"net/http"

	"github.com/frozentech/logs"
	merchantError "github.com/frozentech/merchant/error"
	"github.com/frozentech/merchant/model"
	"github.com/gorilla/mux"
)

// Merchant ...
func Merchant(w http.ResponseWriter, r *http.Request) {
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
	case http.MethodPut:
		MerchantPut(resp, r)
		return
	case http.MethodDelete:
		MerchantDelete(resp, r)
		return
	case http.MethodOptions:
		Options(resp, r)
		return
	default:
		MethodNotAllowed(resp, r)
		return
	}
}

// MerchantPut replace a merchant record
func MerchantPut(w *Response, r *http.Request) {
	var (
		merchant   = model.NewMerchant()
		request    = &model.Merchant{}
		resp       = model.Response{}
		merchantID = mux.Vars(r)["merchantId"]
		err        error
	)

	if err = RequestBody(r, request, w); err != nil {
		return
	}

	if err = merchant.FindID(merchantID); err != nil {
		resp = merchantError.StatusRecord(merchantError.RecordNotFound)
		Output(w, resp.Error.Status, resp)
		return
	}

	merchant.Name = request.Name

	if err = merchant.Update(); err != nil {
		resp = merchantError.StatusRecord(merchantError.DatabaseOperationFailure)
		Output(w, resp.Error.Status, resp)
		return
	}

	resp = merchantError.StatusRecord(merchantError.NoError)
	resp.SetResult(merchant)
	Output(w, http.StatusOK, resp)
	return
}

// MerchantDelete delete a merchant record
func MerchantDelete(w *Response, r *http.Request) {
	var (
		merchant   = model.NewMerchant()
		resp       = model.Response{}
		merchantID = mux.Vars(r)["merchantId"]
		err        error
	)

	if err = merchant.FindID(merchantID); err != nil {
		resp = merchantError.StatusRecord(merchantError.RecordNotFound)
		Output(w, resp.Error.Status, resp)
		return
	}

	if err = merchant.Delete(); err != nil {
		resp = merchantError.StatusRecord(merchantError.DatabaseOperationFailure)
		Output(w, resp.Error.Status, resp)
		return
	}

	resp = merchantError.StatusRecord(merchantError.NoError)
	Output(w, http.StatusOK, resp)
}
