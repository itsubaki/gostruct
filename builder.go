package gostruct

import "reflect"

type Builder struct {
	field []reflect.StructField
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) AddField(name string, ftype reflect.Type) *Builder {
	b.field = append(
		b.field,
		reflect.StructField{
			Name: name,
			Type: ftype,
		})

	return b
}

func (b *Builder) Build() Struct {
	strct := reflect.StructOf(b.field)

	index := make(map[string]int)
	for i := 0; i < strct.NumField(); i++ {
		index[strct.Field(i).Name] = i
	}

	return Struct{strct, index}
}

type Struct struct {
	strct reflect.Type
	index map[string]int
}

func (s *Struct) New() *Instance {
	instance := reflect.New(s.strct).Elem()
	return &Instance{instance, s.index}
}

type Instance struct {
	internal reflect.Value
	index    map[string]int
}

func (i *Instance) Field(name string) reflect.Value {
	return i.internal.Field(i.index[name])
}

func (i *Instance) SetString(name, value string) {
	i.Field(name).SetString(value)
}

func (i *Instance) SetBool(name string, value bool) {
	i.Field(name).SetBool(value)
}

func (i *Instance) SetInt64(name string, value int64) {
	i.Field(name).SetInt(value)
}

func (i *Instance) SetFloat64(name string, value float64) {
	i.Field(name).SetFloat(value)
}

func (i *Instance) Interface() interface{} {
	return i.internal.Interface()
}

func (i *Instance) Addr() interface{} {
	return i.internal.Addr().Interface()
}
