// Package util provides some helpers for working with GopherJS.
package util

import (
	"time"

	"github.com/neelance/gopherjs/js"
)

func Float64Slice(o js.Object) []float64 {
	d := o.Interface().([]interface{})
	ret := make([]float64, len(d))
	for i, e := range d {
		ret[i] = e.(float64)
	}
	return ret
}

func IntSlice(o js.Object) []int {
	d := o.Interface().([]interface{})
	ret := make([]int, len(d))
	for i, e := range d {
		ret[i] = int(e.(float64))
	}
	return ret
}

func StringSlice(o js.Object) []string {
	d := o.Interface().([]interface{})
	ret := make([]string, len(d))
	for i, e := range d {
		ret[i] = e.(string)
	}
	return ret
}

func Time(o js.Object) time.Time {
	ms := int64(o.Call("getTime").Float())
	s := ms / 1000
	ns := (ms % 1000 * 1e6)
	return time.Unix(int64(s), int64(ns))
}

type Err struct {
	js.Object
	Message string `js:"message"`
	Name    string `js:"name"`
	File    string `js:"fileName"`   // Mozilla extension
	Line    int    `js:"lineNumber"` // Mozilla extension
	Stack   string `js:"stack"`      // Chrome/Microsoft extension
}

func (err Err) Error() string {
	return err.Message
}

func Error(o js.Object) error {
	if o.IsNull() {
		return nil
	}
	return Err{Object: o}
}
