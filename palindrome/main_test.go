package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "correct hiragana palindrome verification",
			str:  "たけやぶやけた",
			want: true,
		}, {
			name: "correct katakana palindrome verification",
			str:  "マタマ",
			want: true,
		}, {
			name: "correct kanji palindrome verification",
			str:  "研学心肝心学研",
			want: true,
		}, {
			name: "correct alphabet palindrome verification",
			str:  "level",
			want: true,
		}, {
			name: "correct compound characters palindrome verification",
			str:  "東京バナナナナバ京東",
			want: true,
		}, {
			name: "Incorrect hiragana palindrome verification",
			str:  "おはよう",
			want: false,
		}, {
			name: "correct numbers palindrome verification",
			str:  "0220",
			want: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.str)
			if got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
