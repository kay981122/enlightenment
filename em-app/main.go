package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//core.RunWindowsServer()
	a := "329.6333"
	b, _ := strconv.ParseFloat(a, 64)
	c := math.Floor(b)
	fmt.Println(c)
}
