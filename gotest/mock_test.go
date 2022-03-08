package main

import (
	"fmt"
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(3, 2)
	assert.Equal(t, 5, sum)
}

func TestNeedMock(t *testing.T) {
	patches := gomonkey.ApplyFunc(Add, func(a, b int) int {
		fmt.Println("mocked")
		return 18
	})
	defer patches.Reset()
	result := NeedMock(3)
	assert.Equal(t, 21, result)
}
