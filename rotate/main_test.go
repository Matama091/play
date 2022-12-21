package main

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		name    string
		array   [][]int
		want    [][]int
		wantErr bool
	}{
		{
			name:    "rotate 1x1 2D array",
			array:   [][]int{{1}},
			want:    [][]int{{1}},
			wantErr: false,
		}, {
			name:    "rotate 2x2 2D array",
			array:   [][]int{{1, 2}, {3, 4}},
			want:    [][]int{{2, 4}, {1, 3}},
			wantErr: false,
		}, {
			name:    "rotate 3x3 2D array",
			array:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:    [][]int{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}},
			wantErr: false,
		}, {
			name:    "rotate 4x4 2D array",
			array:   [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
			want:    [][]int{{4, 8, 12, 16}, {3, 7, 11, 15}, {2, 6, 10, 14}, {1, 5, 9, 13}},
			wantErr: false,
		}, {
			name:    "rotate 3x2 2D array",
			array:   [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:    nil,
			wantErr: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Rotate(tt.array)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fibonacci() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for row := 0; row < len(got); row++ {
				for col := 0; col < len(got); col++ {
					if got[row][col] != tt.want[row][col] {
						t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
						return
					}
				}
			}
		})
	}
}
