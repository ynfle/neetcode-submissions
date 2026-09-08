func characterReplacement(s string, k int) int {
	maxLen := 0
	window := map[byte]int{}
	windowStart := 0
	for i := 0; i < len(s); i++ {
		window[s[i]] += 1


		nonMaxValueSum := findCharsToChange(window)
		if nonMaxValueSum <= k {
			// if true 
			windowLen := i - windowStart + 1
			if windowLen > maxLen {
				maxLen = windowLen
			}
		} else {
			// delete from front until true
			for ; windowStart < i; {
				fmt.Println(s[i], s[windowStart])
				if window[s[windowStart]] == 1 {
					delete(window, s[windowStart])
				} else {
					window[s[windowStart]] -= 1
				}
				windowStart++
				nonMaxValueSum := findCharsToChange(window)
				if nonMaxValueSum <= k {
					break
				}
			}
		}
	}
	return maxLen
}

func findCharsToChange(window map[byte]int) int {
	// find key of max value
	maxChar := keyOfMaxValue(window)

	// other chars values should add up to k
	nonMaxValueSum := 0
	for k, v := range window {
		if k != maxChar {
			nonMaxValueSum += v
		}
	}
	return nonMaxValueSum
}

func keyOfMaxValue(m map[byte]int) byte {
	maxK := byte(0)
	maxV := 0
	for k, v := range m {
		if v > maxV {
			maxV = v
			maxK = k
		}
	}
	return maxK
}