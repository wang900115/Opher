package object

import (
	"bytes"
	"encoding/gob"
)

type Serializable interface {
	gob.GobEncoder
	gob.GobDecoder
}

func Serialize[T Serializable](obj T) ([]byte, error) {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(obj)
	if err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}

func Deserialize[T Serializable](data []byte) (T, error) {
	var res T
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}
