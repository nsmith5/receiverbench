# Q: Is a pointer or value reciever faster?

This is a small test to check if a pointer or value type reciever performs
better with a simple struct of two strings. 

The struct principal1 and principal2 are identical. Only difference is that 1
is a value receiver and 2 is a pointer reciever.

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: nfsmith.ca/receiverbench
cpu: AMD Ryzen 7 PRO 5850U with Radeon Graphics     
BenchmarkPrincipal1Name-16     	1000000000	         0.2283 ns/op
BenchmarkPrincipal1Embed-16    	 1367580	       882.9 ns/op
BenchmarkPrincipal2Name-16     	1000000000	         0.2445 ns/op
BenchmarkPrincipal2Embed-16    	 1432716	       934.0 ns/op
PASS
ok  	nfsmith.ca/receiverbench	4.808s
```

On my laptop I find the value receiver is marginally faster.
