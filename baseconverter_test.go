package baseconvert

import (
	"fmt"
	"testing"
)

func TestConversion(t *testing.T) {
	num, _ := Conversion("0", 2, 10)
	if num != "0" {
		t.Error("Expected 0,", "got ", num)
	}

	num, _ = Conversion("56f", 16, 2)
	if num != "10101101111" {
		t.Error("Expected 10101101111,", "got ", num)
	}

	num, _ = Conversion("-58586655672", 10, 62)
	if num != "-11WTDte" {
		t.Error("Expected -11WTDte,", "got ", num)
	}

	num, _ = Conversion("11WTDte", 62, 2)
	if num != "110110100100000010010110101110111000" {
		t.Error("Expected 110110100100000010010110101110111000,", "got ", num)
	}

}

func TestConvertToDec(t *testing.T) {
	num := ConvertToDec([]rune("453fa"), 16)
	if num != 283642 {
		t.Error("Expected 0,", "got ", num)
	}
}

func TestConvertFromDec(t *testing.T) {
	num := ConvertFromDec(6455537849492, 62)
	if num != "1PEvwVPC" {
		t.Error("Expected 0,", "got ", num)
	}
}

func ExampleConversion() {
	fmt.Println(Conversion("11WTDte", 62, 2))
	// Output:
	// 110110100100000010010110101110111000
}

func ExampleConvertToDec() {
	fmt.Println(ConvertToDec([]rune("453fa"), 16))
	// Output:
	// 283642
}

func ExampleConvertFromDec() {
	fmt.Println(ConvertFromDec(6455537849492, 62))
	// Output:
	// 1PEvwVPC
}

func BenchmarkConversion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Conversion("11WTDtejxjvosjsshs3455", 62, 2)
	}
}
