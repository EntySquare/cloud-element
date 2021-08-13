package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println("run test api /params /response")
	http.HandleFunc("/params", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		req := r.Body
		paramReq, err := ioutil.ReadAll(req)
		if err != nil {
			fmt.Println(err)
		}
		fp := &FetchParams{}
		err = json.Unmarshal(paramReq, fp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fp)
		inb, err := ioutil.ReadFile(".././file/c2.params")
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(inb)
		//fmt.Println(err)
		var c2in2 Commit2Params2
		var c2in Commit2Params
		if err := json.Unmarshal(inb, &c2in2); err != nil {
			fmt.Println(err)
		}
		c2in.Commit1Out = c2in2.Phase1Out
		paramsJson, err := json.Marshal(c2in)
		if err != nil {
			fmt.Println(err)
		}
		_, err = w.Write(paramsJson)
	})
	http.HandleFunc("/response", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		req := r.Body
		body, err := ioutil.ReadAll(req)
		if err != nil {
			fmt.Println(err)
		}
		pr := &PostResp{}
		err = json.Unmarshal(body, pr)
		if err != nil {
			fmt.Println(err)
		}
	})
	_ = http.ListenAndServe("localhost:9999", nil)
}

type FetchParams struct {
	SectorNum string
	TaskType  string
}
type Commit2Params struct {
	Commit1Out []byte
}
type Commit2Params2 struct {
	Phase1Out []byte
}
type PostResp struct {
	SectorNum string
	TaskType  string
	Body      []byte
}
