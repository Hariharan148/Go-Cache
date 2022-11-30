package main

import "fmt"

const SIZE = 5


type Node struct{
	Left *Node
	Right *Node
	Val string
}

type Queue struct{
	Head *Node
	Tail *Node
	Length int
}

type Cache struct{
	Queue Queue
	Hash Hash
}

type Hash map[string]*Node


func NewQueue() Queue{
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func NewCache() Cache{
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func (c *Cache) Check(str string){
	node := &Node{}

	if val, ok := c.Hash[str]; ok{
		node = c.Remove(val)
	} else{
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache)Add(n *Node){
	fmt.Printf("\nAdd:%s\n", n.Val)

	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE{
		c.Remove(c.Queue.Tail.Left)
	}

}

func (c *Cache) Remove(n *Node) *Node{
	fmt.Printf("Remove:%s", n.Val)

	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	delete(c.Hash,n.Val )
	c.Queue.Length -= 1

	return n
}


func (c *Cache) Display(){
	c.Queue.Display()
}

func (c *Queue) Display(){
	node := c.Head.Right

	fmt.Printf("\n%d - [", c.Length)
	for i := 0; i < c.Length; i++{
		fmt.Printf("{%s}", node.Val)
		if i < c.Length - 1 {
			fmt.Printf(" < -- > ")
		}
		node = node.Right
	}
	fmt.Println("]")
}


func main(){
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{"ironman", "captain america", "spiderman", "zendaya", "ironman", "black widow", "thanos"}{
		cache.Check(word)
		cache.Display()
	}
}