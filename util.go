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

func Time(o js.Object) time.Time {
	ms := int64(o.Call("getTime").Float())
	s := ms / 1000
	ns := (ms % 1000 * 1e6)
	return time.Unix(int64(s), int64(ns))
}
