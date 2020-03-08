package jumble

import (
	"reflect"
	"testing"
)

func TestIndices(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{"Test11", args{"Jumble.  changes äöü !  a"},
			[][]int{{0, 5}, {9, 15}, {17, 19}, {24, 24}}, false},
		{"Test12", args{"  é@%zz** ab-c     u ! e m12m n   "},
			[][]int{{2, 2}, {5, 6}, {10, 11}, {13, 13}, {19, 19}, {23, 23}, {25, 25}, {28, 28}, {30, 30}}, false},
		{"Test13", args{" %&- 84. 846 -@$`! "}, nil, true},
		{"Test14", args{""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Indices(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Indices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Indices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWord(t *testing.T) {
	type args struct {
		str  []rune
		seed int64
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"Test21", args{[]rune("characters"), 1000000000}, []rune("cathrcreas")},
		{"Test22", args{[]rune(""), 1000000000}, nil},
		{"Test23", args{[]rune("a"), 1000000000}, []rune("a")},
		{"Test24", args{[]rune("ab"), 1000000000}, []rune("ab")},
		{"Test25", args{nil, 1000000000}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Word(tt.args.str, tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Word() = %v, want %v", got, tt.want)
			}
		})
	}
}
