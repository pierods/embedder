# embedder

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![](https://godoc.org/github.com/pierods/embedder?status.svg)](http://godoc.org/github.com/pierods/embedder)
[![Go Report Card](https://goreportcard.com/badge/github.com/pierods/embedder)](https://goreportcard.com/report/github.com/pierods/embedder)
[![Build Status](https://travis-ci.org/pierods/embedder.svg?branch=master)](https://travis-ci.org/pierods/embedder)

Embedder  is a function/command line utility to embed a binary or non-binary resource into a Go package.

## Usage
As a library:

```Go
    hW := []byte("hello, world")
    
    embedded, err := Embed("AssetVar", hW)
    if err != nil {
        t.Fatal(err)
    }
    embedded = append([]byte("package mypackage \n"), embedded...)
    if err != nil {
        ...
    }
    ioutil.WriteFile("MyEmbeddedAsset.go", embedded, os.ModePerm)
    ...

	
```

As a command-line utility:
```Go
    embed -package mypackage -var MyImageGif -asset myImage.gif > MyEmbeddedAsset.go
```

In go:generate:

```Go
//go:generate  embed -package mypackage -var MyImageGif -asset myImage.gif -o MyEmbeddedAsset.go 
```

## Installation:

```Go
    go get github.com/pierods/embedder
    cd %GOPATH/src/github.com/pierods/embedder/embed
    go install
```

## Why not go-binddata, statik etc 

go-bindata: 47 issues, compiles with Makefile

statik: it creates a statikFS fileserver - I don't need a fileserver

rice: it crates an HTTPBox - I don't need a box.

etc.

Embedder does one thing and does it well - embedding resources. You can then wrap the embedded resource in a file server or whatever.

