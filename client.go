package go_avro_rpc

import (
	"github.com/hamba/avro/v2"
)

type Transport interface {
	sendData([]byte) []byte
	isActive() bool
}

type ClientAvro struct {
	tr       *Transport
	protocol *avro.Protocol
}

func (receiver *ClientAvro) initClient(protocolFiles []string) {
	for _, protocolFile := range protocolFiles {
		protocol, err := avro.ParseProtocolFile(protocolFile)
		if err != nil {
			panic(err)
		}
		receiver.protocol = protocol
	}
}
func (receiver *ClientAvro) sendHandshake() {

}

func (receiver *ClientAvro) parseHandshake() {

}

func (receiver *ClientAvro) packRequest() {

}

func (receiver *ClientAvro) unpackResponse() {

}
