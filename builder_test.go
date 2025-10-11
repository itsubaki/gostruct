package gostruct_test

import (
	"encoding/json"
	"fmt"

	"github.com/itsubaki/gostruct"
)

func Example() {
	person := gostruct.New().
		String("Name").
		Int64("Age").
		Build()

	p := person.New()
	p.SetString("Name", "gopher")
	p.SetInt64("Age", 11)

	fmt.Printf(" %T:  %+v\n", p.Interface(), p.Interface())
	fmt.Printf("%T: %+v\n", p.Addr(), p.Addr())

	// Output:
	//  struct { Name string; Age int64 }:  {Name:gopher Age:11}
	// *struct { Name string; Age int64 }: &{Name:gopher Age:11}
}

func Example_tag() {
	person := gostruct.New().
		String("Name", `json:"name"`).
		Int64("Age", `json:"age"`).
		Build()

	p := person.New()
	p.SetString("Name", "gopher")
	p.SetInt64("Age", 11)

	data, err := json.Marshal(p.Interface())
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// Output:
	// {"name":"gopher","age":11}
}
