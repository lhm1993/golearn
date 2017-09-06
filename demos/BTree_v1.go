// BTree project BTree_v1.go
package BTree_v1

import (
	"fmt"
)

type Elemtype int

type BTree struct {
	Date        Elemtype
	Left, Right *BTree
}

//init func
func InitBTree(root **BTree) {
	*root = new(BTree)
	(**root).Date = Elemtype(1111111)
	(**root).Left = nil
	(**root).Right = nil

}

//add func  return ok or not
func AddNode(root **BTree, date Elemtype) {
	//fmt.Println(*root)
	if *root == nil {
		*root = new(BTree)
		(**root).Date = date
		//fmt.Println((**root).Date)
		return
	}
	//judge insert point
	if date >= (**root).Date {
		AddNode(&((**root).Right), date)
	} else {
		AddNode(&((**root).Left), date)
	}
	return

}
