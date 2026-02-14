package usecase

import (
	"testing"

	"github.com/jiro4989/tene3rm/infra"
	"github.com/stretchr/testify/assert"
)

func TestMathServiceSimpleOperation(t *testing.T) {
	oAdd := infra.NewMockRandom(0)
	oMinus := infra.NewMockRandom(1)
	oMulti := infra.NewMockRandom(2)
	n0 := infra.NewMockRandom(0)
	n1 := infra.NewMockRandom(1)
	n3 := infra.NewMockRandom(3)
	n5 := infra.NewMockRandom(5)

	tests := []struct {
		name    string
		svc     MathService
		o, a, b infra.RandomGenerator
		want    int
		wantA   int
		wantB   int
		wantOp  string
	}{
		{
			name:   "正常系: 1 + 3 = 4",
			svc:    NewMathService(),
			o:      &oAdd,
			a:      &n1,
			b:      &n3,
			want:   4,
			wantA:  1,
			wantB:  3,
			wantOp: "+",
		},
		{
			name:   "正常系: 3 - 5 = -2",
			svc:    NewMathService(),
			o:      &oMinus,
			a:      &n3,
			b:      &n5,
			want:   -2,
			wantA:  3,
			wantB:  5,
			wantOp: "-",
		},
		{
			name:   "正常系: 1 * 5 = 5",
			svc:    NewMathService(),
			o:      &oMulti,
			a:      &n1,
			b:      &n5,
			want:   5,
			wantA:  1,
			wantB:  5,
			wantOp: "*",
		},
		{
			name:   "正常系: 乱数は最低でも1になる",
			svc:    NewMathService(),
			o:      &oAdd,
			a:      &n0,
			b:      &n0,
			want:   2,
			wantA:  1,
			wantB:  1,
			wantOp: "+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, gotA, gotB, gotOp := tt.svc.SimpleOperations(tt.o, tt.a, tt.b)
			a.Equal(tt.want, got)
			a.Equal(tt.wantA, gotA)
			a.Equal(tt.wantB, gotB)
			a.Equal(tt.wantOp, gotOp)
		})
	}
}
