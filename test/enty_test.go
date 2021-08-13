package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test1(t *testing.T) {
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
		inb, err := ioutil.ReadFile("../file/c2.params")
		if err != nil {
			inb, err = ioutil.ReadFile("../../file/c2.params")
		}
		var c2in Commit2Params
		if err := json.Unmarshal(inb, &c2in); err != nil {
			fmt.Println(err)
		}
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
}

type FetchParams struct {
	SectorNum string
	TaskType  string
}
type Commit2Params struct {
	Commit1Out []byte
}
type PostResp struct {
	SectorNum string
	TaskType  string
	Body      []byte
}
