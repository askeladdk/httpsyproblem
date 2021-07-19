# HTT-Peasy Problem

[![GoDoc](https://godoc.org/github.com/askeladdk/httpsyproblem?status.png)](https://godoc.org/github.com/askeladdk/httpsyproblem)
[![Go Report Card](https://goreportcard.com/badge/github.com/askeladdk/httpsyproblem)](https://goreportcard.com/report/github.com/askeladdk/httpsyproblem)

## Overview

Package httpsyproblem provides a standard interface for handling API error responses in web applications. It implements [RFC 7807](https://datatracker.ietf.org/doc/html/rfc7807) which specifies a way to carry machine-readable details of errors in a HTTP response to avoid the need to define new error response formats for HTTP APIs.

## Install

```
go get -u github.com/askeladdk/httpsyproblem
```

## Quickstart

The two basic functions are `Wrap` and `Error`. Wrap associates an error with a status code. `Error` replies to requests by checking if an error implements `http.Handler`. Use it instead of `http.Error`.

```go
func endpoint(w http.ResponseWriter, r *http.Request) {
    httpsyproblem.Error(w, r, httpsyproblem.Wrap(http.StatusBadRequest, io.EOF))
}
```

Use the `Details` type directly if you need more control. The error value returns by `Wrap` is of type `Details`.

```go
var err error = &httpsyproblem.Details{
    Detail: "This is not the Jedi that you are looking for",
    Instance: "/jedi/obi-wan",
    Status: http.StatusNotFound,
    Title: "Jedi Mind Trick",
}
```

Embed `Details` inside another type to add custom fields and use `New` to initialize it.

```go
type MoreDetails struct {
    httpsyproblem.Details
    TraceID string `json:"trace_id" xml:"trace_id"`
}

func (err *MoreDetails) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    httpsyproblem.Serve(w, r, err)
}

var err error = &MoreDetails{
    Details: httpsyproblem.New(http.StatusBadRequest, io.EOF),
    TraceID: "42",
}
```

Read the rest of the [documentation on pkg.go.dev](https://pkg.go.dev/github.com/askeladdk/httpsyproblem). It's easy-peasy!

## License

Package httpsyproblem is released under the terms of the ISC license.
