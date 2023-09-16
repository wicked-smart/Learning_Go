package main

import "fmt"

type rect struct {
	height int
	width int
}

//method area with reciever type *rect
func (r *rect) area() int {
		return r.height * r.width
	}

//method perim with reciever type rect
func (r rect) perim() int {
	return 2*(r.height  + r.width )
}

func main(){

	r1 := rect{height: 23, width: 45}
	
	r2 := rect{height: 20, width: 50}

	fmt.Println("r1 area := %d", r1.area())
	fmt.Println("r2 area := %d", r2.area())

	//permimters
	fmt.Println("r1 perimeter := ", r1.perim())
	fmt.Println("r2 perimeter := ", r2.perim())


}
