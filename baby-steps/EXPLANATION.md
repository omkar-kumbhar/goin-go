## Baby Steps 
This folder will have all the experimental codes which were used while tring out Go language.

#### Lessons to be learnt:

#### first.go

Go is a compiled time language and also handles text UTF8 natively; so any text can be used.
to compile, link and run: 

```go
$ go run first.go
```

You may want to have an executable for future purposes; 
```go
$ go build first.go
$ ./first
```

#### os_args.go

This script tells the usage of command line arguments in go. use `os.Args`
Also checkout the usage of pakages `time` and `strings`

important lines: `time.Now()` and `time.Since(start).Seconds()`
