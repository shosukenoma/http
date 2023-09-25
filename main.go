package main

import (
	"fmt"
)

// type logWriter struct{}

// func main() {
// 	resp, err := http.Get("http://google.com")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		os.Exit(1)
// 	}

// 	// bs := make([]byte, 99999)
// 	// // create a slice of type byte with 99999 empty elements.
// 	// // this is because read() cannot resize []byte if it becomes full.
// 	// resp.Body.Read(bs)

// 	// fmt.Println(string(bs))

// 	// io.Copy(os.Stdout, resp.Body)
// 	// // os.Stdout implements the Writer interface
// 	// // resp.Body implements the Reader interface

// 	lw := logWriter{}
// 	io.Copy(lw, resp.Body)
// }

// func (logWriter) Write(bs []byte) (int, error) {
// 	fmt.Println(string(bs))
// 	fmt.Println("Finished writing", len(bs), "bytes")

// 	return len(bs), nil
// }

type triangle struct{
	height float64
	base float64
}
type square struct{
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{
		height: 2.0,
		base: 3.0,
	}

	sqr := square{
		sideLength: 4.0,
	}

	printArea(tri)
	printArea(sqr)
}

func printArea(s shape) {
	fmt.Println("The area of this shape is", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}