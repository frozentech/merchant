package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"github.com/frozentech/logs"
	merchantError "github.com/frozentech/merchant/error"
	"github.com/frozentech/merchant/model"
)

// Log ...
var Log *logs.Log

// Options ...
func Options(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Allow", "POST, PUT, GET, DELETE")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Length", "0")
	w.WriteHeader(200)
}

// MethodNotAllowed ...
func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Allow", "POST, PUT, GET, DELETE")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Length", "0")
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// Strpad ...
func Strpad(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "BOTH":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}

// ReadBody ...
func ReadBody(r *http.Request) []byte {
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes
}

// RequestBody ...
func RequestBody(r *http.Request, request interface{}, w http.ResponseWriter) error {
	d := json.NewDecoder(r.Body)

	if err := d.Decode(request); err != nil {
		resp := merchantError.StatusRecord(merchantError.RequestBodyStatusUnprocessableEntity)
		Output(w, resp.Error.Status, resp)
		return err
	}

	return nil
}

// Output ...
func Output(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(model.BYTE(v))
	// fmt.Fprintf(w, model.JSON(v))
}
