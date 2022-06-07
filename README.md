# Ngram for Golang ![Last release](https://img.shields.io/github/release/euskadi31/go-ngram.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/euskadi31/go-ngram)](https://goreportcard.com/report/github.com/euskadi31/go-ngram)

| Branch | Status                                                                                                                                            | Coverage                                                                                                                                     |
| ------ | ------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| master | [![Go](https://github.com/euskadi31/go-ngram/actions/workflows/go.yml/badge.svg)](https://github.com/euskadi31/go-ngram/actions/workflows/go.yml) | [![Coveralls](https://img.shields.io/coveralls/euskadi31/go-ngram/master.svg)](https://coveralls.io/github/euskadi31/go-ngram?branch=master) |

> an **_n_-gram** is a contiguous sequence of _n_ items from a given sequence of text or speech.

```go
import (
    "fmt"

    "github.com/euskadi31/go-ngram"
)

func main() {
    tokens := ngram.New(ngram.Bigram).Tokenize("Hello")

    fmt.Printf("%v\n", tokens) // ["He", "el", "ll", "lo"]
}

```

## License

go-ngram is licensed under [the MIT license](LICENSE.md).
