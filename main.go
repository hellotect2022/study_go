package main

import (
	"fmt"
	"log"
	"newProj/accounts"
	"newProj/channel"
	"newProj/mydict"
	"newProj/tutorial"
)

func main() {
	//goroutineTest()
	urlChecker()
	// dictionary()
	// account()
	// test()
}

func goroutineTest() {
	// go channel
	c := make(chan string)

	// go routine 은 main 프로세스가 끝나면 전부 끝나버린다.
	people := [6]string{"nico", "dhhan", "1", "2", "3", "4"}
	for _, person := range people {
		go channel.IsSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

func urlChecker() {
	channel.UrlChecker()
}

func dictionary() {
	//2.
	dictionary := mydict.Dictionary{"first": "first word"}
	word := "hello"
	dictionary.Add(word, "first")
	err := dictionary.Update("word", "second")
	if err != nil {
		fmt.Println(err)
	}
	err = dictionary.Delete(word)
	if err != nil {
		fmt.Println(err)
	}
	definition, err := dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	fmt.Println(dictionary)
}

func account() {
	// accounts 예제
	account := accounts.NewAccount("dhhan")
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(10)
	if err != nil {
		log.Fatalln(err)
	}
	account.ChangeOwner("FUCKER")
	fmt.Println(account)

}

func test() {
	totalLength, upperName := tutorial.NakedReturnlenAndUpper("NiCO")
	fmt.Println(totalLength, upperName)
	total := tutorial.SuperAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)

	fmt.Println(tutorial.CanIDrink(16))
	tutorial.PointerTest()
	tutorial.ArrayTest()
	tutorial.MapTest()
	tutorial.StructTest()
}
