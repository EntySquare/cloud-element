package main

type Code int64

type Event struct {
	Head  Header
	Body  []byte
	Error string
}

type Header struct {
	MsgTyp    string
	SectorNum string
}
