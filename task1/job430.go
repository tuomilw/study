package task1

/*
*
430. 扁平化多级双向链表：多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。可以定义一个结构体来表示链表节点，包含 val、prev、next 和 child 指针，然后使用递归的方法来扁平化链表，先处理当前节点的子链表，再将子链表插入到当前节点和下一个节点之间。
*/
func GetFlatLinkList(node *Node) *Node {
	if node == nil {
		return nil
	}
	return flattenDFS(node)
}

/*
*

	深度优先遍历DFS：优先处理子链表，再处理后续节点。
	递归处理：递归展平子链表并返回尾节点，将其插入到当前节点与后续节点之间。
*/
func flattenDFS(node *Node) *Node {
	curr := node   //当前处理的节点
	var last *Node //记录当前链表的最后一个尾节点
	for curr != nil {
		next := curr.Next      //保存原next节点
		if curr.Child != nil { //如果当前节点的子链表不为空
			//递归展平子链表，获取子链表的尾节点
			childTail := flattenDFS(curr.Child)
			//将子链表插入current和next之间
			curr.Next = curr.Child
			curr.Child.Prev = curr
			childTail.Next = next
			if next != nil {
				next.Prev = childTail
			}
			//清空child指针，避免循环引用
			curr.Child = nil
			//更新last为子链表的尾节点
			last = childTail
		} else {
			//没有子链表，当前节点就是最后一个尾节点
			last = curr
		}
		curr = next
	}
	return last
}

type Node struct {
	Num   int
	Prev  *Node
	Next  *Node
	Child *Node //子链表指针
}
