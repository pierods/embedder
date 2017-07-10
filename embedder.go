/*
Package embedder contains a single function to embed (binary) assets in a program
*/
package embedder

import (
	"bytes"
	"strconv"
)

/*
Embed takes in an output variable name and a []byte and returns this string:
var OuputVar []byte = []byte{...}
The content of assetBytes is converted to its string representation.
*/
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
