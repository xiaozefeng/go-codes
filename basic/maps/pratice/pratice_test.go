package pratice

import "testing"

func Test_lengthOfNonRepeatingSubStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{s: "aabbcc"}, 2},
		{"normal", args{s: "bbbbb"}, 1},
		{"test3", args{s: "pwwkew"}, 3},
		{"中文", args{s: "吃饭了吗"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfNonRepeatingSubStr(tt.args.s); got != tt.want {
				t.Errorf("lengthOfNonRepeatingSubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
