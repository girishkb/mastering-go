package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CallResponse struct {
	Resp *Response
	Err error
}

type Response struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func helper(ctx context.Context)  <-chan *CallResponse{
	respChan := make(chan *CallResponse,1)
	go func() {
		res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
		if err != nil {
			respChan <- &CallResponse{Resp: nil,Err:  fmt.Errorf("error in http call")}
			return
		}
		defer res.Body.Close()
		byteResp, err := ioutil.ReadAll(res.Body)

		if err != nil {
			respChan <- &CallResponse{Resp: nil,Err:  fmt.Errorf("error in reading respone")}
			return
		}

		structResp := &Response{}
		err = json.Unmarshal(byteResp, structResp)

		if err != nil {
			respChan <- &CallResponse{Resp: nil,Err:  fmt.Errorf("error in unmarshalling response")}
			return
		}
		respChan <- &CallResponse{Resp: structResp,Err:  nil,}
	}()
	return respChan
}

func GetHttpResponse(ctx context.Context) (*Response,error) {
	select {
	case <-ctx.Done():
		return nil,fmt.Errorf("context timeout, ran out of time")
	case respChan := <-helper(ctx):
		return respChan.Resp,nil
	}
}

func main()  {
	ctx,cancel := context.WithTimeout(context.Background(),1*time.Millisecond)
	defer cancel()
	res, err := GetHttpResponse(ctx)

	if err != nil {
		fmt.Printf("err %v", err)
	} else {
		fmt.Printf("res %v", res)
	}
}