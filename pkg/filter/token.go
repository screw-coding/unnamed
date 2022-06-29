package filter

import "strings"

//
// Lowercase
// @Description: 将字符串转为小写
// @param tokens
// @return []string
//
func Lowercase(tokens []string) []string {
	result := make([]string, len(tokens))
	for i, token := range tokens {
		result[i] = strings.ToLower(token)
	}
	return result
}

//
// StopWord
// @Description: 去除停用词
// @param tokens
// @return []string
//
func StopWord(tokens []string) []string {
	var stopWords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {}, "in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}
	result := make([]string, 0, len(tokens))
	for _, token := range tokens {
		_, ok := stopWords[token]
		if !ok {
			result = append(result, token)
		}
	}
	return result
}
