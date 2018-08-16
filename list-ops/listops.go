package listops

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// IntList is []int type
type IntList []int

// Foldr it return fold right of IntList
func (list IntList) Foldr(fn binFunc, initial int) int {
	fold := initial
	for i := range list {
		fold = fn(list[len(list)-1-i], fold)
	}
	return fold
}

// Foldl it return fold left of IntList
func (list IntList) Foldl(fn binFunc, initial int) int {
	fold := initial
	for _, v := range list {
		fold = fn(fold, v)
	}
	return fold
}

// Filter it return filter of IntList
func (list IntList) Filter(fn predFunc) IntList {
	l := make(IntList, 0)
	for _, v := range list {
		if fn(v) {
			l = append(l, v)
		}
	}
	return l
}

func (list IntList) Length() int {
	return len(list)
}

// Map it return map of IntList
func (list IntList) Map(fn unaryFunc) IntList {
	l := make(IntList, len(list))
	for i, v := range list {
		l[i] = fn(v)
	}
	return l
}

// Reverse it return reverse of IntList
func (list IntList) Reverse() IntList {
	l := make(IntList, len(list))
	for i := range list {
		l[i] = list[len(list)-1-i]
	}
	return l
}

// Append it return append of IntList
func (list IntList) Append(append IntList) IntList {
	l := make(IntList, len(list)+len(append))
	i := 0
	for _, v := range list {
		l[i] = v
		i++
	}
	for _, v := range append {
		l[i] = v
		i++
	}
	return l
}

// Concat it return concat of IntList
func (list IntList) Concat(lists []IntList) IntList {
	l := make(IntList, 0)
	l = l.Append(list)
	for _, other := range lists {
		l = l.Append(other)
	}
	return l
}
