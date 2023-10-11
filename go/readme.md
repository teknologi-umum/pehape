## Import

`import "github.com/teknologi-umum/pehape/go"`

## Usage?

### `Levenshtein`
 
 ```go
 str1 := "kitten"
 str2 := "sitting"
 fmt.Println(pehape.Levenshtein(str1, str2))
 //result : 3
 ```
 - with custom cost
 ```go
 str1 := "kitten"
 str2 := "sitting"
 insertionCost := 1
 replacementCost := 2
 deletionCost := 3
 fmt.Println(pehape.Levenshtein(str1, str2, insertionCost, replacementCost, deletionCost))
 //result : 5
 ```
 
### `Add Slashes`
 
- sample string
  
```go
pehape.AddSlashes(`What does "yolo" mean?`)
//result : What does \"yolo\" mean?
```

### `Implode`

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

### `Explode`

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

### `Ucwords`

```go
var bar = "HELLO WORLD!"
fmt.Println(pehape.Ucwords(bar))
//result : "Hello World!"
```

### `Bin2Hex`

```go
var str = "Hello World!!"
fmt.Println(pehape.Bin2Hex(str))
//result : "48656c6c6f20576f726c642121"
```

### `Hex2Bin`

```go
var str = "48656c6c6f20576f726c642121"
fmt.Println(pehaphe.Hex2Bin(str))
//result : "Hello World!!"
```

### `LTrim`

- with no chars specified

```go
var str = "	\n\x0BHello World"
res, err := pehape.LTrim(str)
if err != nil {
  panic(err)
}
fmt.Println(res)
//result : "Hello World"
```

- with specified chars

```go
var str = "123456\t    \n\x0BHello World"
res, err := pehape.LTrim(str, "123456", " ", "\t\n\x0B")
if err != nil {
  panic(err)
}
fmt.Println(res)
//result: "Hello World"
```

- with specified char range

```go
var str = "abcdefghijklmnopqrstuvwxyzHELLO WORLD!"
res, err := pehape.LTrim(str, "a..z")
if err != nil {
  panic(err)
}
fmt.Println(res)
//result: "HELLO WOLRD!"
```

### `RTrim`

- with no chars specified

```go
var str = "Hello World	\n\x0B"
res, err := pehape.RTrim(str)
if err != nil {
  panic(err)
}
fmt.Println(res)
//result : "Hello World"
```

- with specified chars

```go
var str = "Hello World123456\t    \n\x0B"
res, err := pehape.RTrim(str, "123456", " ", "\t\n\x0B")
if err != nil {
  panic(err)
}
fmt.Println(res)
//result : "Hello World"
```

- with specified char range

```go
var str = "HELLO WORLD!abcdefghijklmnopqrstuvwxyz"
res, err := pehape.RTrim(str, "a..z")
if err != nil {
  panic(err)
}
fmt.Println(res)
//result : "HELLO WORLD!"
```

### `Strrev`

```go
var foo = "Hello World!"
fmt.Println(pehape.Strrev(foo))
//result : "!dlroW olleH"

var bar = "kasur rusak"
fmt.Println(pehape.Strrev(bar))
//result : "kasur rusak"
```

### `Chr`

```go
var charFromNumber = pehape.Chr(65)
fmt.Println(charFromNumber)
//result : "A"

var charFromOctal = pehape.Chr(053)
fmt.Println(charFromOctal)
//result: "+"

var charFromHex = pehape.Chr(0x52)
fmt.Println(charFromHex)
//result: "R"

var symbol = pehape.Chr(240) + pehape.Chr(159) + pehape.Chr(144) + pehape.Chr(152)
fmt.Println(symbol)
//result : "🐘"
```

### `StrStartsWith`

```go
var result = pehape.StrStartsWith("Hello World", "Hell")
fmt.Println(result)
//result : true

var result2 = pehape.StrStartsWith("Hello World", "World")
fmt.Println(result2)
//result : false

### `StrPad`

```go
input := "Alien";
fmt.Println(StrPad(input, 10));                     // produces "Alien     "
fmt.Println(StrPad(input, 10, "-=", STR_PAD_LEFT)); // produces "-=-=-Alien"
fmt.Println(StrPad(input, 10, "_", STR_PAD_BOTH));  // produces "__Alien___"
fmt.Println(StrPad(input,  6, "___"));              // produces "Alien_"
fmt.Println(StrPad(input,  3, "*"));                // produces "Alien"
```

### `StrShuffle`

```go
shuffle := StrShuffle("abcdef")

//This will print something like: "bfdaec"
fmt.Println(shuffle)
```