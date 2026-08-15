package main

import "fmt"

func main() {
	// data type in go and how to use it
	total := 23 // direct infer as int data type
	fmt.Println(total)
	var x int8 = 124
	var y int16 = 32000
	var z int32 = 342342344
	var a int64 = 234324234

	var abc uint = 34234234 // no nagative for unsigned data type

	var onte uint8 = 124    // no nagative for unsigned data type
	var zy uint16 = 3423    // no nagative for unsigned data type
	onerauint32 := 34234234 // no nagative for unsigned data type

	var teston uint64 = 34234234 // no nagative for unsigned data type
	fmt.Println(x, y, z, a, abc, onte, zy, onerauint32, teston)
}
