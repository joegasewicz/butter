# Butter
Merge Go structs

### Install
```
go get -u github.com/joegasewicz/butter
```

#### Merge 2 structs & return a map
```go
type user struct {
    Name string
    Msg  string
}

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