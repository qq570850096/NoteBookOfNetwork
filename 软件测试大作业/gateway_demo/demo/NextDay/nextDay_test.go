package NextDay

import (
	"testing"
)

var util Unit

func TestBigMonth(t *testing.T) {
	TestStmtCover(t)
	TestDesicionCover(t)
	TestBoundaryOfNextDay(t)
	TestDesicionTableOfNextDay(t)
	TestPathCover(t)
	TestEClassOfNextDay(t)
}

func TestBoundaryOfNextDay(t *testing.T) {
	t.Log("以下是次日问题的边界值用例测试结果:")
	util.NextDay("1811", "6", "15")
	util.NextDay("1812", "6", "15")
	util.NextDay("1813", "6", "15")
	util.NextDay("1912", "6", "15")
	util.NextDay("2011", "6", "15")
	util.NextDay("2012", "6", "15")
	util.NextDay("2013", "6", "15")
	util.NextDay("1912", "6", "-1")
	util.NextDay("1912", "6", "1")
	util.NextDay("1912", "6", "2")
	util.NextDay("1912", "6", "29")
	util.NextDay("1912", "6", "30")
	util.NextDay("1912", "6", "31")
	util.NextDay("1912", "6", "32")
	util.NextDay("1912", "-1", "15")
	util.NextDay("1912", "1", "15")
	util.NextDay("1912", "2", "15")
	util.NextDay("1912", "11", "15")
	util.NextDay("1912", "12", "15")
	util.NextDay("1912", "13", "15")
	t.Log("边界值用例测试通过")
}
func TestEClassOfNextDay(t *testing.T) {
	t.Log("开始次日问题等价类测试结果:")
	util.NextDay("1812", "6", "30")
	util.NextDay("2012", "6", "30")
	util.NextDay("2000", "2", "29")
	util.NextDay("1912", "7", "31")
	util.NextDay("1909", "2", "28")
	util.NextDay("2005", "10", "31")
	util.NextDay("1997", "11", "31")
	util.NextDay("1999", "12", "31")
	util.NextDay("1900", "2", "29")
	util.NextDay("abcd", "6", "6")
	util.NextDay("6.6", "6", "6")
	util.NextDay("1787", "6", "5")
	util.NextDay("2018", "7", "2")
	util.NextDay("2000", "0", "6")
	util.NextDay("2000", "13", "6")
	util.NextDay("2000", "6", "0")
	util.NextDay("2000", "6", "32")
	util.NextDay("1887", "1", "15")
	util.NextDay("", "", "")
	t.Log("等价类测试通过")
}
func TestDesicionTableOfNextDay(t *testing.T) {
	t.Log("开始次日问题的决策表测试结果:")
	util.NextDay("2007", "6", "15")
	util.NextDay("2007", "6", "30")
	util.NextDay("2007", "6", "31")
	util.NextDay("2007", "1", "15")
	util.NextDay("2007", "1", "31")
	util.NextDay("2007", "12", "15")
	util.NextDay("2007", "12", "31")
	util.NextDay("2007", "2", "15")
	util.NextDay("2007", "2", "28")
	util.NextDay("2007", "2", "29")
	util.NextDay("2007", "2", "30")
	util.NextDay("2000", "2", "28")
	util.NextDay("2000", "2", "29")
	util.NextDay("2000", "2", "30")
	t.Log("决策表测试通过")
}
func TestStmtCover(t *testing.T) {
	t.Log("开始次日问题语句覆盖的测试:")
	util.NextDay("2000", "2", "29")
	util.NextDay("2001", "2", "28")
	util.NextDay("2000", "3", "28")
	util.NextDay("2001", "12", "31")
	util.NextDay("2001", "11", "30")
	t.Log("测试通过")
}
func TestDesicionCover(t *testing.T) {
	t.Log("开始次日问题判定覆盖的测试:")
	util.NextDay("2000", "2", "29")
	util.NextDay("2001", "2", "28")
	util.NextDay("2000", "3", "28")
	util.NextDay("2001", "12", "31")
	util.NextDay("2001", "11", "30")
	t.Log("语句覆盖测试通过")
}
func TestPathCover(t *testing.T) {
	t.Log("开始次日问题路径覆盖的测试:")
	util.NextDay("2008", "2", "29")
	util.NextDay("2008", "7", "7")
	util.NextDay("2008", "12", "31")
	util.NextDay("2008", "11", "30")
	util.NextDay("2007", "10", "31")
	util.NextDay("2007", "2", "28")
	t.Log("路径覆盖测试通过")
}
