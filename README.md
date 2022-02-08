# go-collection

English | [简体中文](README_zh-CN.md)

`go-collection` provides developers with a convenient set of functions for working with common slices, dicts, and arrays of data. These functions are based on the generic types of Go 1.18, which makes it easier to use them without annoying type assertions. In addition to using these functions directly, it also supports method chaining.

```go
collect.Reduce(collect.Filter(collect.Map([]int{1, 2, 3}, fn), fn), fn)
```

Equivalent to:

```go
collect.UseSlice([]int{1, 2, 3}).Map(fn).Filter(fn).Reduce(fn).All()
```

## Installation

```shell
go get -u github.com/sxyazi/go-collection
```

Then import it

```go
import collect "github.com/sxyazi/go-collection"
```

## API

Its API is very simple and if you have used other similar packages, you should be able to get started with it in a few minutes. **For convenience, they are described below in function form**.

### Slice

The corresponding chained function is `collect.UseSlice()`

- Len: Gets the length of the slice

  ```go
  d1 := []int{1, 2, 3}
  collect.Len(d1) // 3

  d2 := []string{"a", "b", "c"}
  collect.Len(d2) // 3
  ```

- Each: Iterates over each element in the slice

  ```go
  d := []float64{1, 2, 3}
  collect.Each(d, func(value float64, index int) {
    fmt.Println(index, value)
  })
  ```

- Empty: Checks if the slice is empty

  ```go
  var d []int
  collect.Empty(d) // true
  ```

- Same: Checks if the contents of two slices are the same

  ```go
  d1 := []int{1, 2, 3}
  d2 := []int{1, 2, 3}
  collect.Same(d1, d2) // true

  d3 := [][]int{{1, 2, 3}, {4, 5, 6}}
  d4 := [][]int{{1, 2, 3}, {4, 5, 6}}
  collect.Same(d3, d4) // true
  ```

- First: Gets the first element of the slice

  ```go
  d1 := []int{1, 2, 3}
  value, ok := collect.First(d1) // 1, true

  var d2 []int
  value, ok = collect.First(d2) // 0, false
  ```

- Last: Gets the last element of the slice

  ```go
  d1 := []int{1, 2, 3}
  value, ok := collect.Last(d1) // 3, true

  var d2 []int
  value, ok = collect.Last(d2) // 0, false
  ```

- Index: Gets the index of the specified element in the slice, and returns -1 if it does not exist.

  ```go
  d1 := []int{1, 2, 3}
  collect.Index(d1, 2) // 1

  s1 := []string{"a", "b", "c"}
  s2 := []string{"d", "e", "f"}
  collect.Index([][]string{s1, s2}, s2) // 1
  ```

- Contains: Checks if the slice contains the specified element

  ```go
  d1 := []int{1, 2, 3}
  collect.Contains(d1, 2) // true

  s1 := []string{"a", "b", "c"}
  s2 := []string{"d", "e", "f"}
  collect.Contains([][]string{s1, s2}, s2) // true
  ```

- Diff: Computes the difference set of two slices

  ```go
  d := []int{1, 2, 3}
  collect.Diff(d, []int{2, 3})  // []int{1}
  ```

- Filter: Filters the elements in the slice

  ```go
  collect.Filter([]int{1, 2, 3, 4, 5}, func(value, index int) bool {
    return value % 2 == 0
  })  // []int{2, 4}
  ```

- Map: Iterates over and sets the value of the elements in the slice

  ```go
  collect.Map([]int{1, 2, 3}, func(value, index int) int {
    return value * 2
  })  // []int{2, 4, 6}
  ```

- Unique: Removes duplicate elements from slices

  ```go
  d := []int{1, 2, 3, 3, 4}
  collect.Unique(d)  // []int{1, 2, 3, 4}
  ```

- Merge: Merges the current slice with other slices

  ```go
  d1 := []int{1, 2}
  d2 := []int{3, 4}
  d3 := []int{5, 6}

  collect.Merge(d1, d2)      // []int{1, 2, 3, 4}
  collect.Merge(d1, d2, d3)  // []int{1, 2, 3, 4, 5, 6}
  ```

