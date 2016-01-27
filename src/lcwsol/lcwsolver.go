package lcwsolver

import (
	"bufio"
	"os"
)

// Longest Compound Word Solver
type LcwSolver struct {
	words []string
	ht    map[string]bool
}

func New() *LcwSolver {
	s := LcwSolver{}
	s.words = make([]string, 0)
	s.ht = make(map[string]bool)
	return &s
}

func (s *LcwSolver) LoadWordList(file_name string) error {
	file, err := os.Open(file_name)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		s.words = append(s.words, word)
		s.ht[word] = true
	}

	return scanner.Err()
}

func (s *LcwSolver) AppendWord(word string) {
	s.words = append(s.words, word)
	s.ht[word] = true
}

func (s *LcwSolver) getPrefixes(candidate string, start_len int, end_len int) []string {
	var prefixes []string

	for length := start_len; length <= end_len; length++ {
		prefix := candidate[0:length]

		if _, ok := s.ht[prefix]; ok {
			prefixes = append(prefixes, prefix)
		}
	}

	return prefixes
}

func (s *LcwSolver) isCompoundWord(candidate string, count *int) bool {
	prefixes := s.getPrefixes(candidate, 2, len(candidate))
	found := false

	for _, prefix := range prefixes {
		*count++

		suffix := candidate[len(prefix):]
		if len(suffix) == 0 && *count > 1 {
			return true
		}

		found = s.isCompoundWord(suffix, count)
		if found {
			return true
		}

		*count--
	}

	return found
}

func (s *LcwSolver) GetResult() string {
	var longest_word string
	var count int

	for _, candidate := range s.words {
		count = 0

		if len(candidate) < len(longest_word) {
			continue
		}

		if found := s.isCompoundWord(candidate, &count); found {
			if len(candidate) > len(longest_word) {
				longest_word = candidate
			}
		}
	}

	return longest_word
}
