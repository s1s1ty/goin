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
	fmt.Println(re)
}
```
## Available Methods

- `InFloat64(ar []float64, value float64)`
- `InString(ar []string, value string)`
- `InInt(ar []int, value int)`
- `In(ar []interface{}, value interface{}, typ string)`


#### License
    Licenced under MIT Licence

##### Any Suggestions and Bug Report will be gladly appreciated.