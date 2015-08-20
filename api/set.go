package api

import "fmt"

// http://programmers.stackexchange.com/a/177446
type GSet struct {
	set map[string]bool
}

type SimpleFunc func(addedKey string)

func NewGSet() *GSet {
	return &GSet{make(map[string]bool)}
}

func (this *GSet) Keys() []string {
	keys := make([]string, 0, len(this.set))
	for k := range this.set {
		keys = append(keys, k)
	}
	return keys
}

func (this *GSet) Add(key string) *GSet {
	var stub SimpleFunc = func(_ string) {}
	return this.AddX(key, stub)
}

func (this *GSet) AddX(key string, addedCb SimpleFunc) *GSet {
	if !this.set[key] {
		this.set[key] = true
		addedCb(key)
	}
	return this
}

func (this *GSet) PrintIfNotInSet(key string) {
	var cb SimpleFunc = func(k string) { fmt.Println(key) }
	this.AddX(key, cb)
}
