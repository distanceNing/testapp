package statemachine

import "log"

type RetryMessage struct {
	req      interface{}
	flowName string
	cur      State
}

func NewRetryMessage() *RetryMessage {
	return &RetryMessage{}
}

type RetryMessageCodec interface {
	Deserialize(data []byte) (*RetryMessage, error)
	Serialize(sm *RetryMessage) ([]byte, error)
}

const (
	PBType = 1
	JsonType
)

type RetryMessageCodecOpt struct {
	Type int
}
type PBRetryMessageCodec struct {
}

func (pbCodec *PBRetryMessageCodec) Deserialize(data []byte) (*RetryMessage, error) {
	log.Println("protocol codec")
	return NewRetryMessage(), nil
}

func (pbCodec *PBRetryMessageCodec) Serialize(sm *RetryMessage) ([]byte, error) {
	log.Println("protocol codec")
	return []byte{}, nil
}

type JsonRetryMessageCodec struct {
}

// Deserialize /*
func (jsonCodec *JsonRetryMessageCodec) Deserialize(data []byte) (*RetryMessage, error) {
	log.Println("json codec")
	return NewRetryMessage(), nil
}

func (jsonCodec *JsonRetryMessageCodec) Serialize(sm *RetryMessage) ([]byte, error) {
	log.Println("json codec")
	return []byte{}, nil
}

func NewRetryMessageCodec(opt *RetryMessageCodecOpt) RetryMessageCodec {
	if opt.Type == PBType {
		return &PBRetryMessageCodec{}
	}
	return &JsonRetryMessageCodec{}
}
