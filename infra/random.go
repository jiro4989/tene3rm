package infra

type RandomGenerator interface {
	Intn(n int) int
}

type MockRandom struct {
	num int
}

func NewMockRandom(n int) MockRandom {
	return MockRandom{
		num: n,
	}
}

func (m *MockRandom) Intn(n int) int {
	return m.num
}
