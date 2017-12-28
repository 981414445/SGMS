package util

import (
	"encoding/json"
)

// type Set interface {
// }
type SetElement interface {
	// CompareTo(e SetElement) int
	HashCode() int64
}
type HashSet struct {
	Items map[int64]SetElement
}

func NewHashSet() *HashSet {
	return &HashSet{Items: make(map[int64]SetElement)}
}

func (this *HashSet) Contains(e SetElement) bool {
	if _, ok := this.Items[e.HashCode()]; ok {
		return true
	}
	return false
}
func (this *HashSet) Add(e SetElement) {
	this.Items[e.HashCode()] = e
}
func (this *HashSet) Adds(es ...SetElement) {
	for _, e := range es {
		this.Add(e)
	}
}
func (this *HashSet) AddAll(set *HashSet) {
	for _, e := range set.Items {
		this.Add(e)
	}
}

func (this *HashSet) Remove(e SetElement) {
	delete(this.Items, e.HashCode())
}
func (this *HashSet) RemoveAll(set *HashSet) {
	set.Each(func(e SetElement) {
		this.Remove(e)
	})
}
func (this *HashSet) Clear() {
	for k := range this.Items {
		delete(this.Items, k)
	}
}

func (this *HashSet) Size() int {
	return len(this.Items)
}

func (this *HashSet) Empty() bool {
	return this.Size() == 0
}
func (this *HashSet) Values() []SetElement {
	es := make([]SetElement, this.Size())
	i := 0
	for _, e := range this.Items {
		es[i] = e
		i++
	}
	return es
}
func (this *HashSet) Each(cb func(e SetElement)) {
	for _, e := range this.Items {
		cb(e)
	}
}

func (this *HashSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.Values())
}

func (this *HashSet) HashCode() int64 {
	return GetPointer(this)
}
