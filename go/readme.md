## Import
`import "github.com/teknologi-umum/pehape/go"`

## Usage?
 * `Levenshtein`
 ```go
 str1 := "kitten"
 str2 := "sitting"
 fmt.Println(pehape.Levenshtein(str1, str2))
 //result : 3
 ```
 with custom cost
 ```go
 str1 := "kitten"
 str2 := "sitting"
 insertion_cost := 1
 replacement_cost := 2
 deletion_cost := 3
 fmt.Println(pehape.Levenshtein(str1,str2,insertion_cost,replacement_cost,deletion_cost))
 //result : 5
 ```

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
