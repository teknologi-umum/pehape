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
 insertionCost := 1
 replacementCost := 2
 deletionCost := 3
 fmt.Println(pehape.Levenshtein(str1, str2, insertionCost, replacementCost, deletionCost))
 //result : 5
 ```
 * `Add Slashes`
  - sample string
```go
pehape.AddSlashes(`What does "yolo" mean?`)
//result : What does \"yolo\" mean?
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
fmt.Println(pehape.Ucwords(bar))
//result : "Hello World!"
```

 * `Strrev`
 ```go
 var foo = "Hello World!"
 fmt.Println(pehape.Strrev(foo))
 //result : "!dlroW olleH"

 var bar = "kasur rusak"
 fmt.Println(pehape.Strrev(bar))
 //result : "kasur rusak"
 ```
 
