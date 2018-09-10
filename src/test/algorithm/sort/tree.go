package sort

const (
	nodeLeft  = 0
	nodeRight = 1
)

type Node struct {
	key      int
	val      int
	parent   *Node
	children [2]*Node
}

func (root *Node) Find(k int) (n **Node, pnode *Node, dep int) {
	n = &root
	for *n != nil {
		dep++
		pnode = *n
		if k < (*n).key {
			n = &(*n).children[0]
		} else {
			n = &(*n).children[1]
		}
	}
	return
}

//二叉树
type BinaryTree struct {
	length int
	head   *Node
}

func (bt *BinaryTree) Insert(k int) {
	//查找k应该插入的位置，使用循环因为可能 k 有重复值
	var root, p *Node
	n := &bt.head
	for *n != nil {
		root = *n
		n, p, _ = root.Find(k)
	}
	//添加新节点
	*n = &Node{
		key:      k,
		val:      k,
		parent:   p,
		children: [2]*Node{},
	}
	bt.length++
	return
}

func (bt *BinaryTree) PreOrder(root *Node) (sorted []int) {
	var n = root
	for n != nil {
		//遍历中间节点
		sorted = append(sorted, n.key)
		if n.children[0] != nil {
			//遍历左孩子
			n = n.children[0]
			continue
		} else if n.children[1] != nil {
			//无左孩子，遍历右孩子
			n = n.children[1]
			continue
		}
		//递归查找右兄弟节点
		for n.parent != nil {
			//已到叶子节点，若自身为左孩子，则将自己替换为右孩子
			if n.parent.children[0] == n && n.parent.children[1] != nil {
				n = n.parent.children[1]
				break
			}
			n = n.parent
		}
		//已回溯到根节点，且根节点左右孩子已遍历完成
		if n.parent == nil {
			break
		}
	}
	return sorted
}

func (bt *BinaryTree) InOrder(root *Node) (sorted []int) {
	var n = root
	var i = 0
	for n != nil && i < 100 {
		i++
		if n.children[0] != nil {
			//遍历左孩子
			n = n.children[0]
			continue
		} else if n.children[1] != nil {
			//无左孩子，遍历右孩子
			sorted = append(sorted, n.key)
			n = n.children[1]
			continue
		}
		//左右孩子都没有，且父节点有左孩子
		sorted = append(sorted, n.key)
		for n.parent != nil && i < 100 {
			i++
			if n.parent.children[0] == n {
				sorted = append(sorted, n.parent.key)
				if n.parent.children[1] != nil {
					n = n.parent.children[1]
					break
				}
			}
			n = n.parent
		}
		if n.parent == nil {
			break
		}
	}
	return sorted
}

func (bt *BinaryTree) PostOrder(root *Node) (sorted []int) {
	var n = root
	for n != nil {
		if n.children[0] != nil {
			n = n.children[0]
			continue
		} else if n.children[1] != nil {
			n = n.children[1]
			continue
		}
		sorted = append(sorted, n.key)
		for n.parent != nil {
			if n.parent.children[0] == n && n.parent.children[1] != nil {
				n = n.parent.children[1]
				break
			}
			sorted = append(sorted, n.parent.key)
			n = n.parent
		}
		if n.parent == nil {
			break
		}
	}
	return
}

//堆
type Heap struct {
	nl []*Node
}

func (h *Heap) addTail(k int) {
	var p *Node
	pIndex := h.pIndex(len(h.nl))
	if pIndex >= 0 {
		p = h.nl[pIndex]
	}
	cIndex := (len(h.nl) - 1) % 2
	p.children[cIndex] = &Node{
		key:      k,
		val:      k,
		parent:   p,
		children: [2]*Node{},
	}
	h.nl = append(h.nl, p.children[cIndex])
}

func (h *Heap) pIndex(i int) int {
	return i/2 + i%2 - 1
}

func (h *Heap) exchange(nl []*Node, i, j int) {
	//仅交换key和val，省略指针的变换
	nl[i].key, nl[i].val, nl[j].key, nl[j].val = nl[j].key, nl[j].val, nl[i].key, nl[i].val
}

func (h *Heap) Insert(k int) {
	//添加至末尾
	h.addTail(k)
	var i = len(h.nl) - 1
	var p int
	for i > 0 {
		p = h.pIndex(i)
		if k < h.nl[p].key {
			h.exchange(h.nl, i, p)
		}
		i = p
	}
}

func (h *Heap) Pop() (node *Node) {
	return
}

func (h *Heap) Queue() (nl []*Node) {
	for len(h.nl) > 0 {
		nl = append(nl, h.Pop())
	}
	return
}

func (h *Heap) leftTurn() {

}

func (h *Heap) rightTurn() {

}
