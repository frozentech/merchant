package controller_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/frozentech/merchant/model"
)

func TestMerchantsPost(t *testing.T) {
	req, _ := http.NewRequest("POST", "/merchant", bytes.NewBuffer(model.BYTE(model.Merchant{Name: "Test"})))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)
	fmt.Println(response.Body.String())
}

func TestMerchantsEmptyPost(t *testing.T) {
	req, _ := http.NewRequest("POST", "/merchant", bytes.NewBuffer([]byte(``)))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
	fmt.Println(response.Body.String())
}

func TestMerchantsGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/merchant", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	fmt.Println(response.Body.String())
}

func TestMerchantsInvalidPageGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/merchant?page=a", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
	fmt.Println(response.Body.String())
}

func TestMerchantsInvalidLimitGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/merchant?limit=a", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusBadRequest, response.Code)
	fmt.Println(response.Body.String())
}
