package usecase

import (
	"testing"

	"github.com/jiro4989/tene3rm/infra"
	"github.com/stretchr/testify/assert"
)

func TestSimpleOperationUsecaseSimpleOperation(t *testing.T) {
	oAdd := infra.NewMockRandomGenerator(0)
	oMinus := infra.NewMockRandomGenerator(1)
	oMulti := infra.NewMockRandomGenerator(2)
	n0 := infra.NewMockRandomGenerator(0)
	n1 := infra.NewMockRandomGenerator(1)
	n3 := infra.NewMockRandomGenerator(3)
	n5 := infra.NewMockRandomGenerator(5)

	tests := []struct {
		name    string
		svc     SimpleOperationUsecase
		o, a, b infra.RandomGenerator
		want    int
		wantA   int
		wantB   int
		wantOp  string
	}{
		{
			name:   "正常系: 1 + 3 = 4",
			svc:    NewSimpleOperationUsecase(),
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
			svc:    NewSimpleOperationUsecase(),
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
			svc:    NewSimpleOperationUsecase(),
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
			svc:    NewSimpleOperationUsecase(),
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
