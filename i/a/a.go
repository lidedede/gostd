package a

type StructA struct {
	elem MockB
}

type MockB interface {
	SayNum() int
}

