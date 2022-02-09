# go-collection

English | [简体中文](README_zh-CN.md)

`go-collection` provides developers with a convenient set of functions for working with common slices, dicts, and arrays data. These functions are based on the generic types of Go 1.18, which makes it easier to use them without annoying type assertions. In addition to using these functions directly, it also supports method chaining.

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

  <details>
  <summary>Examples</summary>

  ```go
  d1 := []int{1, 2, 3}
  collect.Len(d1) // 3

  d2 := []string{"a", "b", "c"}
  collect.Len(d2) // 3
  ```

  </details>

- Each: Iterates over each element in the slice

  <details>
  <summary>Examples</summary>

  ```go
  d := []float64{1, 2, 3}
  collect.Each(d, func(value float64, index int) {
    fmt.Println(index, value)
  })
  ```

  </details>

- Empty: Checks if the slice is empty

  <details>
  <summary>Examples</summary>

  ```go
  var d []int
  collect.Empty(d) // true
  ```

  </details>

- Same: Checks if the contents of two slices are the same

  <details>
  <summary>Examples</summary>

  ```go
  d1 := []int{1, 2, 3}
  d2 := []int{1, 2, 3}
  collect.Same(d1, d2) // true

  d3 := [][]int{{1, 2, 3}, {4, 5, 6}}
  d4 := [][]int{{1, 2, 3}, {4, 5, 6}}
  collect.Same(d3, d4) // true
  ```

  </details>

- First: Gets the first element of the slice

  <details>
  <summary>Examples</summary>

  ```go
  d1 := []int{1, 2, 3}
  value, ok := collect.First(d1) // 1, true

  var d2 []int
  value, ok = collect.First(d2) // 0, false
  ```

  </details>

- Last: Gets the last element of the slice

  <details>
  <summary>Examples</summary>

  ```go
  d1 := []int{1, 2, 3}
  value, ok := collect.Last(d1) // 3, true

  var d2 []int
  value, ok = collect.Last(d2) // 0, false
  ```

  </details>

- Index: Gets the index of the specified element in the slice, and returns -1 if it does not exist.

  <details>
  <summary>Examples</summary>

  ```go
  d1 := []int{1, 2, 3}
  collect.Index(d1, 2) // 1

  s1 := []string{"a", "b", "c"}
  s2 := []string{"d", "e", "f"}
  collect.Index([][]string{s1, s2}, s2) // 1
  ```

  </details>

- Contains: Checks if the slice contains the specified element

  <details>
  <summary>Examples</summary>

  ```go
  d1 := []int{1, 2, 3}
  collect.Contains(d1, 2) // true

  s1 := []string{"a", "b", "c"}
  s2 := []string{"d", "e", "f"}
  collect.Contains([][]string{s1, s2}, s2) // true
  ```

  </details>

- Diff: Computes the difference set of two slices

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2, 3}
  collect.Diff(d, []int{2, 3})  // []int{1}
  ```

  </details>

- Filter: Filters the elements in the slice

  <details>
  <summary>Examples</summary>

  ```go
  collect.Filter([]int{1, 2, 3, 4, 5}, func(value, index int) bool {
    return value % 2 == 0
  })  // []int{2, 4}
  ```

  </details>

- Map: Iterates over and sets the value of the elements in the slice

  <details>
  <summary>Examples</summary>

  ```go
  collect.Map([]int{1, 2, 3}, func(value, index int) int {
    return value * 2
  })  // []int{2, 4, 6}
  ```

  </details>

- Unique: Removes duplicate elements from slices

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2, 3, 3, 4}
  collect.Unique(d)  // []int{1, 2, 3, 4}
  ```

  </details>

- Merge: Merges the current slice with other slices

  <details>
  <summary>Examples</summary>

  ```go
  d1 := []int{1, 2}
  d2 := []int{3, 4}
  d3 := []int{5, 6}

  collect.Merge(d1, d2)      // []int{1, 2, 3, 4}
  collect.Merge(d1, d2, d3)  // []int{1, 2, 3, 4, 5, 6}
  ```

  </details>

