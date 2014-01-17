// Package util provides some helpers for working with GopherJS.
package util

import "github.com/neelance/gopherjs/js"

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
