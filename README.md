# About baseconverter
Go package to convert any number from an arbitrary base to another arbitrary base. From binary up to base62. Not only from base2 to base36 like in standard library. 

The package also contains two additional functions to convert from an arbitrary base to decimal and from decimal to an arbitrary base.

All functions get an return numbers in the format of string. 

## Installation ##

Install the package via the following:

    go get -u github.com/belinskiydm/baseconverter

## More Details ##

An example of the main.go file:
```go
// example/main.go
package main

import (
	"fmt"

	"github.com/belinskiydm/baseconverter"
)

func main() {
	numBase62 := "gdguydg6767"
	numBase10 := "674386482479820024"
	numBase2 := "110101011110101110101010"
	newNumBase10, _ := baseconverter.ConvertToDec(numBase62, 62)
	newNumBase16, _ := baseconverter.ConvertFromDec(numBase10, 16)
	newNumBase20, _ := baseconverter.Conversion(numBase2, 2, 20)

	fmt.Println("Base10:", newNumBase10)
	fmt.Println("Base16:", newNumBase16)
	fmt.Println("Base20:", newNumBase20)

}
```
Example output:
```sh
$ go run main.go
Base10: 3540159472514319
Base16: 95be6eec988ecf8
Base20: 47c8ei
```




