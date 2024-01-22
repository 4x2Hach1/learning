# install
edit `.env`
```bash
task init
go mod tidy
task db-migrate -- migrations/{sql-file-name}.sql
```

# Getting Started
```bash
task run
```
