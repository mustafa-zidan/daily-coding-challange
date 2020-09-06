package strings

// Given a word w and string s, find all indecies in s which are the starting locations
// of anagrams of w.

func findAnagramIndicies(s, w string) []int {
	if len(s) < len(w) {
		return nil
	}
	wHash, windowHash, result := 0, 0, make([]int, 0)
	// calculate w hash
	for _, c := range w {
		wHash += 1 << uint(c-'a')
	}
	start, end, window := 0, 0, len(w)-1
	for end < len(s) {
		windowHash += 1 << (s[end] - 'a')
		//keep accumelating till you reach the window size
		if end-start < window {
			end++
			continue
		}

		if windowHash == wHash {
			result = append(result, start)
		}
		windowHash -= 1 << uint(s[start]-'a')
		start++
		end++
	}
	return result
}
