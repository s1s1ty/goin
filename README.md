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
import "gopkg.in/s1s1ty/goin"
```

## Quick Start

```go
func main() {
	ar := []int{1, 2, 4, 7, 8, 3}
	re := goin.IntIn(ar, 7)
	fmt.Println(re)
}
```
## Available Methods

- `Float64In(ar []float64, value float64)`
- `StringIn(ar []string, value string)`
- `IntIn(ar []int, value int)`


#### License
    Licenced under MIT Licence

##### Any Suggestions and Bug Report will be gladly appreciated.