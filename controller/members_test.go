package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/frozentech/merchant/model"
)

func TestMembersPost(t *testing.T) {
	m := model.NewMerchant()
	m.Name = "This is a test"
	m.Create()

	req, _ := http.NewRequest("POST", "/merchant/"+m.ID+"/member", bytes.NewBuffer(model.BYTE(model.Member{
		Name:  "Test",
		Email: "email@email.com",
	})))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)
	fmt.Println(response.Body.String())
}

func TestMembersGet(t *testing.T) {

	m := model.NewMerchant()
	m.Name = "This is a test"
	m.Create()

	mm := model.NewMember()
	mm.MerchantID = m.ID
	mm.Name = "Lester Soriano"
	mm.Email = "email@email.com"
	mm.Create()

	req, _ := http.NewRequest("GET", "/merchant/"+m.ID+"/member", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	resp := model.Response{}
	json.Unmarshal([]byte(response.Body.String()), &resp)
	fmt.Println(response.Body.String())

	if fmt.Sprint(resp.Result) == "[]" {
		t.Fail()
	}

}
