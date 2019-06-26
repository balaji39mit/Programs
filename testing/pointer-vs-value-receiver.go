package main

import (
	"fmt"

	"github.com/mitchellh/copystructure"
)

type test struct {
	X int
}

//How go works?
//if it is a pointer, change the caller to pointer
//if it is a value, change the caller to value.
func (t test) New() test {
	data, err := copystructure.Copy(t)
	if err == nil {
		if t1, ok := data.(test); ok {
			fmt.Println("works without prob")
			return t1
		} else if t1, ok := data.(*test); ok {
			fmt.Println("works genuinely")
			return *t1
		}
	}
	return t
}

func main() {

	t := test{X: 5}

	t1 := t.New()

	fmt.Println(t1.X)

	t2 := &test{X: 10}

	t3 := t2.New()

	fmt.Println(t3.X)
}
