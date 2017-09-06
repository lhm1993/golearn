// BTree project BTree.go
package BTree

type Elemtype int

type BTree struct {
	Date        Elemtype
	Left, Right *BTree
}

//init func
func InitBTree(root *BTree) *BTree {
	root = new(BTree)
	root.Date = Elemtype(1111111)
	root.Left = nil
	root.Right = nil
	return root
}

//add func  return ok or not
func AddNode(root *BTree, date Elemtype) *BTree {
	if root == nil {
		root = new(BTree)
		root.Date = date
		return root
	}
	//judge insert point
	if date >= root.Date {
		root.Right = AddNode(root.Right, date)
	} else {
		root.Left = AddNode(root.Left, date)
	}
	return root
}
