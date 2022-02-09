# go-collection

[English](README.md) | 简体中文

`go-collection` 向开发者提供了一组便利的函数，用于处理常见的切片、字典、数组数据。这些函数基于 Go 1.18 中的泛型实现，这让在使用时更加方便，而无需烦人的类型断言。除了直接使用这些函数外，它同样支持链式调用。

```go
collect.Reduce(collect.Filter(collect.Map([]int{1, 2, 3}, fn), fn), fn)
```

等价于：

```go
collect.UseSlice([]int{1, 2, 3}).Map(fn).Filter(fn).Reduce(fn).All()
```

## 安装

```shell
go get -u github.com/sxyazi/go-collection
```

然后导入它

```go
import collect "github.com/sxyazi/go-collection"
```

## API

它的 API 非常简单，如果你用过其它类似的包，应该可以在几分钟内上手它。**为了方便，下面以函数的形式介绍它们**。

### 切片

对应的链式函数为 `collect.UseSlice()`

- Len：获取切片的长度

  <details>
  <summary>例子</summary>

  ```go
  d1 := []int{1, 2, 3}
  collect.Len(d1) // 3

  d2 := []string{"a", "b", "c"}
  collect.Len(d2) // 3
  ```

  </details>

- Each：遍历切片中的每个元素

  <details>
  <summary>例子</summary>

  ```go
  d := []float64{1, 2, 3}
  collect.Each(d, func(value float64, index int) {
    fmt.Println(index, value)
  })
  ```

  </details>

- Empty：检查切片是否为空

  <details>
  <summary>例子</summary>

  ```go
  var d []int
  collect.Empty(d) // true
  ```

  </details>

- Same：检查两个切片的内容是否相同

  <details>
  <summary>例子</summary>

  ```go
  d1 := []int{1, 2, 3}
  d2 := []int{1, 2, 3}
  collect.Same(d1, d2) // true

  d3 := [][]int{{1, 2, 3}, {4, 5, 6}}
  d4 := [][]int{{1, 2, 3}, {4, 5, 6}}
  collect.Same(d3, d4) // true
  ```

  </details>

- First：获取切片的第一个元素

  <details>
  <summary>例子</summary>

  ```go
  d1 := []int{1, 2, 3}
  value, ok := collect.First(d1) // 1, true

  var d2 []int
  value, ok = collect.First(d2) // 0, false
  ```

  </details>

- Last：获取切片的最后一个元素

  <details>
  <summary>例子</summary>

  ```go
  d1 := []int{1, 2, 3}
  value, ok := collect.Last(d1) // 3, true

  var d2 []int
  value, ok = collect.Last(d2) // 0, false
  ```

  </details>

- Index：获取指定元素在切片中的索引，如果不存在返回 -1

  <details>
  <summary>例子</summary>

  ```go
  d1 := []int{1, 2, 3}
  collect.Index(d1, 2) // 1

  s1 := []string{"a", "b", "c"}
  s2 := []string{"d", "e", "f"}
  collect.Index([][]string{s1, s2}, s2) // 1
  ```

  </details>

- Contains：检查切片中是否包含指定元素

  <details>
  <summary>例子</summary>

  ```go
  d1 := []int{1, 2, 3}
  collect.Contains(d1, 2) // true

  s1 := []string{"a", "b", "c"}
  s2 := []string{"d", "e", "f"}
  collect.Contains([][]string{s1, s2}, s2) // true
  ```

  </details>

- Diff：计算两个切片的差集

  <details>
  <summary>例子</summary>

  ```go
  d := []int{1, 2, 3}
  collect.Diff(d, []int{2, 3})  // []int{1}
  ```

  </details>

- Filter：过滤切片中的元素

  <details>
  <summary>例子</summary>

  ```go
  collect.Filter([]int{1, 2, 3, 4, 5}, func(value, index int) bool {
    return value % 2 == 0
  })  // []int{2, 4}
  ```

  </details>

- Map：遍历并设置切片中元素的值

  <details>
  <summary>例子</summary>

  ```go
  collect.Map([]int{1, 2, 3}, func(value, index int) int {
    return value * 2
  })  // []int{2, 4, 6}
  ```

  </details>

- Unique：去除切片中重复的元素

  <details>
  <summary>例子</summary>

  ```go
  d := []int{1, 2, 3, 3, 4}
  collect.Unique(d)  // []int{1, 2, 3, 4}
  ```

  </details>

