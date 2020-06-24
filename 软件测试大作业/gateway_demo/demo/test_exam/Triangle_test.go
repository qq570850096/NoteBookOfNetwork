package test_exam

import (
	"testing"
)

var util Triangle

func TestTriangle(t *testing.T) {
	TestTriangle_IsTriangle(t)
	TestGeneralBoundaryOfTriangle(t)
	TestConditionCover(t)
	TestDesicionCover(t)
	TestEClassOfTriangle(t)
	TestDesicionTableOfTriangle(t)
	TestRobustBoundaryOfTriangle(t)
	TestStmtCover(t)
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
	t.Log("一般边界测试成功！")
}

func TestRobustBoundaryOfTriangle(t *testing.T) {
	t.Log("这是三角形问题健壮性边界值测试的结果:")
	util.IsTriangle("0", "50", "50")
	util.IsTriangle("1", "50", "50")
	util.IsTriangle("2", "50", "50")
	util.IsTriangle("50", "50", "50")
	util.IsTriangle("99", "50", "50")
	util.IsTriangle("100", "50", "50")
	util.IsTriangle("101", "50", "50")
	util.IsTriangle("50", "0", "50")
	util.IsTriangle("50", "1", "50")
	util.IsTriangle("50", "2", "50")
	util.IsTriangle("50", "99", "50")
	util.IsTriangle("50", "100", "50")
	util.IsTriangle("50", "101", "50")
	util.IsTriangle("50", "50", "0")
	util.IsTriangle("50", "50", "1")
	util.IsTriangle("50", "50", "2")
	util.IsTriangle("50", "50", "99")
	util.IsTriangle("50", "50", "100")
	util.IsTriangle("50", "50", "101")
	t.Log("健壮性边界值测试成功")
}

func TestEClassOfTriangle(t *testing.T){
	t.Log("这是三角形问题等价类划分的结果:")
	util.IsTriangle("2", "3", "4")
	util.IsTriangle("2", "2", "3")
	util.IsTriangle("2", "3", "2")
	util.IsTriangle("3", "2", "2")
	util.IsTriangle("2", "2", "2")
	util.IsTriangle("1", "2", "3")
	util.IsTriangle("2.1", "3", "4")
	util.IsTriangle("2.1", "2.1", "3")
	util.IsTriangle("2.1", "2.1", "2.1")
	util.IsTriangle("0", "2", "3")
	util.IsTriangle("0", "0", "2")
	util.IsTriangle("0", "0", "0")
	util.IsTriangle("-1", "2", "2")
	util.IsTriangle("-1", "-2", "2")
	util.IsTriangle("-2", "-2", "-2")
	util.IsTriangle("111", "99", "99")
	util.IsTriangle("111", "111", "99")
	util.IsTriangle("111", "111", "111")
	util.IsTriangle("5", "3", "2")
	util.IsTriangle("5", "2", "3")
	util.IsTriangle("2", "5", "3")
	util.IsTriangle("", "", "")
	util.IsTriangle("2", "3", "")
	util.IsTriangle("3", "4", "5")
	t.Log("等价类划分测试成功！")
}

func TestDesicionTableOfTriangle(t *testing.T){
	t.Log("这是三角形问题决策表测试的结果:")
	util.IsTriangle("4", "1", "2")
	util.IsTriangle("1", "4", "2")
	util.IsTriangle("1", "2", "4")
	util.IsTriangle("5", "5", "5")
	util.IsTriangle("?", "?", "?")
	util.IsTriangle("2", "2", "3")
	util.IsTriangle("2", "3", "2")
	util.IsTriangle("3", "2", "2")
	util.IsTriangle("3", "4", "5")
	t.Log("决策表测试成功")
}

func TestStmtCover(t *testing.T){
	t.Log("这是三角形问题语句覆盖的结果:")
	util.IsTriangle("-2", "3", "4")
	util.IsTriangle("A", "3", "4")
	util.IsTriangle("3.1", "3", "4")
	util.IsTriangle("2", "3", "4")
	util.IsTriangle("2", "2", "3")
	util.IsTriangle("2", "2", "2")
	util.IsTriangle("", "", "")
	t.Log("语句覆盖测试通过")
}

func TestDesicionCover(t *testing.T) {
	t.Log("这是三角形问题判定覆盖的结果:")
	util.IsTriangle("50", "50", "50")
	util.IsTriangle("1", "2", "3")
	util.IsTriangle("3.14", "5", "6")
	util.IsTriangle("A", "5", "6")
	util.IsTriangle("2", "2", "3")
	util.IsTriangle("4", "5", "6")
	t.Log("问题判定覆盖测试通过")
}

func TestConditionCover(t *testing.T) {
	t.Log("这是三角形问题条件覆盖的结果:")
	util.IsTriangle("-2", "2", "4")
	util.IsTriangle("3", "-4", "120")
	util.IsTriangle("3", "5", "9")
	util.IsTriangle("8", "1", "3")
	util.IsTriangle("2", "8", "4")
	util.IsTriangle("3", "3", "3")
	util.IsTriangle("2", "3", "4")
	util.IsTriangle("2", "2", "3")
	util.IsTriangle("3", "4", "3")
	util.IsTriangle("4", "5", "5")
	t.Log("条件覆盖测试通过")
}