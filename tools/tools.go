package tools

import "strconv"

func RemoveFromStr(s[]string , index int)[]string{
	s = append(s[:index], s[index+1:]...)
	x := len(s)
	s[x-1] = ""
	return s
}

func Inttostr(id uint64)string{
	str := strconv.FormatUint(id, 10)
	return str
}


