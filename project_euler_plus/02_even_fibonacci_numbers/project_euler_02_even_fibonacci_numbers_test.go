package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		n uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{
			name: "search00",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "search01",
			args: args{
				n: 1,
			},
			want: 0,
		},
		{
			name: "search02",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "search03",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "search04",
			args: args{
				n: 4,
			},
			want: 2,
		},
		{
			name: "search05",
			args: args{
				n: 5,
			},
			want: 2,
		},
		{
			name: "search06",
			args: args{
				n: 6,
			},
			want: 2,
		},
		{
			name: "search07",
			args: args{
				n: 7,
			},
			want: 2,
		},
		{
			name: "search08",
			args: args{
				n: 8,
			},
			want: 10,
		},
		{
			name: "search09",
			args: args{
				n: 9,
			},
			want: 10,
		},
		{
			name: "search10",
			args: args{
				n: 10,
			},
			want: 10,
		},
		{
			name: "search11",
			args: args{
				n: 11,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.n); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
