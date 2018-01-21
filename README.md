# About bsconv [![GoDoc](https://godoc.org/github.com/belinskiydm/bsconv?status.svg)](https://godoc.org/github.com/belinskiydm/bsconv) #
Golang package to convert any number from an arbitrary base to another arbitrary base. From binary up to base62. This package can work with numbers bigger than int64.

The package also contains two additional functions to convert from an arbitrary base to decimal and from decimal to an arbitrary base.

All functions get and return numbers in the format of string.

## Installation ##

Install the package via the following:

    go get -u github.com/belinskiydm/bsconv

## More Details ##

An example of the main.go file:
```go
// example/main.go
package main

import (
	"fmt"

	"github.com/belinskiydm/bsconv"
)

func main() {
	numBase62 := "gdguygdgdgdghjjgfdfdgdgdgd767"
	numBase10 := "674386482479820024"
	numBase2 := "110101011110101110101010"
	revNumBase10 := "2493771263131529637640109659423615368150832755252239"
	newNumBase10, _ := bsconv.ConvertToDec(numBase62, 62)
	newNumBase16, _ := bsconv.ConvertFromDec(numBase10, 16)
	newNumBase20, _ := bsconv.Conversion(numBase2, 2, 20)
	newRevNumBase62, _ := bsconv.ConvertFromDec(revNumBase10, 62)

	fmt.Println("Base10:", newNumBase10)
	fmt.Println("Base16:", newNumBase16)
	fmt.Println("Base20:", newNumBase20)
	fmt.Println("Base62:", newRevNumBase62)
	fmt.Println("numBase62 == newRevNumBase62:", numBase62 == newRevNumBase62)
}
```
Example output:
```sh
$ go run main.go
Base10: 2493771263131529637640109659423615368150832755252239
Base16: 95be6eec988ecf8
Base20: 47c8ei
Base62: gdguygdgdgdghjjgfdfdgdgdgd767
numBase62 == newRevNumBase62: true
```