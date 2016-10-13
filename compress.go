package assetgo

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
)

// Compress gzip data and return base64-encoded bytes
func Compress(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	writer := gzip.NewWriter(buf)
	if _, err := writer.Write(data); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	data = buf.Bytes()
	buf.Reset()
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	if _, err := encoder.Write(data); err != nil {
		return nil, err
	}
	if err := encoder.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
