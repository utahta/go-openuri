# OpenURI

[![Build Status](https://travis-ci.org/utahta/go-openuri.svg?branch=master)](https://travis-ci.org/utahta/go-openuri)

Easy to open an io.ReadCloser from a local file or URL.

## Install

```
$ go get -u github.com/utahta/go-openuri
```

## Usage

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
	o, err := openuri.Open("/path/to/file.txt")
	if err != nil {
	    log.Fatal(err)
	}
	
	b, _ := ioutil.ReadAll(o)
	log.Println(string(b))
	
	//
	// Open URL
	//
	o, err = openuri.Open("http://example.com")
	if err != nil {
	    log.Fatal(err)
	}
	
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

