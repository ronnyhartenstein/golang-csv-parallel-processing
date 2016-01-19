# golang-csv-parallel-processing

Example for CSV processing in Golang using goroutines and sync.WaitGroup

For blog post ["Parallel processing of CSV in Elixir and Golang"](http://blog.rh-flow.de/2016/01/19/parallel-processing-of-csv-in-elixir-and-golang/).

**Generate some test data**

```
go build csvparallel/generate && ./generate -n 1000 -f test.csv
```

**Test importer**

```
go build csvparallel/import && ./import -f test.csv
```
