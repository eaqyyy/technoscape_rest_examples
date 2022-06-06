package main

import (
	"fmt"
	"latian_techno/helpers"
	"strings"
)

// A. ii -> bikin folder helpers, isinya helloName

// Golang packages rules:
// * Write source code in "Program files/go/src" (windows) or "home/go/src" (mac)
// * Main func has to be written "package main"
// * There can be no duplicating func name in the same package
// * A file can directly access another file's func if they're on the same package, else, importing is required.
// * The first letter of a func name must be in lowercase so that it can be imported.
func A_ii() {
	fmt.Println("Hello World")

	helpers.HelloName("Iqi")
}

func B_i() {
	var name string
	name = "Iqi"

	var age int
	age = 23
	
	var weight float32
	weight = 71.5

	// int - just like int32/64
	// int8 (-128 to 127)
	// int16 (-32,768 to 32,767)
	// int32
	// int64
	// uint8/16/32/64 (only positive integers)

	var filmCollections int32
	filmCollections = 1_000

	var isStudent bool
	isStudent = false

	car := "Porsche"
	fuelPrice := 50_000 
	isEuropeanCar := true

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age, " years old")
	fmt.Println("Weight: ", weight, " kilograms")
	fmt.Println("Film collections: ", filmCollections)
	fmt.Println("Is student: ", isStudent)
	fmt.Println("Car: ", car)
	fmt.Println("Is European car: ", isEuropeanCar)
	fmt.Println("Fuel price: ", fuelPrice)
}

func B_ii() {
	var carCount int32
	carCount = 5

	var totalFuelPrice float64
	totalFuelPrice = 55_000.25

	var averagePrice float64
	averagePrice = totalFuelPrice / float64(carCount)

	fmt.Println("average price: ", averagePrice)
}

func C_i() {
	const pi = 3.14
	const dollarConversion float64 = 15000.25

	// its illegal to mutate const data
}

func D() {
	var name string
	var age int32
	var weight float64

	fmt.Println("Please input your name: ")
	fmt.Scan(&name)

	fmt.Println("Please input your age: ")
	fmt.Scan(&age)

	fmt.Println("Please input your weight: ")
	fmt.Scan(&weight)

	fmt.Println("Your name: ", name)
	fmt.Println("Your age: ", age)
	fmt.Println("Your weight: ", weight)

	fmt.Println("Please input your name, age, and weight, separated by spaces")
	fmt.Scan(&name, &age, &weight)
	fmt.Println("Your name: ", name)
	fmt.Println("Your age: ", age)
	fmt.Println("Your weight: ", weight)
}

func E() {
	//+ - * / % ++ --

	x := 12
	y := 4
	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	x++
	fmt.Println("x++ ->", x)

	x--
	fmt.Println("x-- ->", x)
}

func F() {
	//isHuman := true
	//isAnimal := true
	//onlyEatMeat := true
	//hasSuperPower := true
	//
	//isSuperHuman := isHuman && hasSuperPower
	//isCreature := isHuman || isAnimal
	//isSuperAnimal := isAnimal && hasSuperPower
	//isHerbivore := !onlyEatMeat
	//
	//if isHuman {
	//	fmt.Println("a human")
	//	// if this block alrdy executed, remaining else(s) will be ignored
	//} else if isSuperHuman {
	//	fmt.Println("not a human")
	//}

	x := 15
	y := 16

	if x < y {
		fmt.Println("x is less than y")
	} else if x > y {
		fmt.Println("x is greater than y")
	} else if x >= y {
		fmt.Println("x is greater or equal than y")
	} else if x == 100 {
		fmt.Println("x is one hundred")
	} else {
		fmt.Println("no condition met")
	}
}

func G_H() {
	var array1 [5]int
	// all items will be set to 0
	fmt.Println(array1)

	array1[0] = 15
	fmt.Println(array1[0])

	// array[6] is impossible because out of bound
	slice := make([]int, 5)

	// slice[5] is impossible because array starts from 0, out of range
	slice[4] = 14
	fmt.Println(slice)

	slice = append(slice, 16)
	fmt.Println(slice[5])
}

func I() {
	firstName := "rayhan"
	lastName := "shidqi"
	completeName := firstName + " " + lastName
	fmt.Println(completeName)

	if completeName == "rayhan shidqi" {
		fmt.Println("you are admin")
	} else if completeName == "teman sejati" {
		fmt.Println("you are a friend")
	} else if completeName == "RAYHAN SHIDQI" {
		fmt.Println("all uppercase user")
	}

	// replaceAll, toUpper, toLower, length, contains -> use goland autocomplete
	sentence := "technoscape is held by binus computer club to understand computer language"
	fmt.Println(strings.ToUpper(sentence))

	fmt.Println(len(sentence))

	// starting from 0, and 8 is excluded
	fmt.Println(sentence[0:8])
	fmt.Println(sentence[:8])
	fmt.Println(sentence[8:]) // useful to get numeric ID
}

func J_K() {
	// type "fori" in goland for autocomplete
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while
	j := 0
	for j < 5 {
		j++
		fmt.Println(j)
	}

	maxLoop := 10
	counter := 0
	for {
		if counter == 3 {
			fmt.Println("skipping 3")
			counter++
			continue
		}

		if counter == maxLoop {
			fmt.Println("counter has reached max... stopping loop")
			break
		}

		fmt.Println("counter: ", counter)
		counter++
	}

	sentence := "technoscape binus is the best"
	for k, v := range sentence {
		fmt.Println("Index: ", k, " char: ", v)
	}
	// use _ if unused variable
}

func L() {
	type Book struct {
		Title     string
		Pages     int32
		Author    string
		IsFiction bool
	}

	bookOne := Book{
		Title:     "Harry Potter",
		Pages:     500,
		Author:    "JK Rowling",
		IsFiction: true,
	}
	fmt.Println(bookOne)

	bookTwo := Book{}
	bookTwo.Pages = 25
	// will be empty string/false by default
	fmt.Println(bookTwo)
}

func Example() {
	fmt.Println("this is example")
}

func SayMyName(name string) {
	fmt.Println("Hello, my name is, ", name)
}

func PowerToTwo(number int64) int64 {
	return number * number
}

func GeneratePatientNumber(id string, age int) (string, bool) {
	var generatedId string
	var isPriority bool

	if age > 65 {
		generatedId = "ELDER-" + id
		isPriority = true
	} else {
		generatedId = "NORMAL-" + id
		isPriority = false
	}

	// can be subtituted by early return
	return generatedId, isPriority
}

func M() {
	Example()

	SayMyName("Iqi")

	quadraticNumber := PowerToTwo(5)

	fmt.Println(quadraticNumber)

	generatedId, priority := GeneratePatientNumber("12345", 55)

	fmt.Println(generatedId, " ", priority)
}

type Snack struct {
	Name     string
	Quantity int32
	Price    float64
}

// pass by value -> creates a copy of the var
// pass by reference -> directly access the original object
func buySnack(snack *Snack) {
	snack.Quantity--
}

func N() {
	oreo := Snack{
		Name:     "Oreo Choco",
		Quantity: 10,
		Price:    5000,
	}

	buySnack(&oreo)

	fmt.Println(oreo)
}