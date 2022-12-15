package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "!  ,ab c! d ! e "

	fmt.Println(s)

	s = strings.Trim(s, " ,.!")

	fmt.Println(s)
}
