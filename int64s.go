package main

type int64s []int64

func (s int64s) Len() int           { return len(s) }
func (s int64s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s int64s) Less(i, j int) bool { return s[i] < s[j] }
