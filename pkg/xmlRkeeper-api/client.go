package XmlRkeeper

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	CashServerIP   string
	CashServerPort int
	Username       string
	Password       string
	UserAgent      string

	httpClient *http.Client
}

func NewClient(ip string, port int, username string, password string) *Client {
	return &Client{
		CashServerIP:   ip,
		CashServerPort: port,
		Username:       username,
		Password:       password,
	}
}

func (c *Client) newRequest(method string, body []byte) (*http.Request, error) {
	url := fmt.Sprintf("https://%s:%d/rk7api/v0/xmlinterface.xml", c.CashServerIP, c.CashServerPort)

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/xml")
	}
	req.SetBasicAuth(c.Username, c.Password)

	req.Header.Set("Accept", "application/xml")
	//req.Header.Set("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c.httpClient = &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(body))
	return body, err
}
