// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ngram

// SizeType type.
type SizeType int

// N-Gram size type.
const (
	Unigram SizeType = 1
	Bigram  SizeType = 2
	Trigram SizeType = 3
)
