package embedder

import (
	"strconv"
	"strings"
	"testing"
)

func TestEmbedder(t *testing.T) {

	hW := []byte("hello, world")

	embedded, err := Embed("asset", hW)
	if err != nil {
		t.Fatal(err)
	}

	p1 := strings.Index(string(embedded), "{")
	p2 := strings.Index(string(embedded), "}")
	bytesOnly := string(embedded)[p1+1 : p2]
	bytesOnlySlice := strings.Split(bytesOnly, ",")

	var hW2 []byte

	for _, b := range bytesOnlySlice {
		bb, err := strconv.Atoi(strings.TrimSpace(b))
		if err != nil {
			t.Fatal(err)
		}
		hW2 = append(hW2, byte(bb))
	}
	if string(hW2) != "hello, world" {
		t.Fatal("Should be able to embed an asset")
	}

}
