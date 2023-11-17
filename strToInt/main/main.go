package main

import (
	"fmt"
	"math"
)

func strToInt(str string) int {
	sign := isSigned(str)
	b := strToByte(sign, str)
	return byteToInt(b)
}
func isSigned(str string) int {
	i := 0
	/*
		判断是否是数字且有符号
		sign==0为'-'
		sign==1为'+'
		sign==-1为非数字
	*/
	sign := 0
	for ; i < len(str); i++ {
		if str[i] == '-' || str[i] == '+' {
			sign = 1
			break
		} else if str[i] >= '0' && str[i] <= '9' {
			sign = 0
			break
		} else if str[i] != ' ' {
			sign = -1
			break
		}
	}
	return sign
}
func strToByte(sign int, str string) string {
	var b []byte
	i := 0
	if sign == 0 {
		for ; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				b = append(b, str[i])
				i++
				break
			}
		}
		for ; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				b = append(b, str[i])
			} else {
				break
			}
		}
	} else if sign == 1 {
		for ; i < len(str); i++ {
			if str[i] == '-' || str[i] == '+' {
				b = append(b, str[i])
				i++
				break
			}
		}
		for ; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				b = append(b, str[i])
			} else {
				break
			}
		}
	} else {
		return ""
	}
	return string(b)
}
func byteToInt(b string) int {
	var x int = 0
	/*
		在最后辨别符号
		flag==1为'+'
		flag==0为'-'
	*/
	flag := 1
	for k, v := range b {
		if k == 0 {
			if v == '+' {
				flag = 1
				continue
			} else if v == '-' {
				flag = 0
				continue
			} else {
				x = int(v - '0')
			}
		} else {
			if x*10+int(v-'0') > math.MaxInt32 {
				if flag == 0 {
					x = -math.MinInt32
					break
				} else {
					x = math.MaxInt32
					break
				}
			} else {
				x = x*10 + int(v-'0')
			}
		}
	}
	if flag == 0 {
		return -x
	} else {
		return x
	}
}
func main() {
	n := strToInt("  -3.14")
	fmt.Println(n)
}
