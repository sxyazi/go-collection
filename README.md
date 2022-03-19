# go-collection

English | [简体中文](README_zh-CN.md)

`go-collection` provides developers with a convenient set of functions for working with common slices, maps, and arrays data. These functions are based on the generic types of Go 1.18, which makes it easier to use them without annoying type assertions. In addition to using these functions directly, it also supports method chaining.

```go
collect.Reduce(collect.Filter(collect.Map([]int{1, 2, 3}, fn), fn), fn)
```

Equivalent to:

```go
collect.UseSlice([]int{1, 2, 3}).Map(fn).Filter(fn).Reduce(fn)
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

- `Len` gets the length of the slice

  <details>
  <summary>Examples</summary>

  ```go
  d1 := []int{1, 2, 3}
  collect.Len(d1) // 3

  d2 := []string{"a", "b", "c"}
  collect.Len(d2) // 3
  ```

  </details>

- `Each` iterates over each element in the slice

  <details>
  <summary>Examples</summary>

  ```go
  d := []float64{1, 2, 3}
  collect.Each(d, func(value float64, index int) {
    fmt.Println(index, value)
  })
  ```

  </details>

- `Empty` checks if the slice is empty

  <details>
  <summary>Examples</summary>

  ```go
  var d []int
  collect.Empty(d) // true
  ```

  </details>

- `Same` checks if the contents of two slices are the same

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

- `First` gets the first element of the slice

  <details>
  <summary>Examples</summary>

  ```go
  d1 := []int{1, 2, 3}
  value, ok := collect.First(d1) // 1, true

  var d2 []int
  value, ok = collect.First(d2) // 0, false
  ```

  </details>

- `Last` gets the last element of the slice

  <details>
  <summary>Examples</summary>

  ```go
  d1 := []int{1, 2, 3}
  value, ok := collect.Last(d1) // 3, true

  var d2 []int
  value, ok = collect.Last(d2) // 0, false
  ```

  </details>

- `Index` gets the index of the specified element in the slice, and returns -1 if it does not exist.

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

- `Contains` checks if the slice contains the specified element

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

- `Diff` computes the difference set of two slices

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2, 3}
  collect.Diff(d, []int{2, 3})  // []int{1}
  ```

  </details>

- `Filter` filters the elements in the slice

  <details>
  <summary>Examples</summary>

  ```go
  collect.Filter([]int{1, 2, 3, 4, 5}, func(value, index int) bool {
    return value % 2 == 0
  })  // []int{2, 4}
  ```

  </details>

- `Map` iterates over and sets the value of the elements in the slice

  <details>
  <summary>Examples</summary>

  ```go
  collect.Map([]int{1, 2, 3}, func(value, index int) int {
    return value * 2
  })  // []int{2, 4, 6}
  ```

  </details>

- `Unique` removes duplicate elements in the slice

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2, 3, 3, 4}
  collect.Unique(d)  // []int{1, 2, 3, 4}
  ```

  </details>

- `Duplicates` gets the duplicate elements in the slice

  <details>
  <summary>Examples</summary>

  ```go
  d := []string{"a", "b", "a", "c"}
  collect.Duplicates(d)  // map[int]string{2: "a"}
  ```

  </details>

- `Merge` merges the current slice with other slices

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

- `Random` gets an element of the slice at random

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2}
  value, ok := collect.Random(d)  // 1 or 2, true

  d := []int{}
  value, ok := collect.Random(d)  // 0, false
  ```

  </details>

- `Reverse` reverses the elements in a slice

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2}
  collect.Reverse(d)  // []int{2, 1}
  ```

  </details>

- `Shuffle` randomly shuffles the elements in a slice

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2}
  collect.Shuffle(d)  // []int{1, 2} or []int{2, 1}
  ```

  </details>

- `Slice` takes a segment from a slice

  <details>
  <summary>Examples</summary>

  Function signature: `Slice(items T, offset int)`

  ```go
  d := []int{1, 2, 3, 4, 5}
  collect.Slice(d, 2)   // []int{3, 4, 5}
  collect.Slice(d, -1)  // []int{5}
  collect.Slice(d, -2)  // []int{4, 5}
  ```

  Function signature: `Slice(items T, offset, length int)`

  ```go
  d := []int{1, 2, 3, 4, 5}
  collect.Slice(d, 0, 2)   // []int{1, 2}
  collect.Slice(d, 2, 3)   // []int{3, 4, 5}
  collect.Slice(d, 3, -2)  // []int{3, 4}
  collect.Slice(d, -4, 3)  // []int{2, 3, 4}
  ```

  </details>

