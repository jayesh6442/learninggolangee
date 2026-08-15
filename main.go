// package main
//
// import "fmt"
//
//	func main() {
//		// data type in go and how to use it
//		total := 23 // direct infer as int data type
//		fmt.Println(total)
//		var x int8 = 124
//		var y int16 = 32000
//		var z int32 = 342342344
//		var a int64 = 234324234
//
//		var abc uint = 34234234 // no nagative for unsigned data type
//
//		var onte uint8 = 124         // no nagative for unsigned data type
//		var zy uint16 = 3423         // no nagative for unsigned data type
//		var testone int32 = 34234234 // no nagative for unsigned data type
//		var teston uint64 = 34234234 // no nagative for unsigned data type
//
//		fmt.Println(x, y, z, a, abc, onte, zy, testone, teston)
//
//		var testfloat float32 = 3.55
//		var testfloattwo float64 = 4.23
//		var boolne bool = false
//
//		var str string = " hi there "
//		fmt.Println(testfloat, testfloattwo, boolne, str)
//	}
package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var i int
	var i32 int32
	var f32 float32
	var b bool
	var s string

	fmt.Println("int     ", unsafe.Sizeof(i))   // 8 on amd64/arm64
	fmt.Println("int32   ", unsafe.Sizeof(i32)) // 4
	fmt.Println("float32 ", unsafe.Sizeof(f32)) // 4
	fmt.Println("bool    ", unsafe.Sizeof(b))   // 1
	fmt.Println("string  ", unsafe.Sizeof(s))   // 16 ← the HEADER, not the text
	fmt.Println("IntSize ", strconv.IntSize)    // 64
}
