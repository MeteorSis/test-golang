package packb

import (
	"os"
	"testing"
	"time"
)

func TestSleep1SecAndEchoTwice(t *testing.T) {
	if os.Getenv("FUNC_PARALLEL") == "on" {
		t.Parallel()
	}
	time.Sleep(10 * time.Second)
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "11",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "11",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "11",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "11",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "11",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "11",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "11",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "11",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "11",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "11",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if os.Getenv("CASE_PARALLEL") == "on" {
				t.Parallel()
			}
			if got := Sleep1SecAndEchoTwice(tt.args.input); got != tt.want {
				t.Errorf("Sleep1SecAndEchoTwice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSleep1SecAndEchoThreeTimes(t *testing.T) {
	if os.Getenv("FUNC_PARALLEL") == "on" {
		t.Parallel()
	}
	time.Sleep(10 * time.Second)
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "111",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "111",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "111",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "111",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "111",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "111",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "111",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "111",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "111",
		},
		{
			name: "ok",
			args: args{
				input: "1",
			},
			want: "111",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if os.Getenv("CASE_PARALLEL") == "on" {
				t.Parallel()
			}
			if got := Sleep1SecAndEchoThreeTimes(tt.args.input); got != tt.want {
				t.Errorf("Sleep1SecAndEchoThreeTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
