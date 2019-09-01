package protobuf_decoder

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
	"os"
	"path/filepath"
	"strings"
)

type ProtobufJSONStringify struct {
	protoFileNameWithMessage string
	messageName              string
	fds						 []*desc.FileDescriptor
}

func NewProtobufJSONStringify(protoImportDirs []string, protoFileNameWithMessage string, messageName string) (*ProtobufJSONStringify, error) {

	// Hack to avoid cmd line arg to provide abs file path
	err, protoFilePath := absFilePath(protoImportDirs, protoFileNameWithMessage)
	if err != nil {
		return nil, err
	}

	fds, err := fileDescriptorsFromProtoFiles(protoImportDirs, protoFilePath)
	if err != nil {
		return nil, err
	}

	return &ProtobufJSONStringify{fds: fds, protoFileNameWithMessage: protoFileNameWithMessage, messageName: messageName}, nil
}

func absFilePath(protoImportDirs []string, protoFileNameWithMessage string) (error, string) {
	var protoFilePath string
	var err error
	for _, protoDir := range protoImportDirs {
		fileWithProtoDirPath := filepath.Join(protoDir, protoFileNameWithMessage)
		if _, err = os.Stat(fileWithProtoDirPath); err == nil {
			protoFilePath = fileWithProtoDirPath
			break
		}
	}
	if len(protoFilePath) == 0 {
		return errors.New(fmt.Sprintf("File: %s not found in: %v\n", protoFileNameWithMessage, protoImportDirs)), ""
	}
	return err, protoFilePath
}

func (c *ProtobufJSONStringify) JsonString(protobufMsg []byte, prettyJson bool) (string, error) {

	var fd *desc.FileDescriptor
	for _, value := range c.fds {
		if strings.HasSuffix(c.protoFileNameWithMessage, value.GetName()) {
			fd = value
		}
	}

	if fd == nil {
		return "", errors.New("File not found: " + c.protoFileNameWithMessage)
	}

	md := fd.FindMessage(c.messageName)
	dm := dynamic.NewMessage(md)
	err := dm.Unmarshal(protobufMsg)
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

	var fd *desc.FileDescriptor
	for _, value := range c.fds {
		if value.GetName() == c.protoFileNameWithMessage {
			fd = value
		}
	}

	if fd == nil {
		return "", errors.New("File not found: " + c.protoFileNameWithMessage)
	}

	md := fd.FindMessage(c.messageName)
	dm := dynamic.NewMessage(md)
	err := dm.Unmarshal(protobufMsg)
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
