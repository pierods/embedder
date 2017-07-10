package embedder

import (
	"bytes"
	"strconv"
)

func Embed(outputVar string, assetBytes []byte) ([]byte, error) {

	var buffer bytes.Buffer

	_, err := buffer.Write([]byte("var " + outputVar + " []byte = []byte{"))
	if err != nil {
		return nil, err
	}

	l := len(assetBytes)
	for i, assetByte := range assetBytes {
		if i != l-1 {
			_, err = buffer.Write([]byte(strconv.Itoa(int(assetByte))))
			if err != nil {
				return nil, err
			}
			_, err = buffer.Write([]byte(", "))
			if err != nil {
				return nil, err
			}

		} else {
			_, err = buffer.Write([]byte(strconv.Itoa(int(assetByte))))
			if err != nil {
				return nil, err
			}
			_, err = buffer.Write([]byte("}\n\n"))
			if err != nil {
				return nil, err
			}

		}
	}
	return buffer.Bytes(), nil
}
