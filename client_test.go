package go_avro_rpc

import (
	"fmt"
	"github.com/hamba/avro/v2"
	"testing"
)

func TestClient(t *testing.T) {
	//dat, err := os.ReadFile("./testdata/mail.avpr")
	//if err != nil {
	//	panic(err)
	//}
	//protocol, err := avro.ParseProtocol(string(dat))
	//wantedMd5 := "042b93394447006d0f9bd8e1f8e18b08"
	//fmt.Print(dat)
	//assert.Equal(t, wantedMd5, protocol.Hash())
	//fmt.Print(protocol.Hash())
	type Message struct {
		to   string `avro:"to"`
		from string `avro:"from"`
		body string `avro:"body"`
	}
	type SimpleRecord struct {
		message Message `avro:"message"`
	}
	client := ClientAvro{
		tr:       nil,
		protocol: nil,
	}
	client.initClient([]string{"./testdata/mail.avpr"})
	msg := client.protocol.Message("send")
	requestSchema := msg.Request()
	in := SimpleRecord{message: Message{to: "to", from: "dsds", body: "dsdsd"}}
	data, err := avro.Marshal(requestSchema, in)
	fmt.Print(data, err)
}
