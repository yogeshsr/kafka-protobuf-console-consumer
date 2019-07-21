package protobuf_decoder

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/yogeshsr/kafka-protobuf-console-consumer/proto"
	"testing"
)

func TestJSONStringifyShouldCreateJSONString(t *testing.T) {

	protoImportDirs := []string{"../proto"}
	protoFileNameWithMessage := "sample.proto"
	messageName := "sample_package.SampleMessage"
	simple := grpc.SampleMessage{
		StringField: "some-value",
		IntegerField: 99,
	}

	stringify, _ := NewProtobufJSONStringify(protoImportDirs, protoFileNameWithMessage, messageName)
	messageBytes, _ := proto.Marshal(&simple)
	s, _ := stringify.JsonString(messageBytes, false)
	fmt.Println(s)
	assert.Equal(t, "{\"stringField\":\"some-value\",\"integerField\":99}", s)

}

func TestJSONStringifyShouldGetTheFieldValue(t *testing.T) {

	protoImportDirs := []string{"../proto"}
	protoFileNameWithMessage := "sample.proto"
	messageName := "sample_package.SampleMessage"
	simple := grpc.SampleMessage{
		StringField: "some-value",
		IntegerField: 99,
	}

	stringify, _ := NewProtobufJSONStringify(protoImportDirs, protoFileNameWithMessage, messageName)
	messageBytes, _ := proto.Marshal(&simple)
	s, _ := stringify.FieldValue(messageBytes, "string_field")
	fmt.Println(s)
	assert.Equal(t, "some-value", s)

}
