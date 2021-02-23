package gostruct_test

import (
	"fmt"
	"reflect"

	"github.com/itsubaki/gostruct"
)

func Example() {
	b := gostruct.New()
	b.AddField("Name", reflect.TypeOf(""))
	b.AddField("Age", reflect.TypeOf(int64(0)))
	person := b.Build()

	i := person.New()
	i.SetString("Name", "gopher")
	i.SetInt64("Age", 8)

	fmt.Printf(" %T:  %+v\n", i.Interface(), i.Interface())
	fmt.Printf("%T: %+v\n", i.Addr(), i.Addr())

	// Output:
	//  struct { Name string; Age int64 }:  {Name:gopher Age:8}
	// *struct { Name string; Age int64 }: &{Name:gopher Age:8}
}