- Merge：将当前切片与其它切片合并

  <details>
  <summary>例子</summary>

  ```go
  d1 := []int{1, 2}
  d2 := []int{3, 4}
  d3 := []int{5, 6}

  collect.Merge(d1, d2)      // []int{1, 2, 3, 4}
  collect.Merge(d1, d2, d3)  // []int{1, 2, 3, 4, 5, 6}
  ```

  </details>

- Random：随机获取切片中的一个元素

  <details>
  <summary>例子</summary>

  ```go
  d := []int{1, 2}
  value, ok := collect.Random(d)  // 1 or 2, true

  d := []int{}
  value, ok := collect.Random(d)  // 0, false
  ```

  </details>

- Reverse：反转切片中的元素

  <details>
  <summary>例子</summary>

  ```go
  d := []int{1, 2}
  collect.Reverse(d)  // []int{2, 1}
  ```

  </details>

- Shuffle：随机打乱切片中的元素

  <details>
  <summary>例子</summary>

  ```go
  d := []int{1, 2}
  collect.Shuffle(d)  // []int{1, 2} or []int{2, 1}
  ```

  </details>

- Slice：从切片中截取一段

  <details>
  <summary>例子</summary>

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

- Split：按照指定的数量将切片分割为多个

  <details>
  <summary>例子</summary>

  ```go
  d := []int{1, 2, 3, 4, 5}
  collect.Split(d, 2)  // [][]int{{1, 2}, {3, 4}, {5}}
  ```

  </details>

- Splice：从切片中删除一段

  <details>
  <summary>例子</summary>

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

- Count：统计切片中每个元素出现的次数

  <details>
  <summary>例子</summary>

  ```go
  d := []bool{true, true, false}
  collect.Count(d)  // map[bool]int{true: 2, false: 1}
  ```

  </details>

### 数组

