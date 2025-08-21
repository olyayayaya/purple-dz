package bins

import (
	"time"
)

type Bin struct {
	id string
	private bool
	createdAt time.Time
	name string
}

type BinList struct {
	bins []*Bin
}

func NewBin (id string, private bool, name string) (*Bin) {
	newBin := &Bin{
		id: id,
		private: private,
		createdAt: time.Now(),
		name: name,
	}

	return newBin
}

func NewbinList (bins []*Bin) (*BinList) {
newbinList := &BinList{
	bins: bins,
}

return newbinList
}
