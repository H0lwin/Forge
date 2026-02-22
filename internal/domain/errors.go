package domain

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrMissingRequiredFlag = errors.New("missing required flag")
	ErrToolNotFound        = errors.New("required tool not found")
	ErrValidation          = errors.New("validation error")
)

type WrappedError struct {
	Op  string
	Err error
	KV  map[string]any
}

func (e *WrappedError) Error() string {
	b := strings.Builder{}
	b.WriteString(e.Op)
	b.WriteString(": ")
	b.WriteString(e.Err.Error())
	for k, v := range e.KV {
		b.WriteString(fmt.Sprintf(" %s=%v", k, v))
	}
	return b.String()
}

func (e *WrappedError) Unwrap() error {
	return e.Err
}

func Wrap(op string, err error, kv ...any) error {
	if err == nil {
		return nil
	}
	m := map[string]any{}
	for i := 0; i+1 < len(kv); i += 2 {
		k, ok := kv[i].(string)
		if !ok {
			continue
		}
		m[k] = kv[i+1]
	}
	return &WrappedError{Op: op, Err: err, KV: m}
}
