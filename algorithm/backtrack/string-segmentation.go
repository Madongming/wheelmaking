package backtrack

func partition(str string) [][]string {
	if str == "" {
		return [][]string{}
	}

	return partitionRec(str, 0, nil, nil)
}

func partitionRec(str string, start int, result [][]string, elem []string) [][]string {
	if start >= len(str) {
		return append(result, elem)
	}

	for end := start; end < len(str); end++ {
		elem = append(elem, str[start:end+1])
		result = partitionRec(str, end+1, result, elem)
		tmp := make([]string, len(elem)-1)
		copy(tmp, elem[0:len(elem)-1])
		elem = tmp
	}
	return result
}
