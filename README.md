## From
Golang doesn't have a built-in enum solution, which can make it difficult to create and manage enums in your code. While aliases can be used as a workaround, they don't provide the full functionality of an enum. That's where yet-another-enum comes in - it provides a simple and efficient way to create and work with enums in Golang.

## Example

before, you create enum like this
```go
type timerType string

const (
    every timerType = "every"
    once timerType = "once"
)

//check the string is an option of the enum
(func (val string) bool {
    _, ok := map[timerType]struct{}{every: {}, once: {}}[timerType(val)]
    return ok
})("every")
```

Now, you can just:
```go
type timerType struct {
    Every string `enum:"every"`
    On    string `enum:"on"`
}

//init and export it
var EnumList = enum.Init(&enumList{})

//use it
on := EnumList.On
val, _ := enum.ValueOf(EnumList, "on")
if on != val {
    fmt.Printf("type is mismatch %s %s\n", on.String(), val.String())
}
```

## Basic Implementation
It's uses reflection to get tag and field information from the enum struct. However, to ensure efficient runtime performance, the reflection result is cached.

## Features
- [x] string enum
- [ ] int enum
- [ ] custom enum