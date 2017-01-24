package implementation

import "testing"

func Test_mini_max_sum(t *testing.T) {
	type args struct {
		nums []uint64
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 uint64
	}{
		// TODO: Add test cases.
		{
			name:  "mini_max_sum_01",
			args:  args{nums: []uint64{1, 2, 3, 4, 5}},
			want:  10,
			want1: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := mini_max_sum(tt.args.nums)
			if got != tt.want {
				t.Errorf("mini_max_sum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("mini_max_sum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
