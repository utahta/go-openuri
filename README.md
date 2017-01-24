# OpenURI

[![Build Status](https://travis-ci.org/utahta/go-openuri.svg?branch=master)](https://travis-ci.org/utahta/go-openuri)

Easy way to open an io.ReadCloser from a local file and URL.

## Install

```
$ go get -u github.com/utahta/go-openuri
```

## Usage

```go
import "github.com/utahta/go-openuri"
```

Open the file.
```go
o, err := openuri.Open("/path/to/file")
```

Open the URL.
```go
o, err := openuri.Open("http://localhost")
```

with Google App Engine
```go
o, err := openuri.Open("http://localhost", openuri.WithHTTPClient(urlfetch.Client(ctx)))
```

## Example

```go
package main

import (
	"log"
	"io/ioutil"

	"github.com/utahta/go-openuri"
)

func main() {
	//
	// Open a local file
	//
	o, err := openuri.Open("/path/to/file")
	if err != nil {
	    log.Fatal(err)
	}
	defer o.Close()
	
	b, _ := ioutil.ReadAll(o)
	log.Println(string(b))
	
	//
	// Open URL
	//
	o, err = openuri.Open("http://localhost")
	if err != nil {
	    log.Fatal(err)
	}
	defer o.Close()
	
	b, _ = ioutil.ReadAll(o)
	log.Println(string(b))
}
```

## Contributing

1. Fork it ( https://github.com/utahta/go-openuri/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

