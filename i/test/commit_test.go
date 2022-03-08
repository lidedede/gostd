package test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCheckCommit(t *testing.T) {
	type args struct {
		ts        []time.Time
		checkTime time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case001",
			args: args{
				ts: []time.Time{
					time.Date(2020, 11, 25, 0, 0, 0, 0, time.Local),
					time.Date(2020, 11, 26, 0, 0, 0, 0, time.Local),
					time.Date(2020, 11, 27, 0, 0, 0, 0, time.Local),
					time.Date(2020, 11, 28, 0, 0, 0, 0, time.Local),
				},
				checkTime: time.Date(2020, 11, 30, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCommit(tt.args.ts, tt.args.checkTime); got != tt.want {
				t.Errorf("CheckCommit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastTwoDayCommit(t *testing.T) {
	type args struct {
		ts        []time.Time
		checkTime time.Time
	}
	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastTwoDayCommit(tt.args.ts, tt.args.checkTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLastTwoDayCommit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastTwoDayStartTime(t *testing.T) {
	type args struct {
		cur time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
		{
			name: "case001",
			args: args{
				time.Date(2021, 11, 28, 10, 20, 20, 0, time.Local),
			},
			want: time.Date(2021, 11, 25, 00, 00, 00, 0, time.Local),
		},
		{
			name: "case002",
			args: args{
				time.Date(2021, 11, 27, 10, 20, 20, 0, time.Local),
			},
			want: time.Date(2021, 11, 25, 00, 00, 00, 0, time.Local),
		},
		{
			name: "case003",
			args: args{
				time.Date(2021, 11, 29, 10, 20, 20, 0, time.Local),
			},
			want: time.Date(2021, 11, 25, 00, 00, 00, 0, time.Local),
		},
		{
			name: "case004",
			args: args{
				time.Date(2021, 11, 30, 10, 20, 20, 0, time.Local),
			},
			want: time.Date(2021, 11, 26, 00, 00, 00, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotbef := getLastTwoDayStartTime(tt.args.cur)
			fmt.Println("args time ===", tt.args.cur)
			fmt.Println("want time ===", tt.want)
			fmt.Println("gots time ===", gotbef)

			if got := getLastTwoDayStartTime(tt.args.cur); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLastTwoDayStartTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isWorkDay(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "isWorkDay",
			args: args{
				t: time.Date(2021, 11, 25, 0, 0, 0, 0, time.Local),
			},
			want: true,
		},
		{
			name: "notWorkDay",
			args: args{
				t: time.Date(2021, 11, 27, 0, 0, 0, 0, time.Local),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWorkDay(tt.args.t); got != tt.want {
				t.Errorf("isWorkDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
