## Import
`import "github.com/teknologi-umum/pehape/go"`

## Usage?
 * `Implode`
  - sample string
```go
var array = []string{"Hello", "world"}
fmt.Println(pehape.Implode(array, " "))
//result : Hello world
```
  - sample number
```go
var array = []int{1, 2, 3, 4, 5}
fmt.Println(pehape.Implode(array, " "))
//result : 1 2 3 4 5
```

 * `Explode`
  - With No Limit
 ```go
 var str string = "Hello pehape world"
 fmt.Println(pehape.Explode(" ", str))

 //result : [Hello pehape world]
 ```
  - With Limit
```go
 var str string = "Hello pehape world"
 fmt.Println(pehape.Explode(" ", str, 2))
 //result : [Hello pehape]

 fmt.Println(pehape.Explode(" ", str, -2))
 //result : [Hello]
 ```
 
 * `Ucwords`
```go
var bar = "HELLO WORLD!"
fmt.Println(pehape.Ucwords(foo))
//result : "Hello World!"
```