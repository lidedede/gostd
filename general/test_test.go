package main

import "testing"

func TestStudent_ShowMe(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test001",args{"test"},"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stu := &Student{}
			if got := stu.ShowMe(tt.args.s); got != tt.want {
				t.Errorf("ShowMe() = %v, want %v", got, tt.want)
			}
		})
	}
}
