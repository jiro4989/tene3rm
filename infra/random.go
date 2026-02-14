package infra

type RandomGenerator interface {
	Intn(n int) int
}

type MockRandomGenerator struct {
	num int
}

func NewMockRandomGenerator(n int) MockRandomGenerator {
	return MockRandomGenerator{
		num: n,
	}
}

func (m *MockRandomGenerator) Intn(n int) int {
	return m.num
}
