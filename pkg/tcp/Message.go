package server

type Message struct {
	Id   uint32
	Data []byte
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		id,
		data,
	}
}
