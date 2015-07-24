package api

import "fmt"

// http://programmers.stackexchange.com/a/177446
type GSet struct {
	set map[string]bool
}

func NewGSet() *GSet {
	return &GSet{make(map[string]bool)}
}

func (this *GSet) PrintIfNotInSet(key string) {
	if !this.set[key] {
		fmt.Println(key)
		this.set[key] = true
	}
}
