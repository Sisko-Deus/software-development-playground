package serializer

import (
	"encoding/json"
	"io"
	"time"
)

type JSONSerializer struct {
	EscapeHTML       bool
	IncludeTimestamp bool
	TimestampFormat  string
	Now              func() time.Time
}

func NewJSONSerializer() *JSONSerializer {
	return &JSONSerializer{
		EscapeHTML:      true,
		TimestampFormat: time.RFC3339,
		Now:             time.Now,
	}
}

func (s *JSONSerializer) Encode(w io.Writer, v any) error {
	data := v
	if s.IncludeTimestamp {
		format := s.TimestampFormat
		if format == "" {
			format = time.RFC3339
		}
		now := s.Now
		if now == nil {
			now = time.Now
		}
		data = struct {
			CreatedAt string `json:"created_at"`
			Data      any    `json:"data"`
		}{
			CreatedAt: now().UTC().Format(format),
			Data:      data,
		}
	}

	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(s.EscapeHTML)
	return encoder.Encode(data)
}
