package main

import (
	"fmt"
	"strings"
)

func main() {
	splitString()
	fmt.Println("========================")
	containsCheckString()
	fmt.Println("========================")
	subStringCount()
	fmt.Println("========================")
	subStringIndex()
	fmt.Println("========================")
	stringRepeat()
}

func splitString() {
	str := "a,b,c"
	var split1 = strings.Split(str, ",")
	for _, v := range split1 {
		fmt.Println(" v :  ", v)
	}
	for _, v := range strings.SplitAfter(str, ",") {
		fmt.Println(" v2 :  ", v)
	}

	split3 := strings.SplitAfterN(str, ",", 2)
	for i := 0; i < len(split3); i++ {
		fmt.Println(" split3[i] :  ", split3[i])
	}
}

func containsCheckString() {
	str1 := "seafood"
	fmt.Println(" seafood contains food :", strings.Contains(str1, "food"))
	fmt.Println(" seafood contains f :", strings.Contains(str1, "fe"))
	fmt.Println(" seafood contains \"\" :", strings.Contains(str1, ""))

	fmt.Println(" seafood ContainsAny food :", strings.ContainsAny(str1, "food"))
	fmt.Println(" seafood ContainsAny f :", strings.ContainsAny(str1, "fe"))
	fmt.Println(" seafood ContainsAny \"\" :", strings.ContainsAny(str1, ""))
}
func subStringCount() {
	str1 := "cheese"
	fmt.Println("No of e in cheese:", strings.Count(str1, "e"))
	fmt.Println("No of \"\" in five:", strings.Count("five", ""))
}

func subStringIndex() {
	str1 := "chicken"
	fmt.Println("Index of ken in chicken:", strings.Index(str1, "ken")) //4
	fmt.Println("Index of ck in chicken:", strings.Index(str1, "ck"))   //3
	fmt.Println("Index of cl in chicken:", strings.Index(str1, "cl"))   //-1
	fmt.Println("Index of lc in chicken:", strings.Index(str1, "lc"))   //-1
	fmt.Println("Index of dmr in chicken:", strings.Index(str1, "dmr")) //-1

	fmt.Println("IndexAny of aei in chicken:", strings.IndexAny(str1, "aei")) //2
	fmt.Println("IndexAny of aei in crwt:", strings.IndexAny("crwt", "aei"))  //-1

	fmt.Println("IndexByte of 'g' in golang:", strings.IndexByte("golang", 'g')) //0
	fmt.Println("IndexBtye of 'a' in golang:", strings.IndexByte("golang", 'a')) //3
	fmt.Println("IndexBtye of 'x' in golang:", strings.IndexByte("golang", 'x')) //-1
}
func stringRepeat() {
	str1 := "na"
	fmt.Println("ba" + strings.Repeat(str1, 2)) //banana

	fmt.Println(strings.Repeat(str1, -1)) //Runtime error. panic: strings: negative Repeat count

}
