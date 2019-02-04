package god

import (
	"fmt"
	"sort"
	"strconv"
)

// Smap is sorted map
type Smap struct {
	d map[string]interface{}
	p string
}

// NewSmap create new smap
func NewSmap() Smap {
	return Smap{d: make(map[string]interface{}), p: "%s"}
}

// Pad will return the padding for keys
func (s *Smap) Pad() string {
	return s.p
}

// Pad will set the padding for keys
func (s *Smap) SetPad(p string) *Smap {
	s.p = p
	return s
}

// String will add string KV to map
func (s *Smap) String(k string, v interface{}) *Smap {
	s.d[k] = v
	return s
}

// Int will add int KV to map
func (s *Smap) Int(k int, v interface{}) *Smap {
	s.d[strconv.Itoa(k)] = v
	return s
}

// Keys will return a list of all keys
func (s *Smap) Keys() (l []string) {
	for k := range s.d {
		l = append(l, k)
	}
	sort.Strings(l)
	return
}

// SprintLn will simply the map in sorted order
func (s *Smap) SprintLn() (o string) {
	ks := s.Keys()
	for _, v := range ks {
		o += fmt.Sprintf(s.p, v)
		o += fmt.Sprintln(s.d[v])
	}

	return o
}
