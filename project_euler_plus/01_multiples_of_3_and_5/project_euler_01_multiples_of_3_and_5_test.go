package main

import "testing"

func TestMultiplesOf3Or5(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test01",
			args: args{n: 1},
			want: 0,
		},
		{
			name: "Test02",
			args: args{n: 2},
			want: 0,
		},
		{
			name: "Test03",
			args: args{n: 3},
			want: 0,
		},
		{
			name: "Test04",
			args: args{n: 4},
			want: 3,
		},
		{
			name: "Test05",
			args: args{n: 5},
			want: 3,
		},
		{
			name: "Test06",
			args: args{n: 6},
			want: 8,
		},
		{
			name: "Test07",
			args: args{n: 7},
			want: 14,
		},
		{
			name: "Test08",
			args: args{n: 8},
			want: 14,
		},
		{
			name: "Test09",
			args: args{n: 9},
			want: 14,
		},
		{
			name: "Test10",
			args: args{n: 10},
			want: 23,
		},
		{
			name: "Test11",
			args: args{n: 11},
			want: 33,
		},
		{
			name: "Test12",
			args: args{n: 12},
			want: 33,
		},
		{
			name: "Test13",
			args: args{n: 13},
			want: 45,
		},
		{
			name: "Test14",
			args: args{n: 14},
			want: 45,
		},
		{
			name: "Test15",
			args: args{n: 15},
			want: 45,
		},
		{
			name: "Test16",
			args: args{n: 16},
			want: 60,
		},
		{
			name: "Test100",
			args: args{n: 100},
			want: 2318,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplesOf3Or5(tt.args.n); got != tt.want {
				t.Errorf("MultiplesOf3Or5() = %v, want %v", got, tt.want)
			}
		})
	}
}
