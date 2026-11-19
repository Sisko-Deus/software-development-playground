package serializer

import (
	"bytes"
	"context"
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

type StringResult struct {
	Data string
	Err  error
}

func SerializeToStringAsync(
	ctx context.Context,
	s Serializer,
	in <-chan any,
	outBufSize int,
) <-chan StringResult {
	if outBufSize < 1 {
		outBufSize = 1
	}
	out := make(chan StringResult, outBufSize)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				str, err := SerializeToString(s, v)
				res := StringResult{
					Data: str,
					Err:  err,
				}
				select {
				case out <- res:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

type TypedResult[T any] struct {
	Value T
	Data  string
	Err   error
}

func SerializeToStringAsyncTyped[T any](
	ctx context.Context,
	s Serializer,
	in <-chan T,
	outBufSize int,
) <-chan TypedResult[T] {
	if outBufSize < 1 {
		outBufSize = 1
	}
	out := make(chan TypedResult[T], outBufSize)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				str, err := SerializeToString(s, v)
				res := TypedResult[T]{
					Value: v,
					Data:  str,
					Err:   err,
				}
				select {
				case <-ctx.Done():
					return
				case out <- res:
				}
			}
		}
	}()
	return out
}

func SerializeToWriterAsync(
	ctx context.Context,
	w io.Writer,
	s Serializer,
	in <-chan any,
	errsBufSize int,
) <-chan error {
	if errsBufSize < 1 {
		errsBufSize = 1
	}
	errs := make(chan error, errsBufSize)
	go func() {
		defer close(errs)
		select {
		case <-ctx.Done():
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			if err := SerializeToWriter(w, s, v); err != nil {
				select {
				case errs <- err:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return errs
}

func SerializePrintAsync(
	ctx context.Context,
	s Serializer,
	in <-chan any,
	errsBufSize int,
) <-chan error {
	return SerializeToWriterAsync(ctx, os.Stdout, s, in, errsBufSize)
}
