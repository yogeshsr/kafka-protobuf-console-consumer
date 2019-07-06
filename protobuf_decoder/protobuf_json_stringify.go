package protobuf_decoder

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
)

type ProtobufJSONStringify struct {
	protoImportDirs          []string
	protoImportFiles         []string
	protoFileNameWithMessage string
	messageName              string
}

func NewProtobufJSONStringify(protoImportDirs []string, protoFileNameWithMessage string, messageName string) *ProtobufJSONStringify {
	protoImportFiles := []string{fmt.Sprintf("%s/%s", protoImportDirs[0], protoFileNameWithMessage)}
	return &ProtobufJSONStringify{protoImportDirs: protoImportDirs, protoImportFiles: protoImportFiles, protoFileNameWithMessage: protoFileNameWithMessage, messageName: messageName}
}

func (c *ProtobufJSONStringify) JsonString(protobufMsg []byte, prettyJson bool) (string, error) {

	fds, err := fileDescriptorsFromProtoFiles(c.protoImportDirs, c.protoImportFiles...)

	if err != nil {
		return "", err
	}

	var fd *desc.FileDescriptor
	for _, value := range fds {
		if value.GetName() == c.protoFileNameWithMessage {
			fd = value
		}
	}

	if fd == nil {
		return "", errors.New("File not found: " + c.protoFileNameWithMessage)
	}

	md := fd.FindMessage(c.messageName)
	dm := dynamic.NewMessage(md)
	err = dm.Unmarshal(protobufMsg)
	if err != nil {
		return "", err
	}
 	intent := "\t"
 	if !prettyJson {
 		intent = ""
	}
	bytes, err := dm.MarshalJSONPB(&jsonpb.Marshaler{Indent: intent})

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}


func (c *ProtobufJSONStringify) FieldValue(protobufMsg []byte, field string) (interface{}, error) {

	fds, err := fileDescriptorsFromProtoFiles(c.protoImportDirs, c.protoImportFiles...)

	if err != nil {
		return "", err
	}

	var fd *desc.FileDescriptor
	for _, value := range fds {
		if value.GetName() == c.protoFileNameWithMessage {
			fd = value
		}
	}

	if fd == nil {
		return "", errors.New("File not found: " + c.protoFileNameWithMessage)
	}

	md := fd.FindMessage(c.messageName)
	dm := dynamic.NewMessage(md)
	err = dm.Unmarshal(protobufMsg)
	if err != nil {
		return "", err
	}
	return dm.GetFieldByName(field), nil
}

func fileDescriptorsFromProtoFiles(importPaths []string, fileNames ...string) ([]*desc.FileDescriptor, error) {
	fileNames, err := protoparse.ResolveFilenames(importPaths, fileNames...)
	if err != nil {
		return nil, err
	}
	p := protoparse.Parser{
		ImportPaths:           importPaths,
		InferImportPaths:      len(importPaths) == 0,
		IncludeSourceCodeInfo: true,
	}
	fds, err := p.ParseFiles(fileNames...)
	if err != nil {
		return nil, fmt.Errorf("could not parse given files: %v", err)
	}

	return fds, nil
}
