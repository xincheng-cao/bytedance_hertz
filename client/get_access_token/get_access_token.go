package get_access_token

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"os"
)

type getTokenReqBody struct {
	AppCode   string `json:"appCode"`
	AppSecret string `json:"appSecret"`
}

func GetAccessToken() {
	// tls request
	// https://github.com/cloudwego/hertz-examples/blob/main/protocol/tls/main.go
	clientCfg := &tls.Config{InsecureSkipVerify: false}

	var app_code, app_secret string
	app_secret = os.Getenv("APPSECRET")
	app_code = os.Getenv("APPCODE")

	c, err := client.NewClient(
		client.WithTLSConfig(clientCfg),
		client.WithDialer(standard.NewDialer()),
	)
	if err != nil {
		return
	}
	/*
		var postArgs protocol.Args
		postArgs.Set("appCode", app_code)
		postArgs.Set("appSecret", app_secret)

		//var dst []byte this is nil
		var dst2 = make([]byte, 1, 2)
		status, body, _ := c.Post(context.Background(),
			dst2, "http://127.0.0.1:8888/post", &postArgs)
		// And the content-type used is application/x-www-url-form
		// https://github.com/cloudwego/hertz-examples/tree/main/client/send_request

		fmt.Printf("status=%v body=%v\n, dst=%v", status, string(body), string(dst2))
	*/
	req, res := protocol.AcquireRequest(), protocol.AcquireResponse()
	defer func() {
		protocol.ReleaseRequest(req)
		protocol.ReleaseResponse(res)
	}()

	req.Header.SetMethod(consts.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	req.SetRequestURI("https://open.dingdanll.com/open/platform/pub/getToken")
	reqBody := getTokenReqBody{
		AppCode:   app_code,
		AppSecret: app_secret,
	}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req.SetBody(reqBodyBytes)
	err = c.Do(context.Background(), req, res)

	fmt.Println(string(res.Body()))
}
