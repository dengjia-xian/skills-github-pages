package main

//二叉搜索树
import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func (root *Node) MakeEmpty() {
	if root != nil {
		root.left.MakeEmpty()
		root.right.MakeEmpty()
		root = nil
	}

}
func (root *Node) Find(data int) bool {
	if root == nil {
		return false
	}
	if root.data < data {
		return root.right.Find(data)
	} else if root.data > data {
		return root.left.Find(data)
	} else if root.data == data {
		return true
	}
	return false
}
func (root *Node) FindMin() (int, bool, *Node) {
	if root == nil {
		return 0, false, nil
	}
	if root.left == nil {
		return root.data, true, root
	} else {
		return root.left.FindMin()
	}
}
func (root *Node) FindMax() (int, bool, *Node) {
	if root == nil {
		return 0, false, nil
	}
	if root.right == nil {
		return root.data, true, root
	} else {
		return root.right.FindMax()
	}
}
func (root *Node) Insert(data int) bool {
	p := &Node{
		data: data,
		left: nil, right: nil,
	}

	if root == nil {
		root = p
		return true
	}
	if root.data < data {
		if root.right == nil {
			root.right = p
			return true
		} else {
			return root.right.Insert(data)
		}
	}
	if root.data > data {
		if root.left == nil {
			root.left = p
			return true
		} else {
			return root.left.Insert(data)
		}
	}
	return false
}

func (root *Node) Delete(data int) bool {
	if root == nil {
		return false
	}
	if root.data < data {
		if root.right == nil {
			return false
		} else {
			if root.right.data == data {
				if root.right.left == nil && root.right.right != nil {
					root.right = root.right.right
					return true
				} else if root.right.right == nil && root.right.left != nil {
					root.right = root.right.left
					return true
				} else if root.right.right == nil && root.right.left == nil {
					root.right = nil
					return true
				} else {
					k, _, p := root.right.FindMax()
					root.right.data = k
					p.right = nil
					p = nil
					return true
				}
			}
			return root.right.Delete(data)
		}
	} else {
		if root.left == nil {
			return false
		} else {
			if root.left.data == data {
				if root.left.left == nil && root.left.right != nil {
					root.left = root.left.right
					return true
				} else if root.left.right == nil && root.left.left != nil {
					root.left = root.left.left
					return true
				} else if root.left.right == nil && root.left.left == nil {
					root.left = nil
					return true
				} else {
					k, _, p := root.left.FindMax()
					root.left.data = k
					p.left = nil
					p = nil
					return true
				}
			}
			return root.left.Delete(data)
		}
	}

}
func Traversal(root *Node) {
	if root != nil {
		Traversal(root.left)
		fmt.Printf("%d  ", root.data)
		Traversal(root.right)
	}
}
func main() {
	BST := &Node{}
	BST.MakeEmpty()
	BST.Insert(3)
	BST.Insert(1)
	BST.Insert(2)
	BST.Insert(4)
	BST.Insert(0)
	Traversal(BST)
	k, ok, _ := BST.FindMax()
	if ok {
		fmt.Println()
		fmt.Println("最大值是:", k)
	} else {
		fmt.Println("最大值不存在")
	}
	k, ok, _ = BST.FindMin()
	if ok {
		fmt.Println()
		fmt.Println("最小值是:", k)
	} else {
		fmt.Println("最小值不存在")
	}
	BST.Delete(2)
	Traversal(BST)
	ok = BST.Find(2)
	if ok {
		fmt.Printf("\n数值2存在\n")
	} else {
		fmt.Printf("\n数值2不存在\n")
	}
}
