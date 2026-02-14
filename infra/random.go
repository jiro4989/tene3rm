package infra

type RandomGenerator interface {
	Intn(n int) int
}
