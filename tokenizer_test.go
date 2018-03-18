// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ngram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizerUnigram(t *testing.T) {
	assert.Equal(
		t,
		[]string{"H", "e", "l", "l", "o"},
		New(Unigram).Tokenize("Hello"),
	)

	assert.Equal(
		t,
		[]string{"こ", "ん", "に", "ち", "は"},
		New(Unigram).Tokenize("こんにちは"),
	)
}

func TestTokenizerBigram(t *testing.T) {
	assert.Equal(
		t,
		[]string{"He", "el", "ll", "lo"},
		New(Bigram).Tokenize("Hello"),
	)

	assert.Equal(
		t,
		[]string{"こん", "んに", "にち", "ちは"},
		New(Bigram).Tokenize("こんにちは"),
	)
}

func TestTokenizerTrigram(t *testing.T) {
	assert.Equal(
		t,
		[]string{"Hel", "ell", "llo"},
		New(Trigram).Tokenize("Hello"),
	)

	assert.Equal(
		t,
		[]string{"こんに", "んにち", "にちは"},
		New(Trigram).Tokenize("こんにちは"),
	)
}

func BenchmarkTokenizerUnigram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(Unigram).Tokenize(corpus)
	}
}

func BenchmarkTokenizerBigram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(Bigram).Tokenize(corpus)
	}
}

func BenchmarkTokenizerTrigram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(Trigram).Tokenize(corpus)
	}
}
