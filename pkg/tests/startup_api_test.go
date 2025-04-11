package pkg

import (
	model "ceres/pkg/model/startup"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestCreateStartup(t *testing.T) {
	doTest(func() {
		var (
			err    error
			errMsg string
		)
		url := "http://192.168.0.201:9002/cores/startups"
		// json.Marshal
		reqParam, err := json.Marshal(model.CreateStartupRequest{
			Logo:     "logo",
			Mode:     1,
			Name:     "Startup Unit Test",
			Mission:  "Startup Unit Test Mission 1",
			Overview: "Startup Unit Test Overview 1",
			TxHash:   "",
			ChainID:  43113,
			HashTags: []string{"Developer"},
		})
		if err != nil {
			log.Fatalf("Marshal RequestParam fail, err:%v", err)
		}
		// 准备: HTTP请求
		// fmt.Println(string(reqParam))
		// return
		reqBody := strings.NewReader(string(reqParam))
		httpReq, err := http.NewRequest("POST", url, reqBody)
		if err != nil {
			fmt.Printf("NewRequest fail, url: %v, reqBody: %v, err: %v\n", url, reqBody, err)
			return
		}
		httpReq.Header.Add("Content-Type", "application/json")
		httpReq.Header.Add("X-COMUNION-AUTHORIZATION", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb21lcl91aW4iOiIxNTA2MzI1NTYzNDMyOTYiLCJleHAiOjE2NjI5ODg1NDMsImlhdCI6MTY2MjcyOTM0M30.m2nciAnXlww_lbCJJDC526CqNq9hoja-cVSV845XDoA")

		// DO: HTTP请求
		httpRsp, err := http.DefaultClient.Do(httpReq)
		if err != nil {
			fmt.Printf("do http fail, url: %v, reqBody: %v, err:%v\n", url, reqBody, err)
			return
		}
		defer httpRsp.Body.Close()

		// Read: HTTP结果
		rspBody, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			fmt.Printf("ReadAll failed, url: %v, reqBody: %v, err: %v\n", url, reqBody, err)
			return
		}

		// unmarshal: 解析HTTP返回的结果
		// 		body: {"Result":{"RequestId":"12131","HasError":true,"ResponseItems":{"ErrorMsg":"错误信息"}}}
		//
		fmt.Printf("do post http success, url: %v, reqBody: %v, body: %v %v\n", url, reqBody, string(rspBody), errMsg)

		// t.Log(string(data))
	})
}
