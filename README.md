Ngram for Golang
===========

| Branch  | Status |
| ------- | ------ |
| master  | [![Build Status](https://travis-ci.org/euskadi31/go-ngram.svg?branch=master)](https://travis-ci.org/euskadi31/go-ngram)  |
| develop | [![Build Status](https://travis-ci.org/euskadi31/go-ngram.svg?branch=develop)](https://travis-ci.org/euskadi31/go-ngram) |


> an **_n_-gram** is a contiguous sequence of _n_ items from a given sequence of text or speech.


~~~go
import (
    "fmt"

    "github.com/euskadi31/go-ngram"
)

arr := ngram.Ngram("Hello", ngram.Bigram)

fmt.Printf("%v\n", arr) // ["He", "el", "ll", "lo"]
~~~

## License

go-ngram is licensed under [the MIT license](LICENSE.md).
