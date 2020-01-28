package types

import "strings"

type OrderByKeyValueName []KeyValue

func (a OrderByKeyValueName) Len() int { return len(a) }
func (a OrderByKeyValueName) Less(i int, j int) bool { return 0 <= strings.Compare(a[i].Name, a[j].Name) }
func (a OrderByKeyValueName) Swap(i int, j int) { a[i], a[j] = a[j], a[i] }

// service representation
type Service struct {
	Name string
	Labels []KeyValue
}

// key value pair
type KeyValue struct {
	Name string
	Value string
}