# Crow
`crow` is a simple command-line utility that lets you run arbitrary commands when certain files change.

![Crow Banner](../assets/banner.png)

## Demo
A demonstration of crow being used to re-execute tests every time the file is saved. See other [use cases](#use-cases).

![Crow Demo](../assets/crow.gif)

## Installation
### Install with `go get`
```
go get github.com/maaslalani/crow
```

### Install from source

Clone this repository and `cd` into it.
```
git clone git@github.com:maaslalani/crow.git && cd crow
```

Install `crow` with go install.
```
go install
```

Ensure `~/go/bin` is in your `PATH`.

## Usage
```
crow [--watch path] [--ext extensions] command
```

### Use cases

Use `crow` to run tests once you save `main.go`.
```
crow -w main.go go test ./...
```

Automatically restart your server on changes (watches all files in the current directory).
```
crow go run main.go
```

Live preview markdown in your terminal with [glow](https://github.com/charmbracelet/glow).
```
crow -w README.md glow README.md
```

Use `crow` with `!!` to watch files and run the last command.
```bash
crow !!
```

## Contributing
Pull requests are welcome.

## License
[MIT](https://choosealicense.com/licenses/mit/)
