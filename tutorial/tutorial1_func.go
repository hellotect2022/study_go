package tutorial

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string
}

//go 에는 constructor 가 없다 (물로 class 도 없음)
//javascript -> constructor()
//python -> __init__
//java -> 클래스 이름 다시 적음()

func StructTest() {
	favFood := []string{"kimchi", "ramen"}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico)
}

func MapTest() {
	//map[key]value
	nico := map[string]int{"name": 1, "age": 12}

	for key, value := range nico {
		fmt.Println(key, value)
	}
	fmt.Println(nico)
}

func ArrayTest() {
	names := []string{"nico", "tes1", "test2"}
	names = append(names, "test3")
	fmt.Println(names)
}

func PointerTest() {
	a := 2
	b := a
	a = 10
	c := &a
	*c = 30
	fmt.Println(a, b, *c, &a, &b, &c)
}

func CanIDrink(age int) bool {

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false

	// go 에서는 for 문안에서 변수 선언이 가능
	/*
		if koreanAge := age + 2; koreanAge < 18 {
			return false
		}
		return true
	*/
}

func SuperAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	total := 0
	for _, number := range numbers {
		// go 에서 _ 는 변수 할당 무시 기능
		total += number
	}
	return total
}

func RepeatMe(words ...string) {
	fmt.Println(words)
}

func LenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func NakedReturnlenAndUpper(name string) (length int, uppercase string) {

	length = len(name)
	defer fmt.Println("I'm done (defer 는 method return 후에 실행되는 코드임)")

	uppercase = strings.ToUpper(name)
	return
}

func Mulitply(a, b int) int {
	return a * b
}
