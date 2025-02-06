package json

import (
	"bytes"
	"encoding/json"
)

func MarshalJSON(v any, escapeHTML bool) ([]byte, error) {
	if escapeHTML {
		return json.Marshal(v)
	} else {
		return marshalWithEncoder(v)
	}
}

func marshalWithEncoder(v any) ([]byte, error) {
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)

	err := enc.Encode(v)
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1) // remove trailing newline
	}

	return buf.Bytes(), err
}
