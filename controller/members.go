package controller

import (
	"net/http"
	"strconv"

	"github.com/frozentech/logs"
	merchantError "github.com/frozentech/merchant/error"
	"github.com/frozentech/merchant/model"
	"github.com/gorilla/mux"
)

// Members ...
func Members(w http.ResponseWriter, r *http.Request) {
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
		MembersPost(resp, r)
		return
	case http.MethodGet:
		MembersGet(resp, r)
		return
	case http.MethodOptions:
		Options(resp, r)
		return
	default:
		MethodNotAllowed(resp, r)
		return
	}
}

// MembersPost create a Member record
func MembersPost(w *Response, r *http.Request) {
	var (
		member     = model.NewMember()
		request    = &model.Member{}
		merchant   = model.NewMerchant()
		merchantID = mux.Vars(r)["merchantId"]
		resp       = model.Response{}
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

	member.MerchantID = merchant.ID
	member.Email = request.Email
	member.Name = request.Name
	if err = member.Create(); err != nil {
		resp = merchantError.StatusRecord(merchantError.DuplicateEmailAddress)
		Output(w, resp.Error.Status, resp)
		return
	}

	resp = merchantError.StatusRecord(merchantError.NoError)
	resp.SetResult(member)
	Output(w, http.StatusCreated, resp)
	return
}

// MembersGet list all Member record
func MembersGet(w *Response, r *http.Request) {
	var (
		member            = model.NewMember()
		merchant          = model.NewMerchant()
		merchantID        = mux.Vars(r)["merchantId"]
		members           []model.Member
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

	if err = merchant.FindID(merchantID); err != nil {
		resp = merchantError.StatusRecord(merchantError.RecordNotFound)
		Output(w, resp.Error.Status, resp)
		return
	}

	members, _ = member.FindByMechantID(merchantID, intPage, intLimit)
	resp = merchantError.StatusRecord(merchantError.NoError)
	resp.SetResult(
		model.Records{
			Page:   intPage,
			Count:  len(members),
			Record: members,
		})
	Output(w, http.StatusOK, resp)
	return

}
