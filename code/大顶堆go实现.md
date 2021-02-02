#### 建堆

堆是用数组存储的，而且 0 下标不存，从 1 开始存储，建堆就是在原地通过交换位置，达到建堆的目的。完全二叉树我们知道，如果最后一个元素的下标为 n, 则 1 到 n/2 是非叶子节点，需要自上而下的堆化（和子节点比较），n/2 +1 到 n 是叶子节点，不需要堆化。

#### 插入一个元素

先插入的元素放到堆最后，然后和父节点比较，如果大于父节点，就交换位置，然后再和父节点比较，直到把这个元素放到正确的层。这种也叫自下而上的堆化。

#### 删除一个元素

假如删除大顶堆的，删除最大的元素，然后再它的子节点找到第二大元素，放到堆顶。然后再第二大元素下一层寻找

#### 堆排序

比如有 n 个数据，我们先把数据建堆，生成一个大顶堆，元素个数为 n

获取堆顶数据（也就是最大元素），删除堆顶，并且把最后一个元素放到堆顶，然后堆化成（n-1) 大顶堆，堆化的时间复杂度为 LogN, 底数是 2

重复获取堆顶，堆化成（n-2）大顶堆。我们获取的数据就是从大到小的顺序。



```go
import "fmt"

type heap struct {
	m   []int
	n int
}

/*
建堆，就是在原切片上操作，形成堆结构
只要按照顺序，把切片下标为n/2到1的节点依次堆化，最后就会把整个切片堆化
*/
// 传入的数组第一个元素必须是零值
func buildHeap(data []int) *heap {
	h := &heap{data, len(data) - 1}
	for i := h.n / 2; i > 0; i-- {
		h.heapify(i)
	}
	return h
}

func (h *heap) Push(x int) {
	h.n++
	h.m = append(h.m, x)		//向切片尾部插入数据（推断出父节点下标为i/2）
	// 从新的元素开始，把大的值往上提
	for i := h.n; i/2 > 0 && h.m[i/2] < h.m[i]; i /= 2 {
		h.swap(i, i/2)
	}
}

func (h *heap) Pop() (ret int) {
	ret = h.m[1]
	h.m[1] = h.m[h.n]	// 将最后一个元素放到堆顶
	h.m = h.m[:h.n]		// 修改堆的大小
	h.n--
	h.heapify(1)		// 从下标为1的节点向下堆化
	return
}

func (h *heap) swap(i, j int) {
	h.m[i], h.m[j] = h.m[j], h.m[i]
}

// 对下标为i的节点进行堆化，堆化的过程其实就是把i结点大的左右子结点往上提
func (h *heap) heapify(i int) {
	for {
		maxPos := i
		if 2*i <= h.n && h.m[2*i] > h.m[i] {
			maxPos = 2 * i
		}
		if 2*i+1 <= h.n && h.m[2*i+1] > h.m[maxPos] {
			maxPos = 2*i + 1
		}
		if maxPos == i {
			break
		}
		h.swap(i, maxPos)
		//h.m[i], h.m[maxPos] = h.m[maxPos], h.m[i]
		i = maxPos
	}
}

func main() {
	data := []int{0, 10, 28, 1, 2, 23, 34, 43}
	h := buildHeap(data)
	fmt.Println(h.m)
	fmt.Println(h.Pop())
	fmt.Println(h.Pop())
	h.Push(90)
	fmt.Println(h.m)
}
```

