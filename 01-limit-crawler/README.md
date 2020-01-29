# Limit your crawler

We provide a web crawler that requests pages excessively (a modified version of the code from the first worksheet - The Go Tour) . 
However, we don't want to burden the web-server too much. 

Your task is to change the code to limit the crawler to at most one page per second, while maintaining concurrency (in other words, `Crawl()` must be called concurrently)

## Hint

This exercise can be solved in three lines - maybe less. 
If you struggle where to get started then take a look at [this article][article] from the `Go` website.

## Test your solution

Use `go test` to verify if your solution is correct.

Correct solution:
```
PASS
ok      xxxx/00-limit-crawler  13.009s
```

where `xxxx` is your repository URL.

An incorrect solution might produce something like the following:
```
--- FAIL: TestMain (7.80s)
        main_test.go:18: There exist two crawl tasks that were executed less than 1 second apart.
	        main_test.go:19: The solution is incorrect.
		FAIL
		exit status 1
		FAIL    xxxx/00-limit-crawler  7.808s
```

[article]: https://github.com/golang/go/wiki/RateLimiting
