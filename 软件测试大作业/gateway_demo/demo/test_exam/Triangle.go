package test_exam

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Triangle struct {
}

func (Triangle) IsNum (str []byte) bool {
	reg := regexp.MustCompile("^[-+]?(([0-9]+)([.]([0-9]+))?|([.]([0-9]+))?)$")
	// 返回是否只有数字
	return reg.Match(str)
}
// 判断字符串是否含有"."分隔符
func (Triangle) ContainsDot(str string) bool {
	if strings.ContainsRune(str,'.') {
		return true
	}
	return false
}

// 判断字符串是否为空
func (Triangle) IsEmpty (str string) bool {
	if len(str) == 0 {
		return true
	}
	return false
}

func (f Triangle) IsTriangle (str1,str2,str3 string) {
	// 保证三边非空且都是整数，并且都是数字
	IsNotEmpty,IsNum,NotContainDot := f.AllNotEmpty(str1,str2,str3),f.AllAreNum(str1,str2,str3),f.AllNotContainDot(str1,str2,str3)

	if IsNum && IsNotEmpty && NotContainDot {
		a,_ := strconv.Atoi(str1)
		b,_ := strconv.Atoi(str2)
		c,_ := strconv.Atoi(str3)
		if !mustFit(a,b,c) {
			fmt.Println("这三边无法构成三角形")
		} else {
			switch typeOfTriangle(a, b, c) {
			case 1:
				fmt.Println("等边三角形")
			case 2:
				fmt.Println("等腰三角形")
			case 3:
				fmt.Println("直角三角形")
			case 4:
				fmt.Println("普通三角形")
			}
		}
	} else {
		if !IsNum {
			fmt.Println("输入中有错误字符")
		} else if !IsNotEmpty {
			fmt.Println("数据不完整")
		} else {
			fmt.Println("包含小数点")
		}
	}
}
// 三角形三边条件
func mustFit(a,b,c int) bool {
	if a + b <= c || a + c <= b || b + c <=a{
		return false
	}
	return true
}
// 判定三角形类型 1 等边， 2 等腰，3 直角，4 普通
func typeOfTriangle (a,b,c int) int {
	if a == b && b == c {
		return 1
	} else if a ==b || a ==c || b ==c {
		return 2
	} else if a*a + b*b == c*c || a*a + c*c == b*b || b*b+c*c == a*a {
		return 3
	} else {
		return 4
	}
}

func (f Triangle) AllAreNum (str1,str2,str3 string) bool {
	if f.IsNum([]byte(str1)) && f.IsNum([]byte(str2)) && f.IsNum([]byte(str3)) {
		return true
	}
	return false
}

func (f Triangle) AllNotContainDot(str1,str2,str3 string) bool {
	return !(f.ContainsDot(str1) || f.ContainsDot(str2) || f.ContainsDot(str3))
}

func (f Triangle) AllNotEmpty (str1,str2,str3 string) bool {
	return !(f.IsEmpty(str1) || f.IsEmpty(str2) || f.IsEmpty(str3))
}