package a

import (
	"fmt"
	"github.com/lidedede/gostd/icode/b"
	"testing"
)

func TestStructA_SayNum1(t *testing.T) {
	aaa := &StructA{
		elem: nil,
	}
	fmt.Println(aaa)
	var mb MockB = new(b.StructB)
	a := StructA{
		elem: mb,
	}
	res := a.SayNum()
	if want := "小于100"; res != want {
		t.Errorf("case：小于10，失败")
	}
}

type MockStructB struct{}

func (a *MockStructB) SayNum() int {
	return 200
}

func TestStructA_SayNum2(t *testing.T) {
	var mb MockB = new(MockStructB)
	a := StructA{
		elem: mb,
	}
	res := a.SayNum()
	if want := "大于等于100"; res != want {
		t.Errorf("case：大于等于100，失败")
	}
}
