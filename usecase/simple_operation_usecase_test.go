package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRandomGenerator struct {
	a, b, c int
	counter int
}

func newMockRand(a, b, c int) *MockRandomGenerator {
	return &MockRandomGenerator{
		a: a,
		b: b,
		c: c,
	}
}

func (m *MockRandomGenerator) Intn(n int) int {
	m.counter++
	if m.counter == 1 {
		return m.a
	} else if m.counter == 2 {
		return m.b
	}
	return m.c
}

func TestSimpleOperationUsecaseSimpleOperation(t *testing.T) {
	tests := []struct {
		name   string
		uc     SimpleOperationUsecase
		want   int
		wantA  int
		wantB  int
		wantOp string
	}{
		{
			name:   "正常系: 1 + 3 = 4",
			uc:     NewSimpleOperationUsecase(newMockRand(0, 1, 3)),
			want:   4,
			wantA:  1,
			wantB:  3,
			wantOp: "+",
		},
		{
			name:   "正常系: 3 - 5 = -2",
			uc:     NewSimpleOperationUsecase(newMockRand(1, 3, 5)),
			want:   -2,
			wantA:  3,
			wantB:  5,
			wantOp: "-",
		},
		{
			name:   "正常系: 1 * 5 = 5",
			uc:     NewSimpleOperationUsecase(newMockRand(2, 1, 5)),
			want:   5,
			wantA:  1,
			wantB:  5,
			wantOp: "*",
		},
		{
			name:   "正常系: 乱数は最低でも1になる",
			uc:     NewSimpleOperationUsecase(newMockRand(0, 0, 0)),
			want:   2,
			wantA:  1,
			wantB:  1,
			wantOp: "+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)

			got, gotA, gotB, gotOp := tt.uc.Execute()
			a.Equal(tt.want, got)
			a.Equal(tt.wantA, gotA)
			a.Equal(tt.wantB, gotB)
			a.Equal(tt.wantOp, gotOp)
		})
	}
}
