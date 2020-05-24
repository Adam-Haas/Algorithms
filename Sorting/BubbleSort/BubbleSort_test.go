package BubbleSort

import (
	"reflect"
	"testing"
)

func TestPerformBubbleSortForInt(t *testing.T) {
	type args struct {
		inputArr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test1",
			args{[]int{45, 23, 421, 2, 420}},
			[]int{2, 23, 45, 420, 421},
		},
		{
			"Test2",
			args{[]int{4, 12, 8, 9}},
			[]int{4, 8, 9, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PerformBubbleSortForInt(tt.args.inputArr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PerformBubbleSortForInt() = %v, want %v", got, tt.want)
			}
		})
	}
}