- Random: Gets an element of the slice at random

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2}
  value, ok := collect.Random(d)  // 1 or 2, true

  d := []int{}
  value, ok := collect.Random(d)  // 0, false
  ```

  </details>

- Reverse: Reverses the elements in a slice

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2}
  collect.Reverse(d)  // []int{2, 1}
  ```

  </details>

- Shuffle: Randomly shuffles the elements in a slice

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2}
  collect.Shuffle(d)  // []int{1, 2} or []int{2, 1}
  ```

  </details>

- Slice: Takes a segment from a slice

  <details>
  <summary>Examples</summary>

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

  </details>

- Split: Splits a slice into multiple slices by the specified amount

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2, 3, 4, 5}
  collect.Split(d, 2)  // [][]int{{1, 2}, {3, 4}, {5}}
  ```

  </details>

- Splice: Removes a segment from the slice

  <details>
  <summary>Examples</summary>

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

  </details>

- Count: Counts the number of occurrences of each element in the slice

  <details>
  <summary>Examples</summary>

  ```go
  d := []bool{true, true, false}
  collect.Count(d)  // map[bool]int{true: 2, false: 1}
  ```

  </details>

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

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Only(d, "a")       // map[string]int{"a": 1}
  collect.Only(d, "a", "b")  // map[string]int{"a": 1, "b": 2}
  ```

  </details>

- Except: Gets the elements of the dict with the specified keys removed

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Except(d, "a")       // map[string]int{"b": 2, "c": 3}
  collect.Except(d, "a", "b")  // map[string]int{"c": 3}
  ```

  </details>

- Keys: Gets all the keys in the dict

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Keys(d)  // []string{"a", "b", "c"}
  ```

  </details>

- DiffKeys: Compares with the given collection and returns the key/value pairs in the given collection that do not exist in the original collection

  <details>
  <summary>Examples</summary>

  ```go
  d1 := map[string]int{"a": 1, "b": 2, "c": 3}
  d2 := map[string]int{"b": 22, "c": 33}

  collect.DiffKeys(d1, d2)  // map[string]int{"a": 1}
  ```

  </details>

- Has: Checks if the dict contains the specified key

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1}
  collect.Has(d, "a")  // true
  ```

  </details>

- Set: Sets the value of the specified key in the dict

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1}
  collect.Set(d, "b", 2)  // map[string]int{"a": 1, "b": 2}
  ```

  </details>

- Get: Gets the value of the specified key in the dict

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1}

  value, ok := collect.Get(d, "a")  // 1, true
  value, ok := collect.Get(d, "b")  // 0, false
  ```

  </details>

- Merge: Merges the current dict with other dicts

  <details>
  <summary>Examples</summary>

  ```go
  d1 := map[string]int{"a": 1, "b": 2}
  d2 := map[string]int{"b": 22}
  d3 := map[string]int{"b": 222, "c": 3}

  collect.MapMerge(d1, d2)            // map[string]int{"a": 1, "b": 22}
  collect.UseMap(d1).Merge(d2).All()  // Equal to the above

  collect.MapMerge(d1, d2, d3)            // map[string]int{"a": 1, "b": 222, "c": 3}
  collect.UseMap(d1).Merge(d2, d3).All()  // Equal to the above
  ```

  </details>

- Union: Unites the current dict with other dicts, and the items in the original dict are given priority

  <details>
  <summary>Examples</summary>

  ```go
  d1 := map[string]int{"a": 1, "b": 2}
  d2 := map[string]int{"b": 22, "c": 3}
  collect.Union(d1, d2)  // map[string]int{"a": 1, "b": 2, "c": 3}
  ```

  </details>

### Number slice

