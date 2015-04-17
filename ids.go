package main

import "strconv"

type IDs []int64

func (s IDs) Len() int           { return len(s) }
func (s IDs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IDs) Less(i, j int) bool { return s[i] < s[j] }

func (l IDs) Diff(r IDs) {
	if len(r) == 0 {
		return
	}

	i := 0
	for _, rv := range r {
		for j := i; j < len(l); j++ {
			if l[j] == rv {
				copy(l[j:], l[j+1:])
				l = l[:len(l)-1]
				i = j
				break
			} else if l[j] > rv {
				i = j - 1
				break
			}
		}
	}
}

func (l IDs) Eq(r IDs) bool {
	if len(l) != len(r) {
		return false
	}
	for i, lv := range l {
		if lv != r[i] {
			return false
		}
	}
	return true
}

func (s IDs) String() string {
	if len(s) == 0 {
		return "[]"
	}

	res := make([]byte, 0, len(s)*3)
	res = append(res, '[')
	for _, v := range s {
		res = strconv.AppendInt(res, v, 10)
		res = append(res, ", "...)
	}
	res = res[:len(res)-2]
	res = append(res, ']')
	return string(res)
}
