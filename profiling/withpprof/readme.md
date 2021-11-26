# Trace commands

```bash
go build -o mandlebrot
time ./mandlebrot > cpu.pprof
╰─○ λ go tool pprof cpu.pprof
╰─○ λ (pprof) top -cum
╰─○ λ (pprof) list main.main
```
