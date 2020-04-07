package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestGetRandomInt(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomInt(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("GetRandomInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownLoadimg(t *testing.T) {
	type args struct {
		filename string
		url      string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DownLoadimg(tt.args.filename, tt.args.url)
		})
	}
}

func TestDownLoadimgAsync(t *testing.T) {
	type args struct {
		filename string
		url      string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DownLoadimgAsync(tt.args.filename, tt.args.url)
		})
	}
}

func Test_spiderPhone(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spiderPhone()
		})
	}
}

func Test_spiderLink(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spiderLink()
		})
	}
}

func Test_getHtml(t *testing.T) {
	type args struct {
		tarUrl string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHtml(tt.args.tarUrl); got != tt.want {
				t.Errorf("getHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}
