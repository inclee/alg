package alg

type Node struct {
	left *Node
	right *Node
	value Compareable
}

type BSTree struct {
	root *Node
}

func (bst *BSTree)Insert(node *Node)  {
	if node.value == nil {
		panic("insert null value node to tree")
	}
	if bst.root == nil {
		bst.root = node
		return
	}
	bst.insert(bst.root,node)
}

func (bst *BSTree)insert(root *Node,node *Node) {
	if root.value.Bg(node.value) {
		if root.left == nil {
			root.left = node
		}else {
			bst.insert(root.left,node)
		}
	}else{
		if root.right == nil {
			root.right = node
		}else {
			bst.insert(root.right,node)
		}
	}
}

func (bst *BSTree)InOrderTraversal()(nodes []*Node){
	if bst.root != nil {
		return  bst.inOrderTraversal(bst.root)
	}
	nodes = make([]*Node,0,0)
	return nodes
}

func (bst *BSTree)inOrderTraversal(root *Node)(nodes []*Node){
	nodes = make([]*Node,0,0)
	if root.left != nil {
		nodes = append(nodes,bst.inOrderTraversal(root.left)...)
	}else {
		nodes = append(nodes,root)
	}
	if root.right != nil {
		nodes = append(nodes,bst.inOrderTraversal(root.right)...)
	}
	return nodes
}

func (bst *BSTree)Find(targt Compareable)(node *Node){
	return bst.find(bst.root,targt)
}

func (bst *BSTree)find(root *Node,targt Compareable)(node *Node){
	if root == nil{
		return nil
	}
	if root.value.Eq(targt) {
		return root
	}else if root.value.Lt(targt){
		return bst.find(root.right,targt)
	}else {
		return bst.find(root.left,targt)
	}
}
//
//func (bst *BSTree)Delete(targt Compareable)bool{
//
//}
