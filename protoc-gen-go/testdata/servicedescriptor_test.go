package testdata

import (
	"testing"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/proto"
	"compress/gzip"
	"bytes"
	"io/ioutil"
)

func TestServiceDesc(t *testing.T) {
	desc := GetTestServiceDesc()
	if desc.ServiceName != "testdata.Test" {
		t.Error("Service name does not match")
	}
	if len(desc.Methods) != 1 {
		t.Error("Expected 1 method")
	}
	if len(desc.Streams) != 3 {
		t.Error("Expected 3 streams")
	}
	if desc.Methods[0].MethodName != "UnaryCall" {
		t.Errorf("Expected UnaryCall method but was %s", desc.Methods[0].MethodName)
	}
	if desc.Streams[0].StreamName != "Downstream" {
		t.Errorf("Expected Downstream stream but was %s", desc.Streams[0].StreamName)
	}
	if desc.Streams[1].StreamName != "Upstream" {
		t.Errorf("Expected Upstream stream but was %s", desc.Streams[1].StreamName)
	}
	if desc.Streams[2].StreamName != "Bidi" {
		t.Errorf("Expected Bidi stream but was %s", desc.Streams[2].StreamName)
	}
}

func TestServiceDescriptorExposed(t *testing.T) {
	data, idx := GetTestServiceDescriptor()
	if len(idx) != 1 || idx[0] != 0 {
		t.Errorf("Expected idx to be [0], but was %v", idx)
	}
	decompressor, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	decompressed, err := ioutil.ReadAll(decompressor)
	if err != nil {
		t.Fatal(err)
	}
	desc := &descriptor.FileDescriptorProto{}
	if err := proto.Unmarshal(decompressed, desc); err != nil {
		t.Fatal(err)
	}
	if len(desc.Service) != 1 {
		t.Errorf("Expected 1 service, but found %v", len(desc.Service))
	}
	svc := desc.Service[0]
	if svc.GetName() != "Test" {
		t.Errorf("Expected service name to be Test, but was %s", svc.GetName())
	}
}