package serializer_test

import (
	"bytes"
	"io"
	"object-serializer/serializer"
	"os"
	"testing"
)

type MockPart struct {
	MP1 int
	MP2 string
}

type Mock struct {
	M1 int
	M2 MockPart
}

func buildMocks() (obj Mock, want string) {
	obj = Mock{
		M1: 101,
		M2: MockPart{
			MP1: 201,
			MP2: "TEST",
		},
	}
	want = `{"M1":101,"M2":{"MP1":201,"MP2":"TEST"}}` + "\n"
	return obj, want
}

func TestSerializeToString(t *testing.T) {
	data, want := buildMocks()
	srlz := serializer.NewJSONSerializer()
	got, err := serializer.SerializeToString(srlz, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSerializeToWriter(t *testing.T) {
	data, want := buildMocks()
	srlz := serializer.NewJSONSerializer()
	var buf bytes.Buffer
	err := serializer.SerializeToWriter(&buf, srlz, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSerializePrint(t *testing.T) {
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	data, want := buildMocks()
	srlz := serializer.NewJSONSerializer()

	err = serializer.SerializePrint(srlz, data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_ = w.Close()
	var buf bytes.Buffer
	if _, err = io.Copy(&buf, r); err != nil {
		t.Fatalf("error reading pipe: %v", err)
	}
	got := buf.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
