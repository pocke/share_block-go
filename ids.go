package main

type IDs []int64

func (s IDs) Len() int           { return len(s) }
func (s IDs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IDs) Less(i, j int) bool { return s[i] < s[j] }
