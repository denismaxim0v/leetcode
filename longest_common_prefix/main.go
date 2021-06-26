package longest_common_prefix

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	// sort so the shortest string is first in the array
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	for i := 0; i <= len(strs[0])-1; i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[0]) == 0 || len(strs[j]) == 0 {
				return ""
			}
			if strs[0][i] != strs[j][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}
