package packa

import (
	"os"
	"testing"
	"time"
)

func TestSleep1SecAndEcho(t *testing.T) {
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
				input: "hello",
			},
			want: "hello",
		},
		{
			name: "ok",
			args: args{
				input: "hello",
			},
			want: "hello",
		},
		{
			name: "ok",
			args: args{
				input: "hello",
			},
			want: "hello",
		},
		{
			name: "ok",
			args: args{
				input: "hello",
			},
			want: "hello",
		},
		{
			name: "ok",
			args: args{
				input: "hello",
			},
			want: "hello",
		},
		{
			name: "ok",
			args: args{
				input: "hello",
			},
			want: "hello",
		},
		{
			name: "ok",
			args: args{
				input: "hello",
			},
			want: "hello",
		},
		{
			name: "ok",
			args: args{
				input: "hello",
			},
			want: "hello",
		},
		{
			name: "ok",
			args: args{
				input: "hello",
			},
			want: "hello",
		},
		{
			name: "ok",
			args: args{
				input: "hello",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if os.Getenv("CASE_PARALLEL") == "on" {
				t.Parallel()
			}
			if got := Sleep1SecAndEcho(tt.args.input); got != tt.want {
				t.Errorf("Sleep1SecAndEcho() = %v, want %v", got, tt.want)
			}
		})
	}
}
