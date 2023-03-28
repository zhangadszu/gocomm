package datatime

import "testing"

func TestIsLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				year: 2000,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				year: 1900,
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				year: 2100,
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				year: 2020,
			},
			want: true,
		},
		{
			name: "test5",
			args: args{
				year: 2021,
			},
			want: false,
		},
		{
			name: "test6",
			args: args{
				year: 2022,
			},
			want: false,
		},
		{
			name: "test7",
			args: args{
				year: 2023,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeapYear(tt.args.year); got != tt.want {
				t.Errorf("IsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
