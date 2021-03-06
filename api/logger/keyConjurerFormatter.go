package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type KeyConjurerFormatter struct {
}

func (k *KeyConjurerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Most of this code came from https://github.com/sirupsen/logrus/blob/a6c0064cfaf982707445a1c90368f956421ebcf0/json_formatter.go
	data := make(logrus.Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	data["time"] = entry.Time.Format(time.RFC3339)
	data["level"] = entry.Level.String()
	data["metadata"] = entry.Message
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		data["func"] = funcVal
		data["file"] = fileVal
	}
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	output := map[string]interface{}{
		"jsonEvent":        "keyConjurer",
		"keyConjurerEvent": data}
	encoder := json.NewEncoder(b)
	if err := encoder.Encode(output); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}
