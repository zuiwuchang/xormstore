package encode

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// Encode Encode data
func Encode(data interface{}) ([]byte, error) {
	// return JsonEncode(data)
	return GobEncode(data)
}

// Decode decode data
func Decode(data []byte, to interface{}) error {
	// return JsonDecode(data, to)
	return GobDecode(data, to)
}

// GobEncode encode data with gob
func GobEncode(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// GobDecode decode data with gob
func GobDecode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

// JsonEncode encode data with json
func JsonEncode(data interface{}) ([]byte, error) {
	val, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// JsonDecode decode data with json
func JsonDecode(data []byte, to interface{}) error {
	return json.Unmarshal(data, to)
}
