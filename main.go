package main

import (
	"encoding/json"
	"enty/clouder-element/types"
	"fmt"
	logging "github.com/ipfs/go-log/v2"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var natsUrl, taskType, minerIDStr string
var sectorNumber uint64
var errLog = logging.Logger("ERROR")

func main() {

	if err := ReadFileEnv("../env.json"); err != nil {
		panic("env err!")
	}
	ReadFileLaterSendEvent("../c2_event.json")
	log.Println("run main success")
}

func ReadFileLaterSendEvent(dirFile string) {
	file, err := os.Open(dirFile) //c2_event.json
	if err != nil {
		fmt.Println("[cloud-element] open file err:" + err.Error())
		sendError(spec.JSON_MARSHAL_ERR, err, taskType)
		return
	}
	defer file.Close()
	var event = spec.Event{}
	buf, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(buf, &event)
	bin, err := json.Marshal(event)
	if err != nil {
		fmt.Println("[cloud-element] json marshal err:" + err.Error())
		sendError(spec.JSON_MARSHAL_ERR, err, taskType)
		return
	}
	fmt.Println("[cloud-element] run SendEvent:" + minerIDStr)
	fmt.Println(fmt.Sprintf("[cloud-element] event:\n %+v", bin))
	SendEvent(spec.MinerTopicSealerDone(minerIDStr), bin)
}

func SendEvent(sbj string, src []byte) {

	log.Println("send event to miner subject is : ", sbj)
	// Connect to a server
	nc, err := stan.Connect("knative-nats-streaming", spec.RandString(15), stan.NatsURL(natsUrl))
	if err != nil {
		log.Println("send event connection error : ", err)
		SendEvent(sbj, src)
		return
	}

	// Simple Publisher
	err = nc.Publish(sbj, src)
	if err != nil {
		log.Println("send event publish error : ", err)
		SendEvent(sbj, src)
		return
	}

	// Close connection
	nc.Close()
}

func sendError(code spec.Code, err error, msgType string) {

	errLog.Errorf(msgType+"_ERROR:%+v", err)
	event := spec.Event{
		Head: spec.Header{
			MsgTyp:    msgType,
			SectorNum: spec.Uint64ToString(sectorNumber),
		},
		Body:  nil,
		Error: spec.NewFiltabErr(code, err).ToString(),
	}
	bin, err2 := json.Marshal(event)
	if err2 != nil {
		errLog.Errorf("JSON_MARSHAL_ERROR:%+v", err2)
		time.Sleep(time.Second * 5)
		sendError(code, err, msgType)
		return
	}
	SendEvent(spec.MinerTopicSealerDone(minerIDStr), bin)
}

func ReadFileEnv(dirFile string) error {
	file, err := os.Open(dirFile) //c2_event.json
	if err != nil {
		return err
	}
	defer file.Close()
	buf, _ := ioutil.ReadAll(file)
	var maps = make(map[string]interface{})
	if err = json.Unmarshal(buf, &maps); err != nil {
		return err

	}
	if v, ok := maps["natsUrl"]; ok {
		natsUrl = v.(string)
	} else {
		return fmt.Errorf("file natsUrl is null")
	}

	if v, ok := maps["taskType"]; ok {
		taskType = v.(string)
	} else {
		return fmt.Errorf("file taskType is null")
	}

	if v, ok := maps["minerIDStr"]; ok {
		minerIDStr = v.(string)
	} else {
		return fmt.Errorf("file minerIDStr is null")
	}
	if v, ok := maps["sectorNumber"]; ok {
		sectorNumber = v.(uint64)
	} else {
		return fmt.Errorf("file sectorNumber is null")
	}
	return nil
}
