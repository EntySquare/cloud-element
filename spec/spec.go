package spec

type Event struct {
	Head  Header
	Body  []byte
	Error string
}

type Header struct {
	MsgTyp    string
	SectorNum string
}