- Random: Gets an element of the slice at random

  ```go
  d := []int{1, 2}
  value, ok := collect.Random(d)  // 1 or 2, true

  d := []int{}
  value, ok := collect.Random(d)  // 0, false
  ```

- Reverse: Reverses the elements in a slice

  ```go
  d := []int{1, 2}
  collect.Reverse(d)  // []int{2, 1}
  ```

- Shuffle: Randomly shuffles the elements in a slice

  ```go
  d := []int{1, 2}
  collect.Shuffle(d)  // []int{1, 2} or []int{2, 1}
  ```

- Slice: Takes a segment from a slice

  Function signature: `Slice(items T, offset int)`

  ```go
  d := []int{1, 2, 3, 4}
  collect.Slice(d, 2)  // []int{3, 4}
  ```

  Function signature: `Slice(items T, offset, length int)`

  ```go
  collect.Slice(d, 0, 2)  // []int{1, 2}
  collect.Slice(d, 2, 2)  // []int{3, 4}
  ```

- Split: Splits a slice into multiple slices by the specified amount

  ```go
  d := []int{1, 2, 3, 4, 5}
  collect.Split(d, 2)  // [][]int{{1, 2}, {3, 4}, {5}}
  ```

- Splice: Removes a segment from the slice

  Function signature: `Splice(items T, offset int)`

  ```go
  d := []int{1, 2, 3, 4, 5}
  collect.Splice(d, 2)  // []int{1, 2}
  ```

  Function signature: `Splice(items T, offset, length int)`

  ```go
  d := []int{1, 2, 3, 4, 5}
  collect.Splice(d, 2, 2)  // []int{1, 2, 5}
  ```

  Function signature: `Slice(items T, offset, length int, replacements ...T|E)`

  ```go
  d := []int{1, 2, 3, 4}
  collect.Splice(d, 1, 2, []int{22, 33})             // []int{1, 22, 33, 4}
  collect.Splice(d, 1, 2, 233, 333)                  // []int{1, 222, 333, 4}
  collect.Splice(d, 1, 2, []int{22}, 33, []int{44})  // []int{1, 22, 33, 44, 4}
  ```

- Count: Counts the number of occurrences of each element in the slice
  ```go
  d := []bool{true, true, false}
  collect.Count(d)  // map[bool]int{true: 2, false: 1}
  ```

### Array

