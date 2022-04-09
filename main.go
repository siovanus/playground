package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	_, err := ethGetCurrentHeight("https://damp-empty-morning.matic.quiknode.pro/63f6dc0c1408881c521f255e7902d02d27531370/")
	if err != nil {
		panic(err)
	}
}

type heightReq struct {
	JSONRPC string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	ID      uint     `json:"id"`
}

type heightRep struct {
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      uint   `json:"id"`
}

func ethGetCurrentHeight(url string) (height uint64, err error) {
	req := &heightReq{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		Params:  make([]string, 0),
		ID:      1,
	}
	data, _ := json.Marshal(req)

	body, err := jsonRequest(url, data)
	if err != nil {
		return
	}

	var resp heightRep
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	height, err = strconv.ParseUint(resp.Result, 0, 64)
	if err != nil {
		return
	}

	return
}

func jsonRequest(url string, data []byte) (result []byte, err error) {
	start := time.Now()
	resp, err := http.Post(url, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return
	}
	elapsed := time.Since(start)
	fmt.Println("http post 执行完成耗时：", elapsed)
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}