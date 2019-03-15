package mymath

import (
	"Gogoyan/gotest/mymath"
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := mymath.Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_Division_2(t *testing.T) {
	if i, e := mymath.Division(6, 0); i != 0 || e == nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第二个测试通过了") //记录一些你期望记录的信息
	}
}

// this is an error Test_Xxx case
// func Test_Division_3(t *testing.T) {
// 	t.Error("就是不通过")
// }
