package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pierods/embedder"
)

var (
	outputPackage *string = flag.String("package", "", "package containing the embedded asset")
	outputVar     *string = flag.String("var", "", "name of the variable containing the embedded asset")
	asset         *string = flag.String("asset", "", "path of the asset to be embedded")
	outFile       *string = flag.String("o", "", "path of output file")
)

func main() {

	flag.Parse()
	errFunc := func(msg string) {
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
	if *outputPackage == "" {
		errFunc("package can not be empty")
	}
	if *outputVar == "" {
		errFunc("var cannot be empty")
	}
	if *asset == "" {
		errFunc("asset cannot be empty")
	}

	assetBytes, err := ioutil.ReadFile(*asset)
	if err != nil {
		errFunc(err.Error())
	}

	embedded, err := embedder.Embed(*outputVar, assetBytes)
	if err != nil {
		errFunc(err.Error())
	}

	embedded = append([]byte("package "+*outputPackage+"\n"), embedded...)
	if err != nil {
		os.Exit(1)
	}

	if *outFile != "" {
		err = ioutil.WriteFile(*outFile, embedded, os.ModePerm)
		if err != nil {
			os.Exit(1)
			return
		}
		return
	}
	_, err = os.Stdout.Write(embedded)

	if err != nil {
		os.Exit(1)
	}

	return
}
