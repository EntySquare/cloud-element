package main

import (
	"encoding/json"
	logging "github.com/ipfs/go-log/v2"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var natsUrl string
var errLog = logging.Logger("ERROR")

func main() {
	log.Println("run main success")
}

func SendEvent(sbj string, src []byte) {

	natsStr, ok := os.LookupEnv("NATS_SERVER")
	if !ok {
		natsUrl = "http://localhost:4222"
	} else {
		natsUrl = natsStr
	}

	log.Println("send event to miner subject is : ", sbj)
	// Connect to a server
	nc, err := stan.Connect("knative-nats-streaming", RandString(15), stan.NatsURL(natsUrl))
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

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func ReadFile() string {
	file, err := os.Open("c2_event.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var event = Event{}
	buf, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(buf, &event)
	bin, err := json.Marshal(event)
	if err != nil {
		sendError(spec.JSON_MARSHAL_ERR, err, taskTyp)
		return
	}
	return txt
}

func sendError(code spec.Code, err error, msgType string) {

	errLog.Errorf(msgType+"_ERROR:%+v", err)
	event := spec.Event{
		Head: spec.Header{
			MsgTyp:    msgType,
			SectorNum: Uint64ToString(sectorNumber),
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

func Uint64ToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}
