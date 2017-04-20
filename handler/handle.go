package handler

import (
	"net/http"
	"fmt"
	"encoding/json"

	"gopush/lib"
    qrcode "github.com/skip2/go-qrcode"
)

type Api struct {
}

func NewApi() *Api {
	return &Api{}
}

// qrcode API
//
// DESC: Geneqrcode
// Params:
//		token: secret token
//		message: qrcode message
func (api *Api) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method != lib.HTTP_METHOD_GET {
		formatNormalResponceHeaderText(w)
		api.OutputResponse(w, "HTTP method POST is required.")
		return
	}

	token, err := GetParamString(r, "token")
	if err != nil {
		formatNormalResponceHeaderText(w)
		api.OutputResponse(w, "Param token is required.")
		return
	}

	if token != "4yPqRtS3KHJyF9o5XEiM93FGJqafpd55vEiYOa"{
		formatNormalResponceHeaderText(w)
		api.OutputResponse(w, "Param token is invalid.")
		return
	}

	message, err := GetParamString(r, "message")
	if err != nil {
		formatNormalResponceHeaderText(w)
		api.OutputResponse(w, "Param message is required.")
		return
	}


	var png []byte
	png, err = qrcode.Encode(message, qrcode.Medium, 256);
	if err != nil {
		formatNormalResponceHeaderText(w)
		api.OutputResponse(w, err.Error())
		return
	}

	//formatNormalResponceHeaderText(w)
	formatNormalResponceHeaderPng(w);
	w.Write(png);
	return
}


func (api *Api) FormatResponseJson(resp interface{}) (string, error) {
	result, err := json.Marshal(resp)
	if err != nil {
		return "", err
	}else {
		return string(result), nil
	}
}

func (api *Api) OutputResponse(w http.ResponseWriter, resp interface{}) {
	fmt.Fprintln(w, resp)
}