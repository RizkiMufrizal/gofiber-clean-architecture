package common

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/RizkiMufrizal/gofiber-clean-architecture/exception"
	"io"
	"math/rand"
	"net"
	"net/http"
	"reflect"
	"time"
)

type HttpHeader struct {
	Key   string
	Value string
}

type ClientComponent[T any, E any] struct {
	HttpMethod     string
	UrlApi         string
	ConnectTimeout uint32
	ActiveTimeout  uint32
	Headers        []HttpHeader
	RequestBody    *T
	ResponseBody   *E
}

func (c *ClientComponent[T, E]) Execute(ctx context.Context) error {

	client := &http.Client{
		Timeout: time.Duration(rand.Int31n(int32(c.ActiveTimeout))) * time.Millisecond,
		Transport: &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			TLSHandshakeTimeout: 5 * time.Second,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Duration(rand.Int31n(int32(c.ConnectTimeout)))*time.Millisecond)
			},
		},
	}

	var request *http.Request
	var response *http.Response
	var err error = nil

	//set request body
	if reflect.ValueOf(c.RequestBody).IsZero() || c.RequestBody == nil {
		request, err = http.NewRequest(c.HttpMethod, c.UrlApi, nil)
	} else {
		requestBody, err := json.Marshal(c.RequestBody)
		exception.PanicLogging(err)

		//logging request body
		NewLogger().Info("Request Body ", string(requestBody))

		requestBodyByte := bytes.NewBuffer(requestBody)

		request, err = http.NewRequestWithContext(ctx, c.HttpMethod, c.UrlApi, requestBodyByte)
		exception.PanicLogging(err)
	}

	//set header
	request.Header.Set("Content-Type", "application/json")
	for _, header := range c.Headers {
		request.Header.Set(header.Key, header.Value)
	}

	//logging before
	NewLogger().Info("Request Url ", c.UrlApi)
	NewLogger().Info("Request Method ", c.HttpMethod)
	NewLogger().Info("Request Header ", request.Header)

	//time
	start := time.Now()

	response, err = client.Do(request)
	//error handling for http client
	if err != nil {
		return err
	}

	//time
	elapsed := time.Now().Sub(start)

	responseBody, err := io.ReadAll(response.Body)
	exception.PanicLogging(err)

	err = json.Unmarshal(responseBody, &c.ResponseBody)
	exception.PanicLogging(err)

	NewLogger().Info("Received response for ", c.UrlApi, " in ", elapsed.Milliseconds(), " ms")
	NewLogger().Info("Response Header ", response.Header)
	NewLogger().Info("Response Http Status ", response.Status)
	NewLogger().Info("Response Http Version ", response.Proto)
	NewLogger().Info("Response Body ", string(responseBody))

	return nil
}
