# Per Apollo



<div align="center">
  <h1> Per Apollo repo </h1>
</div>

**Tips**: 
- need [Go 1.21+](https://golang.org/dl/)
- need [pgsql](https://www.postgresql.org/)


## Install

### Install dependencies
```bash
go mod tidy
```
### build
```bash
cd per_apollo
make
```

### Config env

- copy config-demo.yaml to per.yaml

### start syncer
```bash
./per_apollo syncer
```

### start api
```bash
./per_apollo api
```

## Contribute

TBD
