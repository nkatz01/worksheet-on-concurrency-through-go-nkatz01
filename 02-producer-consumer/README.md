# Producer-Consumer Scenario

The producer reads in tweets from a mock stream and a consumer processes the data to determine
whether someone has tweeted about "golang" (or not). 

Your task is to modify the code inside `main.go` so that the producer and the consumer 
can run concurrently to increase the throughput of the program.

## Expected results:

Before: 
```
davecheney      tweets about golang
beertocode      does not tweet about golang
ironzeb         tweets about golang
beertocode      tweets about golang
vampirewalk666  tweets about golang
Process took 3.580866005s
```

After:
```
davecheney      tweets about golang
beertocode      does not tweet about golang
ironzeb         tweets about golang
beertocode      tweets about golang
vampirewalk666  tweets about golang
Process took 1.977756255s
```
