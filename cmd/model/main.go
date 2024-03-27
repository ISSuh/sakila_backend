package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Hook struct {
}

func (h *Hook) Load(model any) {
	examiner(reflect.TypeOf(model), 0)
}

type Item struct {
	Hook

	value int
}

type BindFunc struct {
	params  []interface{}
	results []interface{}
	fn      interface{}
}

func (bf *BindFunc) Params(params ...interface{}) *BindFunc {
	bf.params = params
	return bf
}

func (bf *BindFunc) Results(results ...interface{}) *BindFunc {
	bf.results = results
	return bf
}

func (bf *BindFunc) Do() {
	// Copy params to slice of reflect values.
	var params []reflect.Value
	for _, p := range bf.params {
		params = append(params, reflect.ValueOf(p))
	}

	// Invoke the function.
	results := reflect.ValueOf(bf.fn).Call(params)

	// Copy the results.
	for i, r := range results {
		reflect.ValueOf(bf.results[i]).Elem().Set(r)
	}
}

// NewBindFunc is a proxy function should run the function and bind the results
// to the result
func NewBindFunc(fn interface{}) *BindFunc {
	return &BindFunc{fn: fn}
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}

func main() {
	item := Item{}
	// examiner(reflect.TypeOf(item), 0)
	item.Load(&Item{})
}
