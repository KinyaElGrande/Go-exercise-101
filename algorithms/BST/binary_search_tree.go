package bst

import "sync"

type Item interface{}

type TreeNode struct {
	key   int
	value Item
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

// Insert inserts the element with the given key and value in the BST
func (bst *BinarySearchTree) Insert(key int, value Item) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	node := &TreeNode{key, value, nil, nil}
	if bst.rootNode == nil {
		bst.rootNode = node
	} else {
		insertTreeNode(bst.rootNode, node)
	}
}

// insertTreeNode finds the correct place for the Node in the Tree
func insertTreeNode(node, newNode *TreeNode) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertTreeNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertTreeNode(node.right, newNode)
		}
	}
}

// InOrderTraverse visits all the nodes in order
func (bst *BinarySearchTree) InOrderTraverse(fn func(Item)) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.rootNode, fn)
}

// inOrderTraverse traverses the left, the root and the right tree
func inOrderTraverse(node *TreeNode, fn func(Item)) {
	if node != nil {
		inOrderTraverse(node.left, fn)
		fn(node.value)
		inOrderTraverse(node.right, fn)
	}
}

// PreOrderTraverse visits all the tree nodes with preorder traversing
func (bst *BinarySearchTree) PreOrderTraverse(fn func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.rootNode, fn)
}

func preOrderTraverse(node *TreeNode, fn func(Item)) {
	if node != nil {
		fn(node.value)
		preOrderTraverse(node.left, fn)
		preOrderTraverse(node.right, fn)
	}
}

// PostOrderTraverse traverses all the nodes in a post order method
// (left -> right -> current node)
func (bst *BinarySearchTree) PostOrderTraverse(fn func(Item)) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.rootNode, fn)
}

func postOrderTraverse(node *TreeNode, fn func(Item)) {
	if node != nil {
		postOrderTraverse(node.left, fn)
		postOrderTraverse(node.right, fn)
		fn(node.value)
	}
}

// Min returns the item with the minimum value stored in the tree
func (bst *BinarySearchTree) Min() *Item {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	mainNode := bst.rootNode
	if mainNode == nil {
		return nil
	}
	for {
		if mainNode.left == nil {
			return &mainNode.value
		}
		mainNode = mainNode.left
	}
}

// Max returns the item with the maximum value stored in the tree
func (bst *BinarySearchTree) Max() *Item {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	mainNode := bst.rootNode
	if mainNode == nil {
		return nil
	}
	for {
		if mainNode.right == nil {
			return &mainNode.value
		}
		mainNode = mainNode.right
	}
}

// Search returns true if the Item exists in the tree
func (bst *BinarySearchTree) Search(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.rootNode, key)
}

func search(node *TreeNode, key int) bool {
	if node == nil {
		return false
	}
	if key < node.key {
		return search(node.left, key)
	}
	if key > node.key {
		return search(node.right, key)
	}
	return true
}

// Remove removes the item with the passed in key from the tree
func (bst *BinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.rootNode, key)
}

func remove(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return nil
	}

	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}

	leftMostRightSide := node.right
	for {
		if leftMostRightSide != nil && leftMostRightSide.left != nil {
			leftMostRightSide = leftMostRightSide.left
		} else {
			break
		}
	}

	node.key, node.value = leftMostRightSide.key, leftMostRightSide.value
	node.right = remove(node.right, node.key)

	return node
}