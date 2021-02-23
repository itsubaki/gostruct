package gostruct_test

import (
	"fmt"
	"reflect"

	"github.com/itsubaki/gostruct"
)

func Example() {
	person := gostruct.New().
		AddField("Name", reflect.TypeOf("")).
		AddField("Age", reflect.TypeOf(int64(0))).
		Build()

	p := person.New()
	p.SetString("Name", "gopher")
	p.SetInt64("Age", 8)

	fmt.Printf(" %T:  %+v\n", p.Interface(), p.Interface())
	fmt.Printf("%T: %+v\n", p.Addr(), p.Addr())

	// Output:
	//  struct { Name string; Age int64 }:  {Name:gopher Age:8}
	// *struct { Name string; Age int64 }: &{Name:gopher Age:8}
}
