# Race condition in a caching scenario

The provided code caches *key-value* pairs from a mock database into main memory 
(to reduce access times). 
The code is flawed and contains a **race condition**. 
Modify the code in `main.go` to make this program *thread safe*.

**Please Note**: `golang`'s `map` data structures are not entirely thread safe. 
Multiple readers are fine, but multiple writers are not.

## Test your solution

Use the following command to test for race conditions and correct functionality:
```
go test -race
```

Correct solution:

No output implies a correct solution!
```
$ go test -race
$
```

Incorrect solution:
```
==================
WARNING: DATA RACE
Write by goroutine 7:
...
==================
Found 3 data race(s)
```
