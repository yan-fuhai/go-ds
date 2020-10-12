package cache

type doubleListNode struct {
	key   interface{}
	val   interface{}
	left  *doubleListNode
	right *doubleListNode
}

type LRUCache struct {
	capacity int
	head     *doubleListNode
	tail     *doubleListNode
	keyMap   map[interface{}]*doubleListNode
}

func NewLRUCache(capacity int) *LRUCache {
	head, tail := &doubleListNode{}, &doubleListNode{}
	head.right, tail.left = tail, head
	return &LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		keyMap:   make(map[interface{}]*doubleListNode),
	}
}

func (c *LRUCache) Put(k interface{}, v interface{}) {
	if nPtr, has := c.keyMap[k]; has {
		c.moveToHead(nPtr)
		if v != nPtr.val {
			nPtr.val = v
		}
	} else {
		newNode := &doubleListNode{
			key:   k,
			val:   v,
			left:  nil,
			right: nil,
		}
		c.addToHead(newNode)
		c.keyMap[k] = newNode
		if len(c.keyMap) > c.capacity {
			c.removeTail()
		}
	}
}

func (c *LRUCache) Get(k interface{}) interface{} {
	if nPtr, has := c.keyMap[k]; has {
		c.moveToHead(nPtr)
		return nPtr.val
	}
	return nil
}

func (c *LRUCache) Clear() {
	c.head.right = c.tail
	c.tail.left = c.head
	c.keyMap = make(map[interface{}]*doubleListNode)
}

func (c *LRUCache) moveToHead(node *doubleListNode) {
	if node != nil {
		c.removeNode(node)
		c.addToHead(node)
	}
}

func (c *LRUCache) removeNode(node *doubleListNode) {
	node.right.left, node.left.right = node.left, node.right
}

func (c *LRUCache) addToHead(node *doubleListNode) {
	node.left, node.right = c.head, c.head.right
	c.head.right.left, c.head.right = node, node
}

func (c *LRUCache) removeTail() {
	if c.head.right != c.tail {
		c.removeNode(c.tail.left)
	}
}
