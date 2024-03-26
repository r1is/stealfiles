package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"upload_oss/common"
)

func GetServerURL() string {
	e1 := "7ee90589b0e2cda77765931e3591fe82d103a37756acf0124ffce97a0780b5fed99c25f0d5fa2f6c5610ddb33c6c9ce28edd2656d9091f0b9ee4e2674bc0d03056b3fbba1901f7a572b27df4ac4e1659220080407822c37b5eb73cdfa23e41eb"
	targetURL, _ := common.Sm4_d("speedtest", e1)
	return targetURL
}

func GetOssCfg(targetURL string, argsInfo common.Args) (common.OssCfg, error) {
	var ossCfg common.OssCfg
	encCode, _ := common.Sm4_e("10086.com", argsInfo.Passcode)
	data := &common.Data{Code: encCode}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON marshaling failed:", err)
	}

	req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error creating the request:", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending the request:", err)
	}
	defer resp.Body.Close()

	// 读取并解析返回的 JSON 数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}
	var response common.Resp

	err = json.Unmarshal(body, &response)
	if err != nil {
		return ossCfg, err
	}

	fmt.Println("Response Code:", response.Code)
	if response.Code != 0 {
		fmt.Println("Response Msg:", response.Msg)
		return ossCfg, errors.New("esponse Code is not 0")
	}
	Encoss := response.Msg
	_key := "10086.com" + argsInfo.Passcode
	text, _ := common.Sm4_d(_key, Encoss)
	json.Unmarshal([]byte(text), &ossCfg)
	return ossCfg, nil
}