Exactly the same as [slice](#Slice), you just pass in the array converted to a slice:

```go
arr := [3]int{1, 2, 3}

collect.Len(arr[:])
// or
collect.UseSlice(arr[:]).Len()
```

### Dict

The corresponding chained function is `collect.UseMap()`

- Only: Gets the elements of the dict with the specified keys

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Only(d, "a")       // map[string]int{"a": 1}
  collect.Only(d, "a", "b")  // map[string]int{"a": 1, "b": 2}
  ```

- Except: Gets the elements of the dict with the specified keys removed

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Except(d, "a")       // map[string]int{"b": 2, "c": 3}
  collect.Except(d, "a", "b")  // map[string]int{"c": 3}
  ```

- Keys: Gets all the keys in the dict

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Keys(d)  // []string{"a", "b", "c"}
  ```

- DiffKeys: Compares with the given collection and returns the key/value pairs in the given collection that do not exist in the original collection

  ```go
  d1 := map[string]int{"a": 1, "b": 2, "c": 3}
  d2 := map[string]int{"b": 22, "c": 33}

  collect.DiffKeys(d1, d2)  // map[string]int{"a": 1}
  ```

- Has: Checks if the dict contains the specified key

  ```go
  d := map[string]int{"a": 1}
  collect.Has(d, "a")  // true
  ```

- Set: Sets the value of the specified key in the dict

  ```go
  d := map[string]int{"a": 1}
  collect.Set(d, "b", 2)  // map[string]int{"a": 1, "b": 2}
  ```

- Get: Gets the value of the specified key in the dict

  ```go
  d := map[string]int{"a": 1}

  value, ok := collect.Get(d, "a")  // 1, true
  value, ok := collect.Get(d, "b")  // 0, false
  ```

- Merge: Merges the current dict with other dicts

  ```go
  d1 := map[string]int{"a": 1, "b": 2}
  d2 := map[string]int{"b": 22}
  d3 := map[string]int{"b": 222, "c": 3}

  collect.MapMerge(d1, d2)            // map[string]int{"a": 1, "b": 22}
  collect.UseMap(d1).Merge(d2).All()  // Equal to the above

  collect.MapMerge(d1, d2, d3)            // map[string]int{"a": 1, "b": 222, "c": 3}
  collect.UseMap(d1).Merge(d2, d3).All()  // Equal to the above
  ```

- Union: Unites the current dict with other dicts, and the items in the original dict are given priority

  ```go
  d1 := map[string]int{"a": 1, "b": 2}
  d2 := map[string]int{"b": 22, "c": 3}
  collect.Union(d1, d2)  // map[string]int{"a": 1, "b": 2, "c": 3}
  ```

### Number slice

The corresponding chained function is `collect.UseNumber()`，which is a subset of [slice](#Slice) and includes, in addition to all the methods of slice, the additional:

- Sum: Calculates the sum

  ```go
  collect.Sum([]float64{1.1, 2.2, 3.3})  // 6.6
  ```

- Avg: Calculates the average

  ```go
  collect.Avg([]float64{1.1, 2.2, 3.3})  // 2.2
  ```

- Min: Calculates the minimum value

  ```go
  collect.Min([]int{0, 1, -3})  // -3
  ```

- Max: Calculates the maximum value

  ```go
  collect.Max([]int{0, 1, -3})  // 1
  ```

### Standalone functions

Due to Golang's support for generics, it is [not possible to define generic types in methods](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#no-parameterized-methods), so only their function implementations (which do not support chain calls) are listed below:

- AnyGet: Gets value of arbitrary types (slices, dicts, arrays, structures, and pointers to these) in a non-strict form

  ```go
  m := map[string]int{"a": 1, "b": 2}
  collect.AnyGet[int](m, "b")  // 2

  u := &User{"Email": "user@example.com"}
  collect.AnyGet[string](u, "Email")  // user@example.com

  s := [][]int{{1, 2}, {3, 4}}
  collect.AnyGet[[]int](s, 1)  // []{3, 4}
  ```

- Pluck: Retrieves all values for a given key and supports all values supported by `AnyGet`

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}}
  collect.Pluck[int](d, "ID")  // int[]{33, 193}
  ```

- MapPluck: Retrieves all values of a given key, only dicts are supported

  ```go
  d := []map[string]int{{"ID": 33, "Score": 10}, {"ID": 193, "Score": 6}}
  collect.MapPluck(d, "ID")  // int[]{33, 193}
  ```

- KeyBy: Retrieves a collection with the value of the given key as the identifier (if there are duplicate keys, only the last one will be kept). Supports all values supported by `AnyGet`

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Peter"}}
  collect.KeyBy[string](d, "Name")  // map[Lucy:{33 Lucy} Peter:{194 Peter}]
  ```

- MapKeyBy: Retrieves the set with the value of the given key as the identifier (if there are duplicate keys, only the last one will be kept), only dicts are supported

  ```go
  d := []map[string]int{{"ID": 33, "Score": 6}, {"ID": 193, "Score": 10}, {"ID": 194, "Score": 10}}
  collect.MapKeyBy(d, "Score")  // map[6:map[ID:33 Score:6] 10:map[ID:194 Score:10]]
  ```

- GroupBy: Groups the items in a collection using the value of the given key as the identifier. Supports all values supported by `AnyGet`

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Peter"}}
  collect.GroupBy[string](d, "Name")  // map[Lucy:[{33 Lucy}] Peter:[{193 Peter} {194 Peter}]]
  ```

- MapGroupBy: Groups items in a collection using the value of the given key as the identifier, only dicts are supported

  ```go
  d := []map[string]int{{"ID": 33, "Score": 6}, {"ID": 193, "Score": 10}, {"ID": 194, "Score": 10}}
  collect.MapGroupBy(d, "Score")  // map[6:[map[ID:33 Score:6]] 10:[map[ID:193 Score:10] map[ID:194 Score:10]]]
  ```
