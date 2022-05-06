package main

import "testing"

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "test-1",
			arg:  "",
			want: 0,
		},
		{
			name: "test-2",
			arg:  "a",
			want: 1,
		},
		{
			name: "test-3",
			arg:  "ab",
			want: 2,
		},
		{
			name: "test-4",
			arg:  "aa",
			want: 2,
		},
		{
			name: "test-5",
			arg:  "aaba",
			want: 4,
		},
		{
			name: "test-6",
			arg:  "abcdef",
			want: 2,
		},
		{
			name: "test-7",
			arg:  "aaaaa",
			want: 5,
		},
		{
			name: "test-8",
			arg:  "abaccab",
			want: 4,
		},
		{
			name: "test-9",
			arg:  "vabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaabvabaccabbaaab",
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.arg); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