- `Split` splits a slice into multiple slices by the specified amount

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2, 3, 4, 5}
  collect.Split(d, 2)  // [][]int{{1, 2}, {3, 4}, {5}}
  ```

  </details>

- `Splice` removes a segment from the slice

  <details>
  <summary>Examples</summary>

  Function signature: `Splice(items T, offset int)`

  ```go
  d := []int{1, 2, 3, 4, 5}
  collect.Splice(&d, 2)  // []int{3, 4, 5}
  d                      // []int{1, 2}
  ```

  Function signature: `Splice(items T, offset, length int)`

  ```go
  d := []int{1, 2, 3, 4, 5}
  collect.Splice(&d, 2, 2)  // []int{3, 4}
  d                         // []int{1, 2, 5}
  ```

  Function signature: `Splice(items T, offset, length int, replacements ...T|E)`

  ```go
  d1 := []int{1, 2, 3, 4}
  collect.Splice(&d1, 1, 2, []int{22, 33})  // []int{2, 3}
  d1                                        // []int{1, 22, 33, 4}

  d2 := []int{1, 2, 3, 4}
  collect.Splice(&d2, 1, 2, 22, 33)  // []int{2, 3}
  d2                                 // []int{1, 22, 33, 4}

  d3 := []int{1, 2, 3, 4}
  collect.Splice(&d3, 1, 2, []int{22}, 33, []int{55})  // []int{2, 3}
  d3                                                   // []int{1, 22, 33, 55, 4}
  ```

  It is worth noting that this method also supports the use of negative numbers as arguments, and its behavior is the same as that of `Slice`, which is not repeated here due to space constraints.

  </details>

- `Reduce` reduces the collection to a single value, and the parameters of each iteration are the results of the previous iteration

  <details>
  <summary>Examples</summary>

  ```go
  collect.Reduce([]int{1, 2, 3}, 100, func(carry, value, key int) int {
  	return carry + value
  })  // 106
  ```

  </details>

- `Pop` removes and returns the last element of the collection

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2}
  v, ok := collect.Pop(&d)  // 2, true
  d                         // []int{1}

  c := collect.UseSlice([]int{1, 2})
  v, ok := c.Pop()  // 2, true
  c.All()           // []int{1}
  ```

  </details>

- `Push` appends an element to the end of a collection

  <details>
  <summary>Examples</summary>

  ```go
  d := []int{1, 2}
  Push(&d, 3)
  d  // []int{1, 2, 3}

  collect.UseSlice([]int{1, 2}).Push(3).All()  // []int{1, 2, 3}
  ```

  </details>

- `Where` filters the collection by the specified rules

  <details>
  <summary>Examples</summary>

  Function signature: `Where(items T, target any)`

  ```go
  collect.Where([]int{1, 2, 3}, 2)  // []int{2}
  ```

  Function signature: `Where(items T, operator string, target any)`

  ```go
  d := []int{1, 2, 3, 4}
  collect.Where(d, "=", 2)   // []int{2}
  collect.Where(d, "!=", 2)  // []int{1, 3, 4}
  collect.Where(d, ">", 2)   // []int{3, 4}
  collect.Where(d, ">=", 2)  // []int{2, 3, 4}
  collect.Where(d, "<", 3)   // []int{1, 2}
  collect.Where(d, "<=", 3)  // []int{1, 2, 3}
  ```

  Function signature: `Where(items T, key any, target any)`

  ```go
  d := []User{{ID: 1, Name: "Hugo"}, {ID: 2, Name: "Lisa"}, {ID: 3, Name: "Iris"}, {ID: 4, Name: "Lisa"}}
  collect.Where(d, "Name", "Lisa")  // []User{{2 Lisa} {4 Lisa}}
  ```

  Function signature: `Where(items T, key any, operator string, target any)`

  ```go
  d := []User{{ID: 1, Name: "Hugo"}, {ID: 2, Name: "Lisa"}, {ID: 3, Name: "Iris"}, {ID: 4, Name: "Lisa"}}
  collect.Where(d, "Name", "!=", "Lisa")  // []User{{1 Hugo} {3 Iris}}
  ```

  </details>

