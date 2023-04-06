# Yet Another Enum

### Why
Golang doesn't have a built-in enum solution, which can make it difficult to create and manage enums in your code. While aliases can be used as a workaround, they don't provide the full functionality of an enum. That's where yet-another-enum comes in - it provides a simple and efficient way to create and work with enums in Golang.

### Show me the code
At before, we create enum like this
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
type Enums struct {
    Every *yaenum.Instance[Enums] `enum:"every"`
    On    *yaenum.Instance[Enums] `enum:"on"`
}

//init and export it
var EnumList = yaenum.Init[Enums](&Enums{})

//use it
on := EnumList.On
val, _ := yaenum.ValueOf[type_enum.Enums](type_enum.EnumList, "on")
if on != val {
    fmt.Printf("type is mismatch %s %s\n", on.String(), val.String())
}
```

### Basic Implementation
It's uses reflection to get tag and field information from the enum struct. However, to ensure efficient runtime performance, the reflection result is cached.

### Features
- [x] string enum
- [ ] int enum
- [ ] custom enum