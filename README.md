# goin
[![Go Report Card](https://goreportcard.com/badge/github.com/s1s1ty/goin)](https://goreportcard.com/report/github.com/s1s1ty/goin)
[![Build Status](https://travis-ci.org/s1s1ty/goin.svg?branch=master)](https://travis-ci.org/s1s1ty/goin)
[![license](https://img.shields.io/github/license/s1s1ty/goin.svg)](https://github.com/s1s1ty/goin/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/s1s1ty/goin?status.svg)](https://godoc.org/github.com/s1s1ty/goin)

**goin** Evaluates to true if it finds a variable in the specified sequence and false otherwise.

* **goin** package is built inspired by **python** **in** operator

## Installation and Usage

Install the package with:
```
go get github.com/s1s1ty/goin
```
Import it with:
```go
import "github.com/s1s1ty/goin"
```

## Quick Start

```go
func main() {
	ar := []int{1, 2, 4, 7, 8, 3}
	found, _ := goin.NewValue(7).In(ar)
	fmt.Println(found) // true

	seq := []float64{1.11, 3.20, 5.89, 2.90}
	found, _ = goin.NewValue(2.9).In(seq)
	fmt.Println(found) // true

	dict := map[string]string{"name": "shaon", "id": "110"}
	found, _ = goin.NewValue("fullname").InKey(dict)
	fmt.Println(found) // false
}
```
## Available Methods

- `In(arr interface{})`
- `InKey(arr interface{})`

#### License
    Licenced under MIT Licence

##### Any Suggestions and Bug Report will be gladly appreciated.
