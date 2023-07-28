package rest

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpRequestData struct {
	CashServerIP   string
	CashServerPort int
	Method         string
	Payload        []byte
	Username       string
	Password       string
}

func HttpRequestToRkeeper(reqData HttpRequestData) ([]byte, error) {
	//Отправка xml в http запросе
	fmt.Printf("ReqData: %+v", reqData)
	fmt.Printf("Req Body: %s", string(reqData.Payload))
	url := fmt.Sprintf("https://%s:%d/rk7api/v0/xmlinterface.xml", reqData.CashServerIP, reqData.CashServerPort)
	fmt.Println("Url>", url)
	data := reqData.Payload
	req, err := http.NewRequest(reqData.Method, url, bytes.NewReader(data))
	if err != nil {
		fmt.Println("new req err:", err)
	}

	req.SetBasicAuth(reqData.Username, reqData.Password)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client do err:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body error:", err)
	}
	//fmt.Println("http body:", string(body))
	return body, nil
}
