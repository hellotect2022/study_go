package main

import (
	"fmt"
	"newProj/mydict"
)

func main() {
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

	/* // accounts 예제
	account := accounts.NewAccount("dhhan")
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(10)
	if err != nil {
		log.Fatalln(err)
	}
	account.ChangeOwner("FUCKER")
	fmt.Println(account)
	*/

	//fmt.Println(mulitply(2, 3))
	//totalLength, upperName := lenAndUpper("dhhan")
	//totalLength1, _ := lenAndUpper("john")
	//fmt.Println(totalLength, upperName, totalLength1)
	//tutorial.RepeatMe("nico", "lynn", "dal", "marl")
	//totalLength, upperName := tutorial.NakedReturnlenAndUpper("NiCO")
	//fmt.Println(totalLength, upperName)
	//total := tutorial.SuperAdd(1, 2, 3, 4, 5, 6)
	//fmt.Println(total)

	// fmt.Println(tutorial.CanIDrink(16))
	// tutorial.PointerTest()
	//tutorial.ArrayTest()
	// tutorial.MapTest()
	// tutorial.StructTest()
}
