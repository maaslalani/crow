# Crow
`crow` is a simple command-line utility that lets you run arbitrary commands when certain files change.

## Installation

Clone this repository.

```bash
git clone git@github.com:maaslalani/crow.git
```

Install `crow` with go install.
```bash
go install crow
```

## Usage
```
crow [--watch path] command
```

### Use cases

Use `crow` to run tests once you save `main.go`.
```
crow -w main.go go test ...
```

Automatically restart your server on changes.
```
crow go run main.go
```

Live preview markdown in your terminal with [glow](https://github.com/charmbracelet/glow).
```
crow -w README.md glow README.md
```

## Contributing
Pull requests are welcome.

## License
[MIT](https://choosealicense.com/licenses/mit/)
