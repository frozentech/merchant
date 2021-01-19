package controller_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/frozentech/merchant/model"
)

func TestMemberPut(t *testing.T) {
	m := model.NewMerchant()
	m.Name = "This is a test"
	m.Create()

	mm := model.NewMember()
	mm.MerchantID = m.ID
	mm.Name = "Lester Soriano"
	mm.Email = "email@email.com"
	mm.Create()

	req, _ := http.NewRequest("PUT", "/merchant/"+m.ID+"/member/"+mm.ID, bytes.NewBuffer(model.BYTE(model.Member{
		Name:  "Test",
		Email: "email@email.com",
	})))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	fmt.Println(response.Body.String())
}

func TestMemberDelete(t *testing.T) {
	m := model.NewMerchant()
	m.Name = "This is a test"
	m.Create()

	mm := model.NewMember()
	mm.MerchantID = m.ID
	mm.Name = "Lester Soriano"
	mm.Email = "email@email.com"
	mm.Create()

	req, _ := http.NewRequest("DELETE", "/merchant/"+m.ID+"/member/"+mm.ID, bytes.NewBuffer(model.BYTE(``)))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	fmt.Println(response.Body.String())
}
