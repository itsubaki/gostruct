package gostruct

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrFieldNotFound = errors.New("field not found")
	ErrTypeMismatch  = errors.New("type mismatch")
)

type Builder struct {
	field []reflect.StructField
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) AddField(name string, ftype reflect.Type, tag ...string) *Builder {
	field := reflect.StructField{
		Name: name,
		Type: ftype,
	}

	if len(tag) > 0 {
		field.Tag = reflect.StructTag(tag[0])
	}

	b.field = append(b.field, field)
	return b
}

func (b *Builder) String(name string, tag ...string) *Builder {
	return b.AddField(name, reflect.TypeOf(""), tag...)
}

func (b *Builder) Bool(name string, tag ...string) *Builder {
	return b.AddField(name, reflect.TypeOf(true), tag...)
}

func (b *Builder) Int64(name string, tag ...string) *Builder {
	return b.AddField(name, reflect.TypeOf(int64(0)), tag...)
}

func (b *Builder) Float64(name string, tag ...string) *Builder {
	return b.AddField(name, reflect.TypeOf(float64(1.2)), tag...)
}

func (b *Builder) Build() Struct {
	ref := reflect.StructOf(b.field)

	index := make(map[string]int)
	for i := range ref.NumField() {
		index[ref.Field(i).Name] = i
	}

	return Struct{
		internal: ref,
		index:    index,
	}
}

type Struct struct {
	internal reflect.Type
	index    map[string]int
}

func (s *Struct) New() *Instance {
	return &Instance{
		internal: reflect.New(s.internal).Elem(),
		index:    s.index,
	}
}

type Instance struct {
	internal reflect.Value
	index    map[string]int
}

func (i *Instance) Field(name string) (reflect.Value, error) {
	idx, ok := i.index[name]
	if !ok {
		return reflect.Value{}, ErrFieldNotFound
	}

	return i.internal.Field(idx), nil
}

func (i *Instance) SetString(name, value string) error {
	f, err := i.Field(name)
	if err != nil {
		return fmt.Errorf("set string (%s=%s): %w", name, value, err)
	}

	if f.Kind() != reflect.String {
		return fmt.Errorf("field '%s' has type %s: %w", name, f.Kind(), ErrTypeMismatch)
	}

	f.SetString(value)
	return nil
}

func (i *Instance) SetBool(name string, value bool) error {
	f, err := i.Field(name)
	if err != nil {
		return fmt.Errorf("set bool (%s=%v): %w", name, value, err)
	}

	if f.Kind() != reflect.Bool {
		return fmt.Errorf("field '%s' has type %s: %w", name, f.Kind(), ErrTypeMismatch)
	}

	f.SetBool(value)
	return nil
}

func (i *Instance) SetInt64(name string, value int64) error {
	f, err := i.Field(name)
	if err != nil {
		return fmt.Errorf("set int64 (%s=%d): %w", name, value, err)
	}

	if f.Kind() != reflect.Int64 {
		return fmt.Errorf("field '%s' has type %s: %w", name, f.Kind(), ErrTypeMismatch)
	}

	f.SetInt(value)
	return nil
}

func (i *Instance) SetFloat64(name string, value float64) error {
	f, err := i.Field(name)
	if err != nil {
		return fmt.Errorf("set float64 (%s=%f): %w", name, value, err)
	}

	if f.Kind() != reflect.Float64 {
		return fmt.Errorf("field '%s' has type %s: %w", name, f.Kind(), ErrTypeMismatch)
	}

	f.SetFloat(value)
	return nil
}

func (i *Instance) Interface() any {
	return i.internal.Interface()
}

func (i *Instance) Addr() any {
	return i.internal.Addr().Interface()
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
