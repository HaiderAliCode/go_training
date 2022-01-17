package main

import "fmt"

//defining an interface
type Shape interface {
	IsCompanyISOCert() bool
	IsEmploymentValid() bool
}

type Rect struct {
	Width  int
	height int
}

func (r Rect) IsCompanyISOCert() bool {
	return true
}

func (r Rect) IsEmploymentValid() bool {
	return true
}

//empty interface unlimited parameters
func function1(i ...interface{}) {
	fmt.Println(i)
}

type Object interface {
	getvolume() int
}

type Shape1 interface {
	getsize() int
}

type Cube struct {
	weight int
}

func (c Cube) getsize() int {
	return c.weight
}

func (c Cube) getvolume() int {
	return c.weight
}

type skin interface {
	getskin() string
}

//type switching
func whatami(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i am a string")
	case int:
		fmt.Println("i am a int")
	default:
		fmt.Println("i am neither int nor string")
	}
}

type shape2 interface {
	Area() int
	Color() int
}

type Cube1 struct {
	width  int
	height int
}

func (c *Cube1) Area() int {
	return 1
}

func (c Cube1) Color() int {
	return 2
}

func main() {
	var r Shape
	r = Rect{31, 32}
	fmt.Println(r.IsCompanyISOCert())
	fmt.Println(r.IsEmploymentValid())

	function1("hello", "friend")

	cb := Cube{1}

	var s Shape1 = cb
	var o Object = cb
	fmt.Println(s.getsize())
	fmt.Println(o.getvolume())

	//getting dynamic type of o and calling other function which is not in object
	p := o.(Cube)
	fmt.Println(p.getsize())

	// since dynamic type does not implement skin interface thus this is throwing an error
	val, ok := o.(skin)
	fmt.Println(val, ok)

	//type switching
	whatami("hello")
	whatami(1)
	whatami(2.3)

	c := Cube1{1, 2}
	var s1 shape2 = &c
	area := s1.Area()
	color := s1.Color()
	fmt.Println(area, color)

	//interface comparison can be done
	//If these interfaces are not nil, then their dynamic types (the type of their concrete values) should be the same and the concrete values should be equal.

	/*
		Use of interfaces
		We have learned interfaces and we saw they can take different forms. That’s the definition of polymorphism.
		Interfaces are very useful in case of functions and methods where you need argument of dynamic types, like Println function which accepts any type of values.
		When multiple types implement the same interface, it becomes easy to work with them. Hence whenever we can use interfaces, we should.
	*/
}
