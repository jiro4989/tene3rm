package column

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRow_ToInt(t *testing.T) {
	tests := []struct {
		name    string
		r       Row
		index   int
		want    int
		wantErr bool
	}{
		{
			name: "正常系: 筆算の入力として正しい数値順",
			r: Row{
				value: []Cell{
					NewCell(" "),
					NewCell(" "),
					NewCell("1"),
					NewCell("2"),
				},
			},
			index:   0,
			want:    12,
			wantErr: false,
		},
		{
			name: "異常系: 間に空白があるのは異常",
			r: Row{
				value: []Cell{
					NewCell(" "),
					NewCell("1"),
					NewCell(" "),
					NewCell("2"),
				},
			},
			index:   0,
			want:    0,
			wantErr: true,
		},
		{
			name: "異常系: index 0 では末尾空白2つは異常",
			r: Row{
				value: []Cell{
					NewCell(" "),
					NewCell("1"),
					NewCell("2"),
					NewCell(" "),
				},
			},
			index:   0,
			want:    0,
			wantErr: true,
		},
		{
			name: "異常系: index 0 では末尾空白2つは異常",
			r: Row{
				value: []Cell{
					NewCell("1"),
					NewCell("2"),
					NewCell(" "),
					NewCell(" "),
				},
			},
			index:   0,
			want:    0,
			wantErr: true,
		},
		{
			name: "異常系: index 0 では先頭に数字が入ってるのは異常",
			r: Row{
				value: []Cell{
					NewCell("1"),
					NewCell("2"),
					NewCell("3"),
					NewCell("4"),
				},
			},
			index:   0,
			want:    0,
			wantErr: true,
		},
		{
			name: "異常系: index 1 では中間空白2つは異常",
			r: Row{
				value: []Cell{
					NewCell(" "),
					NewCell(" "),
					NewCell("2"),
					NewCell(" "),
				},
			},
			index:   1,
			want:    0,
			wantErr: true,
		},
		{
			name: "異常系: index 1 では中間空白2つは異常",
			r: Row{
				value: []Cell{
					NewCell(" "),
					NewCell("2"),
					NewCell(" "),
					NewCell(" "),
				},
			},
			index:   1,
			want:    0,
			wantErr: true,
		},
		{
			name: "異常系: index 1 では末尾に空白以外があると異常",
			r: Row{
				value: []Cell{
					NewCell("1"),
					NewCell("2"),
					NewCell("3"),
					NewCell("4"),
				},
			},
			index:   1,
			want:    0,
			wantErr: true,
		},
		{
			name: "異常系: index 2 では末尾空白3つは異常",
			r: Row{
				value: []Cell{
					NewCell("2"),
					NewCell(" "),
					NewCell(" "),
					NewCell(" "),
				},
			},
			index:   2,
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, err := tt.r.ToInt(tt.index)
			if tt.wantErr {
				a.Error(err)
				return
			}

			a.Equal(tt.want, got)
		})
	}
}