- `WhereIn` removes elements from the collection that do not exist in the specified slice

  <details>
  <summary>Examples</summary>

  Function signature: `WhereIn(items T, targets []any)`

  ```go
  d := []int{1, 2, 3, 4}
  collect.WhereIn(d, []int{2, 3})  // []int{2, 3}
  ```

  Function signature: `WhereIn(items T, key any, targets []any)`

  ```go
  d := []User{{ID: 1, Name: "Hugo"}, {ID: 2, Name: "Lisa"}, {ID: 3, Name: "Iris"}, {ID: 4, Name: "Lisa"}}
  collect.WhereIn(d, "Name", []string{"Hugo", "Iris"})  // []User{{1 Hugo} {3 Iris}}
  ```

  </details>

- `WhereNotIn` removes elements from the collection that exist in the specified slice

  <details>
  <summary>Examples</summary>

  Function signature: `WhereNotIn(items T, targets []any)`

  ```go
  d := []int{1, 2, 3, 4}
  collect.WhereNotIn(d, []int{2, 3})  // []int{1, 4}
  ```

  Function signature: `WhereNotIn(items T, key any, targets []any)`

  ```go
  d := []User{{ID: 1, Name: "Hugo"}, {ID: 2, Name: "Lisa"}, {ID: 3, Name: "Iris"}, {ID: 4, Name: "Lisa"}}
  collect.WhereNotIn(d, "Name", []string{"Lisa", "Iris"})  // []User{{1 Hugo}}
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

### Map

The corresponding chained function is `collect.UseMap()`

- `Len` gets the number of elements in the map

  <details>
  <summary>Examples</summary>

  ```go
  d1 := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Len(d1) // 3
  ```

  </details>

- `Empty` checks if the map is empty

  <details>
  <summary>Examples</summary>

  ```go
  var d map[string]int
  collect.Empty(d) // true
  ```

  </details>

- `Only` gets the elements of the map with the specified keys

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Only(d, "a")       // map[string]int{"a": 1}
  collect.Only(d, "a", "b")  // map[string]int{"a": 1, "b": 2}
  ```

  </details>

- `Except` gets the elements of the map with the specified keys removed

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Except(d, "a")       // map[string]int{"b": 2, "c": 3}
  collect.Except(d, "a", "b")  // map[string]int{"c": 3}
  ```

  </details>

- `Keys` gets all the keys in the map

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Keys(d)  // []string{"a", "b", "c"}
  ```

  </details>

- `DiffKeys` compares with the given collection and returns the key/value pairs in the given collection that do not exist in the original collection

  <details>
  <summary>Examples</summary>

  ```go
  d1 := map[string]int{"a": 1, "b": 2, "c": 3}
  d2 := map[string]int{"b": 22, "c": 33}

  collect.DiffKeys(d1, d2)  // map[string]int{"a": 1}
  ```

  </details>

- `Has` checks if the map contains the specified key

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1}
  collect.Has(d, "a")  // true
  ```

  </details>

- `Get` gets the value of the specified key in the map

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1}

  value, ok := collect.Get(d, "a")  // 1, true
  value, ok := collect.Get(d, "b")  // 0, false
  ```

  </details>

- `Put` sets the value of the specified key in the map

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1}
  collect.Put(d, "b", 2)  // map[string]int{"a": 1, "b": 2}
  ```

  </details>

- `Pull` removes the specified key from the collection and returns its value

  <details>
  <summary>Examples</summary>

  ```go
  d := map[string]int{"a": 1, "b": 2}
  v, ok := collect.Pull(d, "b")  // 2, true
  d                              // map[string]int{"a": 1}
  ```

  </details>

- `Merge` merges the current map with other maps

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

- `Union` unites the current map with other maps, and the items in the original map are given priority

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

- `Sum` calculates the sum

  <details>
  <summary>Examples</summary>

  ```go
  collect.Sum([]float64{1, 3.14})  // 4.14
  ```

  </details>

- `Min` calculates the minimum value

  <details>
  <summary>Examples</summary>

  ```go
  collect.Min([]int{0, 1, -3})  // -3
  ```

  </details>

- `Max` calculates the maximum value

  <details>
  <summary>Examples</summary>

  ```go
  collect.Max([]int{0, 1, -3})  // 1
  ```

  </details>

- `Sort` sorts the numbers in the collection in ascending order

  <details>
  <summary>Examples</summary>

  ```go
  collect.Sort([]float64{1, -4, 0, -4.3})  // []float64{-4.3, -4, 0, 1}
  ```

  </details>

- `SortDesc` sorts the numbers in the collection in descending order

  <details>
  <summary>Examples</summary>

  ```go
  collect.SortDesc([]float64{1, -4, 0, -4.3})  // []float64{1, 0, -4, -4.3}
  ```

  </details>

- `Avg` calculates the average

  <details>
  <summary>Examples</summary>

  ```go
  collect.Avg([]int{1, 2, 3, 4})  // 2.5
  ```

  </details>

- `Median` calculates the median

  <details>
  <summary>Examples</summary>

  ```go
  collect.Median([]int{1, 2, 3, 4})  // 2.5
  ```

  </details>

### Standalone functions

Due to Golang's support for generics, it is [not possible to define generic types in methods](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#no-parameterized-methods), so only their function implementations (which do not support chain calls) are listed below:

- `AnyGet` gets value of arbitrary types (slices, maps, arrays, structures, and pointers to these) in a non-strict form

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

- `Pluck` retrieves all values for a given key. supports all values supported by `AnyGet`

  <details>
  <summary>Examples</summary>

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}}
  collect.Pluck[int](d, "ID")  // int[]{33, 193}
  ```

  </details>

