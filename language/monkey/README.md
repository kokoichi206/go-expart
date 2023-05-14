[Go言語でつくるインタプリタ](https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%81%A4%E3%81%8F%E3%82%8B%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF-Thorsten-Ball/dp/4873118220)で勉強させてもらった内容です。

## Usage

``` sh
$ go run main.go
Hello kokoichi! This is the Monkey programming language!
>> let a = {"name": "kokoichi", "age": 99 }                            
>> a["name"]
kokoichi
>> puts(a)
{name: kokoichi, age: 99}
null
>> let add = fn(x, y) { x + y; };
>> add(46, 88)                                
134
>> let arr = [1, "pien", fn(x) { x * x }]
>> arr[2](3)   
9
>> arr[-1]
null
```
