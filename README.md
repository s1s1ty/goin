# goin
[![Go Report Card](https://goreportcard.com/badge/github.com/s1s1ty/goin)](https://goreportcard.com/report/github.com/s1s1ty/goin)
[![Build Status](https://travis-ci.org/s1s1ty/goin.svg?branch=master)](https://travis-ci.org/s1s1ty/goin)
![license](https://img.shields.io/github/license/s1s1ty/goin.svg)


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
	re := goin.InInt(ar, 7)
	fmt.Println(re) // true

	seq := []float64{1.11, 3.20, 5.89, 2.90}
	re := goin.In(seq, 2.9, "float64")
	fmt.Println(re) // true

	dict := map[string]string{"name": "shaon", "id": "110"}
	re := goin.In(dict, "fullname", "[string]string")
	fmt.Println(re) // false
}
```
## Available Methods

- `InFloat64(sequence []float64, value float64)`
- `InFloat32(sequence []float32, value float32)`
- `InString(sequence []string, value string)`
- `InInt(sequence []int, value int)`
- `InInt32(sequence []int32, value int32)`
- `InInt64(sequence []int64, value int64)`
- `In(sequence type, value type, typ type)`
	* `type` will be `string`/`float64`/`float32`/`int`/`int32`/`int64`/`map`


#### License
    Licenced under MIT Licence

##### Any Suggestions and Bug Report will be gladly appreciated.