- `MapPluck` retrieves all values of a given key, only maps are supported

  <details>
  <summary>Examples</summary>

  ```go
  d := []map[string]int{{"ID": 33, "Score": 10}, {"ID": 193, "Score": 6}}
  collect.MapPluck(d, "ID")  // int[]{33, 193}
  ```

  </details>

- `KeyBy` retrieves a collection with the value of the given key as the identifier (if there are duplicate keys, only the last one will be kept). Supports all values supported by `AnyGet`

  <details>
  <summary>Examples</summary>

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Peter"}}
  collect.KeyBy[string](d, "Name")  // map[Lucy:{33 Lucy} Peter:{194 Peter}]
  ```

  </details>

- `MapKeyBy` retrieves the collection with the value of the given key as the identifier (if there are duplicate keys, only the last one will be kept), only maps are supported

  <details>
  <summary>Examples</summary>

  ```go
  d := []map[string]int{{"ID": 33, "Score": 6}, {"ID": 193, "Score": 10}, {"ID": 194, "Score": 10}}
  collect.MapKeyBy(d, "Score")  // map[6:map[ID:33 Score:6] 10:map[ID:194 Score:10]]
  ```

  </details>

- `GroupBy` groups the items in a collection using the value of the given key as the identifier. Supports all values supported by `AnyGet`

  <details>
  <summary>Examples</summary>

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Peter"}}
  collect.GroupBy[string](d, "Name")  // map[Lucy:[{33 Lucy}] Peter:[{193 Peter} {194 Peter}]]
  ```

  </details>

- `MapGroupBy` groups items in a collection using the value of the given key as the identifier, only maps are supported

  <details>
  <summary>Examples</summary>

  ```go
  d := []map[string]int{{"ID": 33, "Score": 6}, {"ID": 193, "Score": 10}, {"ID": 194, "Score": 10}}
  collect.MapGroupBy(d, "Score")  // map[6:[map[ID:33 Score:6]] 10:[map[ID:193 Score:10] map[ID:194 Score:10]]]
  ```

  </details>

- `Count` counts the number of occurrences of each element in the slice

  <details>
  <summary>Examples</summary>

  ```go
  d := []bool{true, true, false}
  collect.Count(d)  // map[bool]int{true: 2, false: 1}
  ```

  </details>

- `Times` creates a new collection of slice by calling the callback with specified number of times

  <details>
  <summary>Examples</summary>

  ```go
  collect.Times(3, func(number int) float64 {
  	return float64(number) * 3.14
  })  // *SliceCollection{[]float64{3.14, 6.28, 9.42}}
  ```

  </details>

- `SortBy` calls a callback for each element and performs an ascending sort by the return value of the callback

  <details>
  <summary>Examples</summary>

  ```go
  collect.SortBy([]int{2, 1, 3}, func(item, index int) string {
  	return strconv.Itoa(item)
  })  // *SliceCollection{[]int{1, 2, 3}}
  ```

  </details>

- `SortByDesc` calls a callback for each element and performs a descending sort by the return value of the callback

  <details>
  <summary>Examples</summary>

  ```go
  collect.SortByDesc([]int{2, 1, 3}, func(item, index int) string {
  	return strconv.Itoa(item)
  })  // *SliceCollection{[]int{3, 2, 1}}
  ```

  </details>

## License

go-collection is [MIT licensed](LICENSE).
