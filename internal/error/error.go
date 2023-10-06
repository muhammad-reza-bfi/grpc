package terminal

import "fmt"

func CheckErr(err error) {
	if err != nil {
		fmt.Println("got error", err)
		panic(err)
	}
}
