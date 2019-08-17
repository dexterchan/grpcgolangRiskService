package hello

import (
	"fmt"

	"rsc.io/quote"
)

type ABC struct {
	id   int
	name string
}

func Hello() string {
	m := ABC{
		id:   1,
		name: "peter",
	}
	fmt.Println(m)
	return quote.Hello()
}
