# Q: Is a pointer or value reciever faster?

This is a small test to check if a pointer or value type reciever performs
better with a simple struct of two strings. 

The struct principal1 and principal2 are identical. Only difference is that 1
is a value receiver and 2 is a pointer reciever.

```
go test -bench=.
goos: linux
goarch: amd64
pkg: nfsmith.ca/receiverbench
cpu: AMD Ryzen 7 PRO 5850U with Radeon Graphics     
BenchmarkPrincipal1Name-16     	1000000000	         0.2289 ns/op	       0 B/op	       0 allocs/op
BenchmarkPrincipal1Embed-16    	 1296943	       904.3 ns/op	     320 B/op	       5 allocs/op
BenchmarkPrincipal2Name-16     	1000000000	         0.2454 ns/op	       0 B/op	       0 allocs/op
BenchmarkPrincipal2Embed-16    	 1336635	       865.1 ns/op	     320 B/op	       5 allocs/op
PASS
ok  	nfsmith.ca/receiverbench	4.710s
```

On my laptop I find the value receiver is marginally faster.
