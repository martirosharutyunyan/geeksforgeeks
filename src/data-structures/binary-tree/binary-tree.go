package main

import (
	"fmt"
	"math"
)

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

type Height struct {
	h int
}

func PreorderRecursive(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)

	PreorderRecursive(root.Left)
	PreorderRecursive(root.Right)
}

func PreorderIterative(root *Node) {
	if root == nil {
		return
	}

	stack := []*Node{}

	for root != nil || len(stack) != 0 {
		if root != nil {
			fmt.Println(root.Value)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
}

func InorderRecursive(root *Node) {
	if root == nil {
		return
	}

	InorderRecursive(root.Left)
	fmt.Println(root.Value)
	InorderRecursive(root.Right)
}

func InorderIterative(root *Node) {
	if root == nil {
		return
	}

	stack := []*Node{}

	current := root

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println(current.Value)
		current = current.Right
	}
}

func PostorderRecursive(root *Node) {
	if root == nil {
		return
	}

	PostorderRecursive(root.Left)
	PostorderRecursive(root.Right)

	fmt.Println(root.Value)
}

func PostOrderIterative(root *Node) {
	firstStack, secondStack := []*Node{}, []*Node{}

	if root == nil {
		return
	}
	firstStack = append(firstStack, root)

	var temp *Node
	for len(firstStack) > 0 {
		temp = firstStack[len(firstStack)-1]
		firstStack = firstStack[:len(firstStack)-1]
		secondStack = append(secondStack, temp)

		if temp.Left != nil {
			firstStack = append(firstStack, temp.Left)
		}

		if temp.Right != nil {
			firstStack = append(firstStack, temp.Right)
		}
	}

	for len(secondStack) > 0 {
		temp = secondStack[len(secondStack)-1]
		secondStack = secondStack[:len(secondStack)-1]
		fmt.Println(temp)
	}
}

func Levelorder(root *Node) {
	q := []*Node{root}
	var node *Node

	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		fmt.Println(node.Value)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
}

func Search(root *Node, data interface{}) int {
	q := []*Node{root}
	var node *Node

	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		if node.Value == data {
			return 1
		}

		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return -1
}

func Insert(temp *Node, value interface{}) {
	if temp == nil {
		temp = &Node{Value: value}
		return
	}

	stack := []*Node{temp}

	for len(stack) != 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if temp.Left == nil {
			temp.Left = &Node{Value: value}
			break
		} else {
			stack = append(stack, temp.Left)
		}

		if temp.Right == nil {
			temp.Right = &Node{Value: value}
			break
		} else {
			stack = append(stack, temp.Right)
		}
	}
}

func DeleteDepest(root *Node, delNode *Node) {
	q := []*Node{root}

	var temp *Node

	for len(q) > 0 {
		temp = q[0]
		q = q[1:]

		if temp == delNode {
			temp = nil
			return
		}

		if temp.Right != nil {
			if temp.Right == delNode {
				temp.Right = nil
				return
			}
		} else {
			q = append(q, temp.Right)
		}

		if temp.Left != nil {
			if temp.Left == delNode {
				temp.Left = nil
				return
			}
		} else {
			q = append(q, temp.Left)
		}
	}
}

func Delete(root *Node, value interface{}) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if root.Value == value {
			root = nil
		}
		return
	}

	q := []*Node{root}

	var temp, keyNode *Node

	for len(q) > 0 {
		temp = q[0]
		q = q[1:]

		if temp.Value == value {
			keyNode = temp
		}

		if temp.Left != nil {
			q = append(q, temp.Left)
		}

		if temp.Right != nil {
			q = append(q, temp.Right)
		}
	}

	if keyNode != nil {
		value := temp.Value
		DeleteDepest(root, temp)
		keyNode.Value = value
	}
}

func GetHeight(root *Node) int {
	var x, y int

	if root != nil {
		x = GetHeight(root.Left)
		y = GetHeight(root.Right)

		if x > y {
			return x + 1
		}
		return y + 1
	}

	return 0
}

func DiameterOpt(root *Node, height *Height) int {
	leftHeight := &Height{}
	rightHeight := &Height{}

	if root == nil {
		height.h = 0
		return 0
	}

	leftDiameter := DiameterOpt(root.Left, leftHeight)
	rightDiameter := DiameterOpt(root.Right, rightHeight)
	height.h = int(math.Max(float64(leftDiameter), float64(rightDiameter))) + 1

	return int(math.Max(float64(leftHeight.h)+float64(rightHeight.h), math.Max(float64(leftDiameter), float64(rightDiameter))))
}

func Diameter(root *Node) int {
	height := &Height{}

	return DiameterOpt(root, height)
}

func IsFullBtree(root *Node) bool {
	q := []*Node{root}
	var node *Node

	for len(q) > 0 {
		node = q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}

		if (node.Right != nil && node.Left == nil) || (node.Right == nil && node.Left != nil) {
			return false
		}
	}

	return true
}

func IntToBool(x int) bool {
	if x == 0 {
		return false
	}
	return true
}

func BoolToInt(boolean bool) int {
	if boolean {
		return 1
	}
	return 0
}

func IsBalanced(root *Node, height *int) int {
	leftHeight, rightHeight := 0, 0

	left, right := 0, 0

	if root == nil {
		*height = 0
		return 1
	}

	left = IsBalanced(root.Left, &leftHeight)
	right = IsBalanced(root.Right, &rightHeight)

	if leftHeight > rightHeight {
		*height = leftHeight
	} else {
		*height = rightHeight
	}
	*height += 1

	if math.Abs(float64(leftHeight)-float64(rightHeight)) >= 2 {
		return 0
	}

	return BoolToInt(IntToBool(left) && IntToBool(right))
}

func main() {
	tree := BinaryTree{Root: &Node{Value: 1}}
	node2 := Node{Value: 2}
	node3 := Node{Value: 3}
	node4 := Node{Value: 4}
	node5 := Node{Value: 5}
	node6 := Node{Value: 6}
	node7 := Node{Value: 7}
	tree.Root.Left = &node2
	tree.Root.Right = &node3
	node2.Left = &node4
	node2.Right = &node5
	node3.Left = &node6
	node3.Right = &node7
	//		 1
	//   2      3
	// 4   5  6   7
	// Preorder(tree.Root)
	// Inorder(tree.Root)
	// Levelorder(tree.Root)
	// height := 0
	// fmt.Println(IsBalanced(tree.Root, &height))

}
