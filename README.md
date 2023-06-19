# promise

[![Go Reference](https://pkg.go.dev/badge/github.com/ammario/promise.svg)](https://pkg.go.dev/github.com/ammario/promise)


The `promise` package implements a basic promise library for Go. It combines
Go's concurrency primitives with generics to provide a simple interface and
implementation.

## Usage

```go
	p := promise.Go(func() (int, error) {
		time.Sleep(time.Second)
		return 1000, nil
	})
	
    // Do some other work...

    i, err := p.Resolve()
    // i == 1000
    // err == nil
```