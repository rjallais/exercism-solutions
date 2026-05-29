package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	result := initial
    for _, v := range s {
        result = fn(result, v)
    }

    return result
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	result := initial
    for _, v := range s.Reverse() {
        result = fn(v, result)
    }

    return result
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := make(IntList, 0)
    for _, item := range s {
        if fn(item) {
            result = result.Append(IntList{item})
        }
    }

    return result
}

func (s IntList) Length() int {
	len := 0
    for _ = range s {
        len++
    }

    return len
}

func (s IntList) Map(fn func(int) int) IntList {
	result := s
    for k, v := range result {
        result[k] = fn(v)
    }

    return result
}

func (s IntList) Reverse() IntList {
	result := make(IntList, s.Length())
    for k, v := range s {
        result[len(s)-k-1] = v
    }

    return result
}

func (s IntList) Append(lst IntList) IntList {
    result := make(IntList, s.Length()+lst.Length())
    for k, v := range s {
        result[k] = v
    }
    for k, v := range lst {
        result[s.Length()+k] = v
    }

    return result
}

func (s IntList) Concat(lists []IntList) IntList {
	result := s
    for _, lst := range lists {
        result = result.Append(lst)
    }

    return result
}
