package utils

type SortingHeapNode struct {
	Val    int
	Father *SortingHeapNode
	Left   *SortingHeapNode
	Right  *SortingHeapNode
}

func InitSortingHeap(rootVal int) *SortingHeapNode {
	return &SortingHeapNode{
		Val: rootVal,
	}
}

func (_this *SortingHeapNode) Insert(nodeVal int) {
	if nodeVal <= _this.Val {
		if _this.Left == nil {
			_this.Left = &SortingHeapNode{
				Val:    nodeVal,
				Father: _this,
			}
		} else {
			_this.Left.Insert(nodeVal)
		}
	} else {
		if _this.Right == nil {
			_this.Right = &SortingHeapNode{
				Val:    nodeVal,
				Father: _this,
			}
		} else {
			_this.Right.Insert(nodeVal)
		}
	}
}

func (_this *SortingHeapNode) SortedTraverse() []int {
	// 结果数组
	result := make([]int, 0)
	// 先序递归遍历左枝丫
	if _this.Left != nil {
		result = append(result, _this.Left.SortedTraverse()...)
	}
	// 中间遍历本枝
	result = append(result, _this.Val)
	// 后续递归遍历右枝丫
	if _this.Right != nil {
		result = append(result, _this.Right.SortedTraverse()...)
	}
	return result
}
