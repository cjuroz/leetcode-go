# LeetCode Go

A collection of LeetCode-style problem solutions implemented in Go, organized by topic and problem number.

## Overview

This repository contains focused Go implementations and tests for common algorithm and data-structure problems from LeetCode. Each problem is stored in its own folder with a `solution.go` file and usually a matching `solution_test.go`.

## Project Structure

```text
leetcode-go/
├── arrays/
│   ├── 0001_two_sum/
│   ├── 0049_group_anagrams/
│   ├── 0169_majority_element/
│   ├── 0217_contains_duplicate/
│   ├── 0242_valid_anagram/
│   ├── 0347_top_k_frequent_elements/
│   └── ...
├── two_pointers/
│   ├── 0125_valid_palindrome/
│   └── 0167_two_sum_ii/
├── go.mod
├── README.md
└── ...
```

## Example

```go
package validpalindrome

func IsPalindrome(s string) bool {
	filtered := make([]byte, 0, len(s))

	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			filtered = append(filtered, s[i]+32)
		} else if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			filtered = append(filtered, s[i])
		}
	}

	for i, j := 0, len(filtered)-1; i < j; i, j = i+1, j-1 {
		if filtered[i] != filtered[j] {
			return false
		}
	}

	return true
}
```

## Running Tests

From the repository root:

```bash
go test ./...
```

This will run the test suite for all implemented problems.

## Notes

- Solutions are kept small and problem-focused.
- Folder names follow the LeetCode numbering and problem title.
- Each problem includes at least one set of unit tests covering expected behavior.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
