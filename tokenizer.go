// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ngram

import "unicode/utf8"

// Tokenizer interface
type Tokenizer interface {
	Tokenize() []string
}

type tokenizer struct {
	size    SizeType
	content string
}

// New N-gram Tokenizer
func New(content string, size SizeType) Tokenizer {
	return &tokenizer{
		size:    size,
		content: content,
	}
}

func (t tokenizer) Tokenize() []string {
	length := len(t.content)
	size := int(t.size)
	runes := make([]int, length+1)
	ridx := 0 // rune index
	bidx := 0 // byte index

	for i, w := 0, 0; i < length; i += w {
		_, width := utf8.DecodeRuneInString(t.content[bidx:])
		runes[ridx] = bidx
		bidx += width
		w = width
		ridx++
	}

	runes[ridx] = len(t.content)

	tokens := make([]string, ridx+1-size)

	for i := 0; i <= ridx-size; i++ {
		end := i + size

		tokens[i] = t.content[runes[i]:runes[end]]
	}

	return tokens
}
