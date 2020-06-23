package test_exam

import "testing"

var util Triangle

func TestTriangle(t *testing.T) {
	TestTriangle_IsTriangle(t)
	TestGeneralBoundaryOfTriangle(t)
}

func TestTriangle_IsTriangle(t *testing.T) {
	util.IsTriangle("a", "1", "2")
	util.IsTriangle("", "50", "50")
	util.IsTriangle("101", "50", "50")
	util.IsTriangle("2.1", "3", "4")
	util.IsTriangle("100", "50", "50")
	util.IsTriangle("1", "50", "50")
	util.IsTriangle("50", "50", "50")
	util.IsTriangle("4", "5", "6")
	util.IsTriangle("3", "4", "5")
}

func TestGeneralBoundaryOfTriangle(t *testing.T){
	t.Log("这是三角形问题一般边界测试的结果")
	util.IsTriangle("60", "60", "1")
	util.IsTriangle("60", "60", "2")
	util.IsTriangle("60", "60", "2")
	util.IsTriangle("60", "60", "60")
	util.IsTriangle("50", "50", "99")
	util.IsTriangle("50", "50", "100")
	util.IsTriangle("60", "1", "60")
	util.IsTriangle("60", "2", "60")
	util.IsTriangle("50", "99", "50")
	util.IsTriangle("50", "100", "50")
	util.IsTriangle("1", "60", "60")
	util.IsTriangle("2", "60", "60")
	util.IsTriangle("99", "50", "50")
	util.IsTriangle("100", "50", "50")
}
