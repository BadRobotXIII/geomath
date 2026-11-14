package main

import (
	"fmt"

	"github.com/BadRobotXIII/geomath/geomath"
)

func main() {

	c := geomath.Circle{Diameter: 25.4, Radius: 53.7}
	c.Circumfrence = c.CalcCircum(57)
	fmt.Printf("\nCirciumfrence = %f", c.Circumfrence)

}