与 [切片](#切片) 完全一致，你只需将数组转换为切片传入：

```go
arr := [3]int{1, 2, 3}

collect.Len(arr[:])
// or
collect.UseSlice(arr[:]).Len()
```

### 字典

对应的链式函数为 `collect.UseMap()`

- Only：获取字典中指定键的元素

  <details>
  <summary>例子</summary>

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Only(d, "a")       // map[string]int{"a": 1}
  collect.Only(d, "a", "b")  // map[string]int{"a": 1, "b": 2}
  ```

  </details>

- Except：获取字典中除去指定键的元素

  <details>
  <summary>例子</summary>

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Except(d, "a")       // map[string]int{"b": 2, "c": 3}
  collect.Except(d, "a", "b")  // map[string]int{"c": 3}
  ```

  </details>

- Keys：获取字典中所有的键

  <details>
  <summary>例子</summary>

  ```go
  d := map[string]int{"a": 1, "b": 2, "c": 3}
  collect.Keys(d)  // []string{"a", "b", "c"}
  ```

  </details>

- DiffKeys：与给定的集合比较，返回给定集合中不存在于原始集合的键/值对

  <details>
  <summary>例子</summary>

  ```go
  d1 := map[string]int{"a": 1, "b": 2, "c": 3}
  d2 := map[string]int{"b": 22, "c": 33}

  collect.DiffKeys(d1, d2)  // map[string]int{"a": 1}
  ```

  </details>

- Has：检查字典中是否包含指定键

  <details>
  <summary>例子</summary>

  ```go
  d := map[string]int{"a": 1}
  collect.Has(d, "a")  // true
  ```

  </details>

- Set：设置字典中指定键的值

  <details>
  <summary>例子</summary>

  ```go
  d := map[string]int{"a": 1}
  collect.Set(d, "b", 2)  // map[string]int{"a": 1, "b": 2}
  ```

  </details>

- Get：获取字典中指定键的值

  <details>
  <summary>例子</summary>

  ```go
  d := map[string]int{"a": 1}

  value, ok := collect.Get(d, "a")  // 1, true
  value, ok := collect.Get(d, "b")  // 0, false
  ```

  </details>

- Merge：将当前字典与其它字典合并

  <details>
  <summary>例子</summary>

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

- Union：将当前字典与其它字典联合，原字典中的项目会被优先考虑

  <details>
  <summary>例子</summary>

  ```go
  d1 := map[string]int{"a": 1, "b": 2}
  d2 := map[string]int{"b": 22, "c": 3}
  collect.Union(d1, d2)  // map[string]int{"a": 1, "b": 2, "c": 3}
  ```

  </details>

### 数字切片

对应的链式函数为 `collect.UseNumber()`，它是 [切片](#切片) 的子集，除切片的所有方法外，还额外包括：

- Sum：求和

  <details>
  <summary>例子</summary>

  ```go
  collect.Sum([]float64{1.1, 2.2, 3.3})  // 6.6
  ```

  </details>

- Avg：求平均数

  <details>
  <summary>例子</summary>

  ```go
  collect.Avg([]float64{1.1, 2.2, 3.3})  // 2.2
  ```

  </details>

- Min：求最小值

  <details>
  <summary>例子</summary>

  ```go
  collect.Min([]int{0, 1, -3})  // -3
  ```

  </details>

- Max：求最大值

  <details>
  <summary>例子</summary>

  ```go
  collect.Max([]int{0, 1, -3})  // 1
  ```

  </details>

### 独立函数

受限于 [Golang 泛型](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#no-parameterized-methods) 的支持，无法在方法中定义泛型类型，因此以下列出的这些只有其函数实现（不支持链式调用）：

- AnyGet：以一种非严格的形式获取任意类型（切片、字典、数组、结构体，以及这些的指针）中的值

  <details>
  <summary>例子</summary>

  ```go
  m := map[string]int{"a": 1, "b": 2}
  collect.AnyGet[int](m, "b")  // 2

  u := &User{"Email": "user@example.com"}
  collect.AnyGet[string](u, "Email")  // user@example.com

  s := [][]int{{1, 2}, {3, 4}}
  collect.AnyGet[[]int](s, 1)  // []{3, 4}
  ```

  </details>

- Pluck：检索给定键的所有值，支持 `AnyGet` 支持的所有值

  <details>
  <summary>例子</summary>

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}}
  collect.Pluck[int](d, "ID")  // int[]{33, 193}
  ```

  </details>

- MapPluck：检索给定键的所有值，只支持字典

  <details>
  <summary>例子</summary>

  ```go
  d := []map[string]int{{"ID": 33, "Score": 10}, {"ID": 193, "Score": 6}}
  collect.MapPluck(d, "ID")  // int[]{33, 193}
  ```

  </details>

- KeyBy：以给定键的值为标识检索集合（若存在重复的键，则只有最后一个会被保留）。支持 `AnyGet` 支持的所有值

  <details>
  <summary>例子</summary>

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Peter"}}
  collect.KeyBy[string](d, "Name")  // map[Lucy:{33 Lucy} Peter:{194 Peter}]
  ```

  </details>

- MapKeyBy：以给定键的值为标识检索集合（若存在重复的键，则只有最后一个会被保留），只支持字典

  <details>
  <summary>例子</summary>

  ```go
  d := []map[string]int{{"ID": 33, "Score": 6}, {"ID": 193, "Score": 10}, {"ID": 194, "Score": 10}}
  collect.MapKeyBy(d, "Score")  // map[6:map[ID:33 Score:6] 10:map[ID:194 Score:10]]
  ```

  </details>

- GroupBy：以给定键的值为标识，对集合中的项目分组。支持 `AnyGet` 支持的所有值

  <details>
  <summary>例子</summary>

  ```go
  d := []User{{ID: 33, Name: "Lucy"}, {ID: 193, Name: "Peter"}, {ID: 194, Name: "Peter"}}
  collect.GroupBy[string](d, "Name")  // map[Lucy:[{33 Lucy}] Peter:[{193 Peter} {194 Peter}]]
  ```

  </details>

- MapGroupBy：以给定键的值为标识，对集合中的项目分组，只支持字典

  <details>
  <summary>例子</summary>

  ```go
  d := []map[string]int{{"ID": 33, "Score": 6}, {"ID": 193, "Score": 10}, {"ID": 194, "Score": 10}}
  collect.MapGroupBy(d, "Score")  // map[6:[map[ID:33 Score:6]] 10:[map[ID:193 Score:10] map[ID:194 Score:10]]]
  ```

  </details>

## 许可

go-collection is [MIT licensed](LICENSE).