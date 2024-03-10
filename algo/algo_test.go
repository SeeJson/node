package algo

import (
	"testing"
)

//func Test_rotate(t *testing.T) {
//	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
//	Rotate(matrix)
//	fmt.Println(matrix)
//}

func TestRotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"xx", args{matrix: matrix1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
