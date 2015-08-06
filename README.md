# gogalert

Ganglia alerting tool written in go

## Installation

### From source

#### Go installed globally

```bash
$ make deps
$ make fix-paths
$ make build
```

### Go not installed globally

We will create a virtual env with [envirius](https://github.com/ekalinin/envirius):

```bash
$ make env-init
$ make env-build
```
