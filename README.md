# gostruct
A Runtime Struct Builder for Go


## Example

```go
func Example() {
	b := gostruct.New()
	b.AddField("Name", reflect.TypeOf(""))
	b.AddField("Age", reflect.TypeOf(int64(0)))
	person := b.Build()

	p := person.New()
	p.SetString("Name", "gopher")
	p.SetInt64("Age", 11)

	fmt.Printf(" %T:  %+v\n", p.Interface(), p.Interface())
	fmt.Printf("%T: %+v\n", p.Addr(), p.Addr())

	// Output:
	//  struct { Name string; Age int64 }:  {Name:gopher Age:11}
	// *struct { Name string; Age int64 }: &{Name:gopher Age:11}
}
```