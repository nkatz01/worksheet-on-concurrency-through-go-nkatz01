# Worksheet - Concurrency in Go

In this worksheet you will investigate further the use of concurrency 
via the `Go` programming language.  
You will focus on the use of (so-called) *Goroutines*, 
and their applicability for large scale concurrency programming.

The Go community has plenty resources to read about `go`'s concurrency model 
and how to use it effectively but relatively few exercises.
This set of questions tries to teach concurrency patterns by following the 
"learning by doing" philosophy.

![Image of excited gopher](https://golang.org/doc/gopher/pkg.png)

## Outline

For each of the following exercises you are provided with some outline code.

+ *Only edit the `main.go`* in each case to solve the problem; **do not** modify 
any of the other supplied files (unless you find an error in the supplied code :-( ).
+ If you find a `*_test.go` file present in the folder then you can test the correctness 
of your solution with `go test`.

These exercises should be treated as the "next step on" from the Go-routine exercises you completed
in the first worksheet (so you need to have completed those first before attempting these tasks).

## Exercises

1. "[Limit your Crawler][crawl]" (an extension of the web crawler example from the first worksheet).
2. A [Producer-Consumer][pc] scenario.
3. A "[race condition][race]" in a cache.
4. [Limit service time][time] for free-tier users.


### Credits
 
The majority of these exercises are based upon a set of questions, and outline code, posed by
Long Hoang; many thanks for making such useful resources available under a *free* licence.

[crawl]: 01-limit-crawler
[pc]: 02-producer-consumer
[race]: 03-race-in-cache
[time]: 04-limit-service-time