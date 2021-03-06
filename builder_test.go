package gostruct_test

import (
	"fmt"

	"github.com/itsubaki/gostruct"
)

func Example() {
	person := gostruct.New().
		AddString("Name").
		AddInt64("Age").
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
