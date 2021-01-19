package controller

import (
	"net/http"

	"github.com/frozentech/logs"
	merchantError "github.com/frozentech/merchant/error"
	"github.com/frozentech/merchant/model"
	"github.com/gorilla/mux"
)

// Member ...
func Member(w http.ResponseWriter, r *http.Request) {
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
		MemberPut(resp, r)
		return
	case http.MethodDelete:
		MemberDelete(resp, r)
		return
	case http.MethodOptions:
		Options(resp, r)
		return
	default:
		MethodNotAllowed(resp, r)
		return
	}
}

// MemberPut replace a Member record
func MemberPut(w *Response, r *http.Request) {
	var (
		member     = model.NewMember()
		request    = model.Member{}
		merchant   = model.NewMerchant()
		resp       = model.Response{}
		merchantID = mux.Vars(r)["merchantId"]
		memberID   = mux.Vars(r)["memberId"]
		err        error
	)

	if err = RequestBody(r, &request, w); err != nil {
		return
	}

	if err = member.FindID(memberID); err != nil {
		resp = merchantError.StatusRecord(merchantError.RecordNotFound)
		Output(w, resp.Error.Status, resp)
		return
	}

	if request.Email != "" {
		if !IsEmailValid(request.Email) {
			resp = merchantError.StatusRecord(merchantError.InvalidEmailAddress)
			Output(w, resp.Error.Status, resp)
			return
		}
	}

	member.ID = memberID
	member.Email = request.Email
	member.Name = request.Name
	if err = merchant.FindID(merchantID); err != nil {
		resp = merchantError.StatusRecord(merchantError.RecordNotFound)
		Output(w, resp.Error.Status, resp)
		return
	}

	if err = member.Update(); err != nil {
		resp = merchantError.StatusRecord(merchantError.DuplicateEmailAddress)
		Output(w, resp.Error.Status, resp)
		return
	}

	resp = merchantError.StatusRecord(merchantError.NoError)
	resp.SetResult(member)
	Output(w, http.StatusOK, resp)
	return
}

// MemberDelete delete a Member record
func MemberDelete(w *Response, r *http.Request) {
	var (
		member     = model.NewMember()
		merchant   = model.NewMerchant()
		resp       = model.Response{}
		merchantID = mux.Vars(r)["merchantId"]
		memberID   = mux.Vars(r)["memberId"]
		err        error
	)

	if err = member.FindID(memberID); err != nil {
		resp = merchantError.StatusRecord(merchantError.RecordNotFound)
		Output(w, resp.Error.Status, resp)
		return
	}

	if err = merchant.FindID(merchantID); err != nil {
		resp = merchantError.StatusRecord(merchantError.RecordNotFound)
		Output(w, resp.Error.Status, resp)
		return
	}

	if err = member.Delete(); err != nil {
		resp = merchantError.StatusRecord(merchantError.DatabaseOperationFailure)
		Output(w, resp.Error.Status, resp)
		return
	}

	resp = merchantError.StatusRecord(merchantError.NoError)
	Output(w, http.StatusOK, resp)
	return
}
