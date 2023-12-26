package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResCode struct {
	OK       int
	EXIST    int
	NOTFOUND int
	ISE      int
	INVALID  int
}

var Code = ResCode{
	OK:       200,
	EXIST:    201,
	NOTFOUND: 404,
	ISE:      500,
	INVALID:  406,
}

func Response(data interface{}, responseStatus int, w http.ResponseWriter) {
	jData, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	ErrorChecker(w.Write(jData))
}

func HttpReq(req *http.Request) (map[string]interface{}, string) {
	body, err := io.ReadAll(req.Body)
	if err == nil {
		jsonData := make(map[string]interface{})
		ErrorChecker(0, json.Unmarshal(body, &jsonData))
		if len(jsonData) != 0 {
			return jsonData, ""
		}
		return nil, "Request parameters not found or empty."
	}
	return nil, "Invalid request."
}
