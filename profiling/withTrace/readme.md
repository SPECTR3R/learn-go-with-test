run

```bash
go build -o mandlebrot
time ./mandlebrot > m.trace
go tool trace m.trace
```
