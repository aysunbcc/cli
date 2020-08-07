package writer

import (
	"encoding/json"
	"io"
	"os"
)

type Writer interface {
	WriteOne(data interface{}) error
	WriteTable(data interface{}) error
}

type jsonWriter struct {
	encoder *json.Encoder
}

func (w *jsonWriter) WriteOne(data interface{}) error {
	return w.encoder.Encode(data)
}

func (w *jsonWriter) WriteTable(data interface{}) error {
	var err error

	for dataRow := range data.([]*interface{}) {
		err = w.WriteOne(dataRow)
	}
	return err
}

func JSONWriter(writer io.Writer) Writer {
	return &jsonWriter{
		encoder: json.NewEncoder(writer),
	}
}

var Default = JSONWriter(os.Stdout)
