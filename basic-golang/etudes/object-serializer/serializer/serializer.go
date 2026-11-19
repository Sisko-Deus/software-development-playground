package serializer

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Serializer interface {
	Encode(io.Writer, any) error
}

func SerializeToString(s Serializer, v any) (string, error) {
	var buf bytes.Buffer
	err := s.Encode(&buf, v)
	if err != nil {
		return "", fmt.Errorf("Serialize error: %w", err)
	}
	return buf.String(), nil
}

func SerializeToWriter(w io.Writer, s Serializer, v any) error {
	if w == nil {
		return fmt.Errorf("Args error: nil io.Writer")
	}
	return s.Encode(w, v)
}

func SerializePrint(s Serializer, v any) error {
	return SerializeToWriter(os.Stdout, s, v)
}
