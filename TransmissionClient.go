package main

import (
	"github.com/go-resty/resty"
	"log"
	"fmt"
)

type TransmissionClient struct {
	url string
	username string
	password string
}

func NewTransmissionClient (url string, username string, password string) *TransmissionClient {
	return &TransmissionClient{url, username, password}
}

func (client *TransmissionClient) addTorrent(uri string) {
	resp, err := client.callAddTorrent(uri, "")
	if (err != nil) {
		log.Fatalf("error %v", err)
	} else {
		// Add X-Transmission-Session-Id header on status 409
		if (resp.StatusCode() == 409) {
			resp, err = client.callAddTorrent(uri, getSessionId(resp))
		}
		fmt.Printf("Response: [%s]", resp.Body())

	}
}

func getSessionId(response *resty.Response) string {
	return response.Header().Get("X-Transmission-Session-Id")
}

func (client *TransmissionClient) buildRequest(sesId string)  *resty.Request {
	request := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Transmission-Session-Id", sesId)
	if(client.username != "") {
		request.SetBasicAuth(client.username, client.password)
	}
	return request
}

func (client *TransmissionClient) callAddTorrent(uri string, sesId string) (resp *resty.Response, err error) {
	body := fmt.Sprintf(`{"method":"torrent-add", "arguments": {"filename": "%s"}}`, uri)
	request := client.buildRequest(sesId)
	request.SetBody(body)
	resp, err = request.Post(client.url)
	return resp, err
}