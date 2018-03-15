// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ngram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnigram(t *testing.T) {
	actual := Ngram("Hello", Unigram)

	assert.Equal(t, []string{"H", "e", "l", "l", "o"}, actual)
}

func TestBigram(t *testing.T) {
	actual := Ngram("Hello", Bigram)

	assert.Equal(t, []string{"He", "el", "ll", "lo"}, actual)
}

func TestTrigram(t *testing.T) {
	actual := Ngram("Hello", Trigram)

	assert.Equal(t, []string{"Hel", "ell", "llo"}, actual)
}

func BenchmarkUnigram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Ngram(corpus, Unigram)
	}
}

func BenchmarkBigram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Ngram(corpus, Bigram)
	}
}

func BenchmarkTrigram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Ngram(corpus, Trigram)
	}
}
