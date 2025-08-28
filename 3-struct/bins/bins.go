package bins

import (
	"time"
)

type Bin struct {
	Id string `json:"id"`
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name string `json:"name"`
}

type BinList struct {
	Bins []*Bin `json:"bins"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewBin (id string, private bool, name string) (*Bin) {
	newBin := &Bin{
		Id: id,
		Private: private,
		CreatedAt: time.Now(),
		Name: name,
	}
 
	return newBin
}

func NewbinList (bins []*Bin) (*BinList) {
newbinList := &BinList{
	Bins: bins,
	UpdatedAt: time.Now(),
}

return newbinList
}


