## Go Maro

a link shortener

### Requirement
- Go 1.24
  - [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html)
  - [air](https://github.com/air-verse/air) (optional for hot reloading)

- npm or pnpm

### Setup
1. clone the project 
```bash
git clone https://github.com/maroisa/go-maro.git && cd go-maro
```
2. Download the required dependencies
```bash
go mod download
pnpm install # or npm install
```

3. Run the project
```bash
air # if air installed
go run cmd/main.go # if not
```
