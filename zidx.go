package zidx

func ZIndex(haystack, needle string) int {
	n := len(haystack)
	m := len(needle)
	if m == 0 {
		return 0
	}

	if n == 0 || m > n {
		return -1
	}

	lenZ := n + m
	z := ZFunc(needle, lenZ)
	li := 0
	ri := 0

	for i := m; i < lenZ; i++ {
		if i <= ri {
			x := ri - i
			if z[i-li] < x {
				z[i] = z[i-li]
			} else {
				z[i] = x
			}
		}

		k := i - m
		for k+z[i] < n && needle[z[i]] == haystack[k+z[i]] {
			z[i]++
			if z[i] == m {
				return k
			}
		}

		if i+z[i] > ri {
			li = i
			ri = i + z[i]
		}
	}

	return -1
}

func ZFunc(s string, lenZ int) []int {
	li := 0
	ri := 0
	z := make([]int, lenZ)
	for i := 1; i < len(s); i++ {
		if i <= ri {
			x := ri - i
			if z[i-li] < x {
				z[i] = z[i-li]
			} else {
				z[i] = x
			}
		}
		for i+z[i] < len(s) && s[z[i]] == s[i+z[i]] {
			z[i]++
		}

		if i+z[i] > ri {
			li = i
			ri = i + z[i]
		}
	}
	return z
}