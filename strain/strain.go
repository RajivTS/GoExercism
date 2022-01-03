package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var result Ints
	for _, item := range i {
		if filter(item) {
			result = append(result, item)
		}
	}
	return result
}

func (i Ints) Discard(filter func(int) bool) Ints {
	return i.Keep(func(i int) bool { return !filter(i) })
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var result Lists
	for _, list := range l {
		if filter(list) {
			result = append(result, list)
		}
	}
	return result
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var result Strings
	for _, str := range s {
		if filter(str) {
			result = append(result, str)
		}
	}
	return result
}
