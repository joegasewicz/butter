# Butter
Merges 2 Go structs

### Merge 2 structs & return a map
```go
user1 := user{
Name: "joe",
}
user2 := user{
Msg: "hello!",
}
result := butter.MergeStructsToMap(user1, user2)

// returns
// Name -> Joe
// Description -> hello!
```