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
		New("Hello", Unigram).Tokenize(),
	)

	assert.Equal(
		t,
		[]string{"こ", "ん", "に", "ち", "は"},
		New("こんにちは", Unigram).Tokenize(),
	)
}

func TestTokenizerBigram(t *testing.T) {
	assert.Equal(
		t,
		[]string{"He", "el", "ll", "lo"},
		New("Hello", Bigram).Tokenize(),
	)

	assert.Equal(
		t,
		[]string{"こん", "んに", "にち", "ちは"},
		New("こんにちは", Bigram).Tokenize(),
	)
}

func TestTokenizerTrigram(t *testing.T) {
	assert.Equal(
		t,
		[]string{"Hel", "ell", "llo"},
		New("Hello", Trigram).Tokenize(),
	)

	assert.Equal(
		t,
		[]string{"こんに", "んにち", "にちは"},
		New("こんにちは", Trigram).Tokenize(),
	)
}

func BenchmarkTokenizerUnigram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(corpus, Unigram).Tokenize()
	}
}

func BenchmarkTokenizerBigram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(corpus, Bigram).Tokenize()
	}
}

func BenchmarkTokenizerTrigram(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(corpus, Trigram).Tokenize()
	}
}
