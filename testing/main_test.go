package main

import (
	"testing"
	"fmt"
)

func testPrint(t *testing.T) {
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

func testPrint2(t *testing.T) {
	res := Print1to20()
	res++
	if res != 211 {
		t.Errorf("Test Print2 failed")
	}
}

func TestAll(t *testing.T)  {
	// t.Run()，第一个参数为显示的名字，可自行命名，第二个参数为函数名
	t.Run("TestPrint*", testPrint)
	t.Run("TestPrint2*", testPrint2)
}

func Testmain(m *testing.M)  {
	fmt.Println("TestMain begin...")
	m.Run()
}

func BenchmarkAll(b *testing.B)  {
	for n := 0; n < b.N; n++ {
		Print1to20()
	}
}