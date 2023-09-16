
package main 

import ( 
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perim() float64
}


type rect struct {
	height, width float64
}

type circle struct {
	radius float64
}

//method area on rect 

func (r rect) area() float64 {
	return r.height * r.width 
}

func (r rect) perim() float64 {
	return 2*(r.height + r.width)
}


func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}


func (c circle) perim() float64 {
	return 2* math.Pi * c.radius
}


//measure area and perimeter of the object
func measure(g geometry) {
	fmt.Println(g)
	fmt.Printf("area := %.3f",g.area())
	fmt.Printf("perim := %.3f",g.perim())
}


func main(){
	
	r := rect{height: 3, width: 8}
	c := circle{radius: 4}

	measure(r)
	measure(c)
}