The corresponding chained function is `collect.UseNumber()`，which is a subset of [slice](#Slice) and includes, in addition to all the methods of slice, the additional:

- Sum: Calculates the sum

  <details>
  <summary>Examples</summary>

  ```go
  collect.Sum([]float64{1.1, 2.2, 3.3})  // 6.6
  ```

  </details>

- Avg: Calculates the average

  <details>
  <summary>Examples</summary>

  ```go
  collect.Avg([]float64{1.1, 2.2, 3.3})  // 2.2
  ```

  </details>

- Min: Calculates the minimum value

  <details>
  <summary>Examples</summary>

  ```go
  collect.Min([]int{0, 1, -3})  // -3
  ```

  </details>

- Max: Calculates the maximum value

  <details>
  <summary>Examples</summary>

  ```go
  collect.Max([]int{0, 1, -3})  // 1
  ```

  </details>

### Standalone functions

Due to Golang's support for generics, it is [not possible to define generic types in methods](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#no-parameterized-methods), so only their function implementations (which do not support chain calls) are listed below:

- AnyGet: Gets value of arbitrary types (slices, dicts, arrays, structures, and pointers to these) in a non-strict form

  <details>
  <summary>Examples</summary>

  ```go
  m := map[string]int{"a": 1, "b": 2}
  collect.AnyGet[int](m, "b")  // 2

  u := &User{"Email": "user@example.com"}
  collect.AnyGet[string](u, "Email")  // user@example.com

  s := [][]int{{1, 2}, {3, 4}}
  collect.AnyGet[[]int](s, 1)  // []{3, 4}
  ```

  </details>

- Pluck: Retrieves all values for a given key and supports all values supported by `AnyGet`

  <details>
  <summary>Examples</summary>

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}}
  collect.Pluck[int](d, "ID")  // int[]{33, 193}
  ```

  </details>

- MapPluck: Retrieves all values of a given key, only dicts are supported

  <details>
  <summary>Examples</summary>

  ```go
  d := []map[string]int{{"ID": 33, "Score": 10}, {"ID": 193, "Score": 6}}
  collect.MapPluck(d, "ID")  // int[]{33, 193}
  ```

  </details>

- KeyBy: Retrieves a collection with the value of the given key as the identifier (if there are duplicate keys, only the last one will be kept). Supports all values supported by `AnyGet`

  <details>
  <summary>Examples</summary>

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Peter"}}
  collect.KeyBy[string](d, "Name")  // map[Lucy:{33 Lucy} Peter:{194 Peter}]
  ```

  </details>

- MapKeyBy: Retrieves the set with the value of the given key as the identifier (if there are duplicate keys, only the last one will be kept), only dicts are supported

  <details>
  <summary>Examples</summary>

  ```go
  d := []map[string]int{{"ID": 33, "Score": 6}, {"ID": 193, "Score": 10}, {"ID": 194, "Score": 10}}
  collect.MapKeyBy(d, "Score")  // map[6:map[ID:33 Score:6] 10:map[ID:194 Score:10]]
  ```

  </details>

- GroupBy: Groups the items in a collection using the value of the given key as the identifier. Supports all values supported by `AnyGet`

  <details>
  <summary>Examples</summary>

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Peter"}}
  collect.GroupBy[string](d, "Name")  // map[Lucy:[{33 Lucy}] Peter:[{193 Peter} {194 Peter}]]
  ```

  </details>

- MapGroupBy: Groups items in a collection using the value of the given key as the identifier, only dicts are supported

  <details>
  <summary>Examples</summary>

  ```go
  d := []map[string]int{{"ID": 33, "Score": 6}, {"ID": 193, "Score": 10}, {"ID": 194, "Score": 10}}
  collect.MapGroupBy(d, "Score")  // map[6:[map[ID:33 Score:6]] 10:[map[ID:193 Score:10] map[ID:194 Score:10]]]
  ```

  </details>

## License

go-collection is [MIT licensed](LICENSE).
