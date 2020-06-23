package NextDay

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Unit struct {}

func (Unit) IsNum (str []byte) bool {
	reg := regexp.MustCompile("^[-+]?(([0-9]+)([.]([0-9]+))?|([.]([0-9]+))?)$")
	// 返回是否只有数字
	return reg.Match(str)
}

// 判断字符串是否为空
func (Unit) IsEmpty (str string) bool {
	if len(str) == 0 {
		return true
	}
	return false
}

// 判断字符串是否含有"."分隔符
func (Unit) ContainsDot(str string) bool {
	if strings.ContainsRune(str,'.') {
		return true
	}
	return false
}
// 判断是否为闰年
func(Unit) IsLeapYear (year int) bool {
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		return true
	}
	return false
}

func(f Unit) NextDay(str1,str2,str3 string) {
	IsNotEmpty,IsNum,NotContainDot := f.AllNotEmpty(str1,str2,str3),f.AllAreNum(str1,str2,str3),f.AllNotContainDot(str1,str2,str3)
	if NotContainDot && IsNum && IsNotEmpty {
		year,_ := strconv.Atoi(str1)
		month,_ := strconv.Atoi(str2)
		day,_ := strconv.Atoi(str3)
		if year <= 2100 && year >= 1900 {
			if month <= 12 && month >= 1 {
				if day <=31 && day >= 1 {
					if f.IsLeapYear(year) && month == 2 && day == 29 {
						fmt.Println("下一日期是 ",year," 年 3 月 1 日")
					} else if !f.IsLeapYear(year) && month ==2 && day == 28 {
						fmt.Println("下一日期是 ",year," 年 3 月 1 日")
					} else if month == 12 && day == 31 {
						fmt.Println("下一日期是 ",year+1," 年 1 月 1 日")
					} else if BigMonth(month) && day == 31 {
						fmt.Println("下一日是",year," 年 ",month+1," 月 1 日")
					} else if month == 4 || month == 6 || month == 9 ||month==11 && day == 30 {
						fmt.Println("下一日是",year," 年 ",month+1," 月 1 日")
					} else if !f.IsLeapYear(year) && month == 2 && day > 28 {
						fmt.Println("本年二月没有这么多天")
					} else if month == 4 || month == 6 || month == 9 ||month==11 && day == 31 {
						fmt.Println(month, " 月没有31天")
					} else {
						day = day + 1
						fmt.Println("下一日是：",year, " 年 ", month," 月 ",day," 日")
					}
				}else {
					fmt.Println("日越界")
				}
			}else{
				fmt.Println("月越界")
			}
		}else {
			fmt.Println("年越界")
		}
	} else {
		if !IsNum {
			fmt.Println("含不合法的字符")
		} else if !NotContainDot {
			fmt.Println("含有小数点")
		} else {
			fmt.Println("数据不全，出现空字符串")
		}
	}
}

func (f Unit) AllAreNum (str1,str2,str3 string) bool {
	if f.IsNum([]byte(str1)) && f.IsNum([]byte(str2)) && f.IsNum([]byte(str3)) {
		return true
	}
	return false
}

func (f Unit) AllNotContainDot(str1,str2,str3 string) bool {
	return !(f.ContainsDot(str1) || f.ContainsDot(str2) || f.ContainsDot(str3))
}

func (f Unit) AllNotEmpty (str1,str2,str3 string) bool {
	return !(f.IsEmpty(str1) || f.IsEmpty(str2) || f.IsEmpty(str3))
}

func BigMonth(a int) bool {
	if a == 1||a==3||a==5||a==7||a==8||a==10||a==12 {
		return true
	}
	return false
}