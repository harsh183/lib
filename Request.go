package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRequestContentFromRequest(req *http.Request) map[string]interface{} {
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	data := buf.Bytes()
	var requestContent map[string]interface{}
	err := json.Unmarshal(data, &requestContent)
	if err != nil {
		fmt.Println(err)
	}
	return requestContent
}

func GetRequestContentFromResponse(resp *http.Response) map[string]interface{} {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	data := buf.Bytes()
	var requestContent map[string]interface{}
	err := json.Unmarshal(data, &requestContent)
	if err != nil {
		fmt.Println(err)
	}
	return requestContent
}
