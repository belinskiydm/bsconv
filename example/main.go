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
