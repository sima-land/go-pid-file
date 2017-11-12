# go-pid-file

Tiny tool to create and read pid files


[![Build Status](https://travis-ci.org/sima-land/go-pid-file.svg?branch=master)](https://travis-ci.org/sima-land/go-pid-file)
[![Go Report Card](https://goreportcard.com/badge/github.com/sima-land/go-pid-file)](https://goreportcard.com/report/github.com/sima-land/go-pid-file)

Before start something:

```go

pf := pid.NewFile("path/to/file")

if err := pf.Create(); err := nil {
    log.Fatal("process already running")
}
defer pf.Remove()

// start something useful

```

Before stop something:

```go

pf := pid.NewFile("path/to/file")

process, err := pf.Process();
if err != nil {
    log.Fatal(err)
}
if process != nil {
    // stop process
}

```
