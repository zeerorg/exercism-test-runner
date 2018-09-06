package main

func (arr StrArray) isPresent(check string) bool {
	for _, ele := range arr {
		if ele == check {
			return true
		}
	}
	return false
}
