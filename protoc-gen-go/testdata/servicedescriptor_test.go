package testdata

import (
	"testing"
	"google.golang.org/grpc"
)

func TestServiceDescriptor(t *testing.T) {
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
	if expectedDescString != ListMethodsAndStreams(desc) {
		t.Errorf("Expected desc string \n%s but was \n%s", expectedDescString, ListMethodsAndStreams(desc))
	}
}

const expectedDescString =
`Methods:
	testdata.Test/UnaryCall
Streams:
	testdata.Test/Downstream
	testdata.Test/Upstream
	testdata.Test/Bidi
`

func ListMethodsAndStreams(desc grpc.ServiceDesc) string {
	s := ""
	s += "Methods:\n"
	for _, m := range desc.Methods {
		s += "\t" + desc.ServiceName + "/" + m.MethodName + "\n"
	}
	s += "Streams:\n"
	for _, st := range desc.Streams {
		s += "\t" + desc.ServiceName + "/" + st.StreamName + "\n"
	}
	return s
}