package god

import (
	"fmt"
	"strconv"
)

// Utils is the container of utility dependencies
type Utils struct{}

// NewUtils returns new util
func NewUtils() Utils {
	return Utils{}
}

// Smap wil return new Smap
func (U *Utils) Smap() Smap {
	return NewSmap()
}

// LeftPad will simply pad the string
func (U *Utils) LeftPad(v interface{}, p string) (o string) {

	if i, ok := v.(int); ok {
		o = strconv.Itoa(i)
	} else if s, ok := v.(string); ok {
		o = s
	} else {
		o = fmt.Sprint(v)
	}

	return fmt.Sprintf(p, o)
}

// SortByKeys will sort map by key
func (U *Utils) SortByKeys(
	m map[interface{}]interface{},
	c func(k string, v interface{}) string,
	p string) (
	o string) {

	var t string

	l := U.SliceKeys(m)

	//U.SortSlice(l)

	for _, i := range l {
		fmt.Println("Yeah", i, m[i])
		t = U.LeftPad(i, p)
		o += c(t, m[i])
	}

	return o
}

// SliceKeys will return all keys in map
func (U *Utils) SliceKeys(m map[interface{}]interface{}) (l []interface{}) {
	for k := range m {
		l = append(l, k)
	}
	return
}
