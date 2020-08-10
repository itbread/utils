package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	HTTP_GET_METHOD    = "GET"
	HTTP_POST_METHOD   = "POST"
	HTTP_DELETE_METHOD = "DELETE"
)

func http_request(method string, url string, mp map[string]string, args ...interface{}) error {
	//mp := make(map[string]interface{})
	var (
		reqStruct  interface{}
		respStruct interface{}
		req        *http.Request
		resp       *http.Response
		client     http.Client
		data       string
		respData   []byte
		err        error
	)

	switch len(args) {
	case 1:
		respStruct = args[0]
	case 2:
		reqStruct = args[0]
		respStruct = args[1]
	default:
		return errors.New("function must have 2 or 3 arguments")
	}
	if sendBody, jsonErr := json.Marshal(reqStruct); jsonErr != nil {
		return jsonErr
	} else {
		data = string(sendBody)
	}
	log.Printf("Send Data:%v\n", data)
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, err = http.NewRequest(method, url, strings.NewReader(data))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	if len(mp) > 0 {
		for k, v := range mp {
			req.Header.Set(k,v)
		}
	}
	//req.Header.Add("Content-Type", "application/json; charset=utf-8") //设置content-type

	resp, err = client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("StatusCode=%v", resp.StatusCode))
	}
	respData, err = ioutil.ReadAll(resp.Body)
	log.Printf("Response Data:%s\n", respData)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	buffer.Write(respData)
	err = json.NewDecoder(buffer).Decode(&respStruct)
	if err != nil {
		return err
	}
	return nil
}



