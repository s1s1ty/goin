# goin
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
}
```
## Available Methods

- `InFloat64(sequence []float64, value float64)`
- `InString(sequence []string, value string)`
- `InInt(sequence []int, value int)`
- `In(sequence type, value type, typ type)`
	* `type` will be `string` or `int` or `float64`


#### License
    Licenced under MIT Licence

##### Any Suggestions and Bug Report will be gladly appreciated.