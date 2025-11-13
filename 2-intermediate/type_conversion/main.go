package main

import "fmt"

func main() {

	var a int = 32
	b := int32(a)
	c := float64(b)
	// d :=  bool(true)
	e := 3.14
	f := int(e)
	fmt.Println(f, c)

	// Type(value)
	g := "Hello @ こんにちは 😋 привет"
	var h []byte
	h = []byte(g)
	fmt.Println(h)
	fmt.Println(string(h[8]))
	fmt.Println(g)
	rn := []rune(g)
	fmt.Println(string(rn[14]), ":", rn[14])
	i := []byte{255, 120, 73}
	j := string(i)
	fmt.Println(j)
}
