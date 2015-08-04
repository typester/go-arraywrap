package arraywrap

import (
	"fmt"
	"reflect"
)

type ArrayWrapper interface {
	Len() int
	Index(i int) interface{}
}

type ArrayWrap struct {
	slice interface{}
}

func New(slice interface{}) *ArrayWrap {
	t := reflect.TypeOf(slice)
	if t.Kind() != reflect.Slice {
		panic(fmt.Sprintf("needs slice but passed: %s", t.Kind().String()))
	}

	return &ArrayWrap{slice}
}

func (a *ArrayWrap) Len() int {
	v := reflect.ValueOf(a.slice)
	return v.Len()
}

func (a *ArrayWrap) Index(i int) interface{} {
	// support negative index
	if i < 0 {
		i = a.Len() + i
	}

	v := reflect.ValueOf(a.slice)
	return v.Index(i).Interface()
}
