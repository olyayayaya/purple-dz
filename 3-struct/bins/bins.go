package bins

import (
	"time"
)

type Db interface {
	Read()([]byte, error)
	CheckFileExtension()(bool)
}

type Bin struct {
	Id string `json:"id"`
	Private bool `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name string `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BinListWithDb struct {
	BinList
	db Db
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

// func NewbinList (bins []*Bin) (*BinList) {

// newbinList := &BinList{
// 	Bins: bins,
// 	UpdatedAt: time.Now(),
// }

// return newbinList
// }

func NewbinList (db Db) (*BinListWithDb) {

newbinList := &BinListWithDb{
	BinList: BinList{
		Bins: []Bin{},
		UpdatedAt: time.Now(),
	},
	db: db,
}

return newbinList
}


