package algo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCache(t *testing.T) {

	println("xx")
}

type Stu struct {
}

func TestCache1(t *testing.T) {
	q := make(chan int, 2)
	q <- 1
	q <- 2
	select {
	case q <- 3:
		fmt.Println("ok")
	default:
		for i := 0; i < 3; i++ {
			c := <-q
			fmt.Println(c)
		}

		fmt.Println("wrong")
	}
}

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	ins := []int{1, 2, 3, 3, 24, 2, 7, 8, 4, 3, 2, 5455, 343, 65464, 54}
	mate := [][]int{ins, ins, ins}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "xxx", args: args{matrix: mate}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
			t.Log(tt.args.matrix)
		})
	}
}

func Test_quickMul(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{args: args{x: 2, n: 3}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickMul(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("quickMul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func cew(x float64, n int) float64 {
	ans := 1.0
	arr_ans := x
	for n > 0 {
		if n%2 == 1 {
			ans *= x * arr_ans
		}
		arr_ans *= arr_ans
		n /= 2
	}
	return ans
}

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	nums := []int{-2, 1, -3, 4}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{want: 6, args: args{nums: nums}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args: args{[]int{1, 2, 2, 0, 0, 0, 1, 2, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			t.Log("nums:", tt.args.nums)
		})
	}
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	node1 := ListNode{Val: 4}
	node2 := ListNode{Val: 6}
	node3 := ListNode{Val: 7}
	tests := []struct {
		name     string
		args     args
		wantHead *ListNode
	}{
		// TODO: Add test cases.
		{
			args: args{l1: &ListNode{Val: 3}, l2: &ListNode{Val: 8}},
		},
	}
	node1.Next = &node3
	node2.Next = &node3
	tests[0].args.l1.Next = &node1
	tests[0].args.l2.Next = &node2
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHead := addTwoNumbers(tt.args.l1, tt.args.l2)
			prinListNode(gotHead)

		})
	}
}

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{want: []int{1, 2}, args: args{
			nums1: []int{1, 2, 3},
			nums2: []int{1, 2, 2},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 测试双链表队列
func Test_CQueue(t *testing.T) {
	cq := Constructor()
	for i := 0; i < 10; i++ {
		cq.AppendTail(i)
	}
	cq.DeleteHead()
	fmt.Println(cq)
}

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		// TODO: Add test cases.
		// 0 - 0 1 -1 2 -1 , 3 -2 4 - 3 5-5 6-8
		{args: args{n: 6}, wantSum: 8},
		{args: args{n: 600}, wantSum: 1317549726222168096},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := fib(tt.args.n); gotSum != tt.wantSum {
				t.Errorf("fib() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func Test_getArr(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{args: args{arr: []int{1, 2, 3, 4, 5}}, want: []int{1, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getArr(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{args: args{root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		}},
			want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	treeNode = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
)

func Test_Traversal(t *testing.T) {
	fmt.Println("前序遍历", preorderTraveral(treeNode))
	fmt.Println("后序遍历", postorderTraversal(treeNode))
	fmt.Println("中序遍历", inorderTraversal(treeNode))

}
func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		// TODO: Add test cases.
		{args: args{treeNode},
			wantRes: []int{2, 1, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := inorderTraversal(tt.args.root); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("inorderTraversal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
