package main

import (
	"time"
)

type bin struct {
	id string
	private bool
	createdAt time.Time
	name string
}

type binList struct {
	bins string
}

func newBin (id string, private bool, name string) (*bin) {
	newBin := &bin{
		id: id,
		private: private,
		createdAt: time.Now(),
		name: name,
	}

	return newBin
}

func newbinList (bins string) (*binList) {
newbinList := &binList{
	bins: bins,
}

return newbinList
}

func main() {

}