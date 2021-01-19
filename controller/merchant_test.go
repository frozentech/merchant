package controller_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/frozentech/merchant/model"
)

func TestMerchantPut(t *testing.T) {
	m := model.NewMerchant()
	m.Name = "This is a test"
	m.Create()

	req, _ := http.NewRequest("PUT", "/merchant/"+m.ID, bytes.NewBuffer(model.BYTE(model.Merchant{Name: "Test"})))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	fmt.Println(response.Body.String())
}

func TestMerchantDelete(t *testing.T) {
	m := model.NewMerchant()
	m.Name = "This is a test"
	m.Create()

	req, _ := http.NewRequest("DELETE", "/merchant/"+m.ID, bytes.NewBuffer(model.BYTE(``)))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	fmt.Println(response.Body.String())
}
