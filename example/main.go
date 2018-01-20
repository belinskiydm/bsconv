package main

import (
	"fmt"

	"github.com/belinskiydm/baseconverter"
)

func main() {
	numBase62 := "gdguyd767"
	numBase10 := "674386482479820024"
	numBase2 := "110101011110101110101010"
	newNumBase10, _ := baseconverter.ConvertToDec(numBase62, 62)
	newNumBase16, _ := baseconverter.ConvertFromDec(numBase10, 16)
	newNumBase20, _ := baseconverter.Conversion(numBase2, 2, 20)

	fmt.Println("Base10:", newNumBase10)
	fmt.Println("Base16:", newNumBase16)
	fmt.Println("Base20:", newNumBase20)

}
