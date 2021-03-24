package timeutil

import (
	"testing"
	"time"
)

func BenchmarkParse(b *testing.B) {
	type args struct {
		t string
	}
	type test struct {
		name string
		args args
	}

	tests := []test{
		{
			name: "when t is 10ms",
			args: args{
				t: "10ms",
			},
		},
		{
			name: "when t is 100ms",
			args: args{
				t: "100ms",
			},
		},
		{
			name: "when t is 1s",
			args: args{
				t: "1s",
			},
		},
		{
			name: "when t is 10s",
			args: args{
				t: "10s",
			},
		},
		{
			name: "when t is 100s",
			args: args{
				t: "100s",
			},
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					got, err := Parse(test.args.t)
					if err != nil {
						b.Error(err)
						return
					}
					if got == 0 {
						b.Error("got is 0")
					}
				}
			})
		})
	}
}

func BenchmarkParseWithDefault(b *testing.B) {
	type args struct {
		t string
		d time.Duration
	}
	type test struct {
		name string
		args args
	}

	tests := []test{
		{
			name: "when t is 10second",
			args: args{
				t: "10second",
				d: 50 * time.Millisecond,
			},
		},
		{
			name: "when t is 100second",
			args: args{
				t: "100second",
				d: 50 * time.Millisecond,
			},
		},
		{
			name: "when t is 1000second",
			args: args{
				t: "1000second",
				d: 50 * time.Millisecond,
			},
		},
		{
			name: "when t is 10000second",
			args: args{
				t: "1000second",
				d: 50 * time.Millisecond,
			},
		},
		{
			name: "when t is 100000second",
			args: args{
				t: "10000second",
				d: 50 * time.Millisecond,
			},
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					got := ParseWithDefault(test.args.t, test.args.d)
					if got == 0 {
						b.Error("got is 0")
					}
				}
			})
		})
	}
}
