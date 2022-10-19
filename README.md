# Butter
Merges 2 Go structs

### Merge 2 structs & return a map
```go
str1 := struct{ Name string, Msg string }{ Name: "Joe" }
str1 := struct{ Name string, Msg string }{ Msg: "hello!" }
butter.MergeStructsToMap()

// returns
// Name -> Joe
// Description -> hello!
```