# easyjson_test

```
goos: darwin
goarch: arm64
pkg: easyjson_test
BenchmarkEasyJSONMarshal-8       8918133               124.6 ns/op           288 B/op          3 allocs/op
BenchmarkJSONMarshal-8           3004923               400.4 ns/op           240 B/op          3 allocs/op
BenchmarkEasyJSONUnmarshal-8     8044593               149.9 ns/op            20 B/op          2 allocs/op
BenchmarkJSONUnmarshal-8         2195044               530.8 ns/op           176 B/op          4 allocs/op
PASS
ok      easyjson_test   6.464s
```
测试下来，easyjson的速度是原生的json的4倍。