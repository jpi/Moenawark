package namegen

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
)

type Markov struct {
	depth int
	dict  map[string][]rune
}

func NewMarkov(file io.Reader, depth int) (m *Markov, err error) {
	if depth < 2 {
		err = fmt.Errorf("NewMarkov: depth must greater or equal to 2")
		return
	}

	reader := bufio.NewReader(file)
	var line string
	dict := make(map[string][]rune)
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return
		}
		name := []rune(line)
		prefix := make([]rune, 0)
		for _, r := range name {
			// Append r to the list of possible characters following prefix
			rest, ok := dict[string(prefix)]
			if !ok {
				rest = make([]rune, 0)
			}
			rest = append(rest, r)
			dict[string(prefix)] = rest
			if len(prefix) < depth {
				prefix = append(prefix, r)
			} else {
				// Rotate prefix: append r and discard first character
				for i := 1; i < len(prefix); i += 1 {
					prefix[i-1] = prefix[i]
				}
				prefix[depth-1] = r
			}
		}
		if err == io.EOF {
			err = nil
			break
		}
	}

	m = &Markov{depth, dict}
	return
}

func (m *Markov) Gen(n int) (name string) {
	letters := make([]rune, n)
	for x := 0; x < 50; x += 1 {
		for i := 0; i < n; i += 1 {
			r := m.randomNextLetter(letters)
			if r == 0 {
				break
			}
			letters[i] = r
		}
		if m.isWordEnding(letters) {
			break
		}
	}
	return string(letters)
}

func (m *Markov) nextLetters(letters []rune) []rune {
	var prefix []rune
	if len(letters) < m.depth {
		prefix = letters[:]
	} else {
		prefix = letters[len(letters)-m.depth:]
	}
	return m.dict[string(prefix)]
}

func (m *Markov) isWordEnding(letters []rune) bool {
	next := m.nextLetters(letters)
	for _, r := range next {
		if r == '\n' {
			return true
		}
	}
	return false
}

func (m *Markov) randomNextLetter(letters []rune) rune {
	next := m.nextLetters(letters)
	if len(next) == 0 {
		return 0
	}
	i := rand.Intn(len(next))
	return next[i]
}
