package goroutine

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 检测数 t1 和 t2 是否含有相同的值
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		// 没关闭，有空改一下
		fmt.Printf("tree 1 sorted value %d", i)
		j, ok := <-ch2
		fmt.Printf("tree 2 sorted value %d", j)
		if i != j || !ok {
			return false
		}
	}
	_, empty := <-ch2
	if !empty {
		return false
	}
	return true
}

func main() {
	res1 := Same(tree.New(1), tree.New(1))
	res2 := Same(tree.New(1), tree.New(2))
	fmt.Printf("result one for %d & result two for %d", res1, res2)
}
