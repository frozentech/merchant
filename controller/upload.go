package controller

import (
	"net/http"

	merchantError "github.com/frozentech/merchant/error"
	"github.com/frozentech/merchant/model"
	"github.com/gorilla/mux"
)

// Upload ...
func Upload(w http.ResponseWriter, r *http.Request) {
	var (
		upload     = model.NewUpload()
		resp       = model.Response{}
		merchant   = model.NewMerchant()
		merchantID = mux.Vars(r)["merchantId"]
		err        error
	)

	if err = merchant.FindID(merchantID); err != nil {
		resp = merchantError.StatusRecord(merchantError.RecordNotFound)
		Output(w, resp.Error.Status, resp)
		return
	}

	err = upload.Upload("myfile", r)
	if err != nil {
		Log.Print("ERROR ", err.Error())
		resp = merchantError.StatusRecord(merchantError.UploadFailed)
		Output(w, resp.Error.Status, resp)
		return
	}

	resp = merchantError.StatusRecord(merchantError.NoError)

	merchant.Image = upload.Filename
	merchant.Update()

	resp.SetResult(merchant)
	Output(w, http.StatusCreated, resp)
	return
}
