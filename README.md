# Butter
Merge Go structs

### Install
```
go get -u github.com/joegasewicz/butter
```

#### Merge 2 structs & return a map
Create a map of the combined results of two structs. The returned map
will merge the incoming struct on top of the original struct. 
The returned map is useful as a value to GORM's `Updates` method for example.
```go
type user struct {
    ID   uint
    Name string
    Msg  string
}

incoming := user{
    Name: "joe",
}
original := user{
    ID:  1,
    Msg: "hello!",
}
merged := butter.MergeStructsToMap(original, incoming)

// returns
// Name -> Joe
// Description -> hello!
// Not we could pass the result to GORM's `Updates` method
db.Model(&user).Updates(merged)
```