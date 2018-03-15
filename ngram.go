// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ngram

// Ngram analysis is a field in textual analysis which uses sliding window character
// sequences in order to aid topic analysis.
func Ngram(content string, size SizeType) []string {
	container := []string{}
	width := int(size)
	var ng []byte

	for i := 0; i < len(content); i++ {
		if i > (width - 2) {

			for j := (width - 1); j >= 0; j-- {
				ng = append(ng, content[i-j])
			}

			container = append(container, string(ng))

			ng = []byte{}
		}
	}

	return container
}
