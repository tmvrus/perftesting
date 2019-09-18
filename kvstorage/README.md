```
BenchmarkShardedStorage-4          41916             31310 ns/op             819 B/op          3 allocs/op
BenchmarkSingleStorage-4           10000            137800 ns/op             189 B/op          1 allocs/op
BenchmarkSyncMap-4                 55057             24127 ns/op            6691 B/op        604 allocs/op
BenchmarkRefStorage-4             390432              2893 ns/op             275 B/op          0 allocs/op
BenchmarkShardedSyncMap-4          76954             20325 ns/op            6467 B/op        603 allocs/op
```
