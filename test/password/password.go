package main

// 假定长度在8-16位
func checkPassword(s string) bool {
	key := make(map[int32]int)
	for _, c := range s {
		key[c]++
		if key[c] >= 3 {
			return false
		}
	}
	return true
}
