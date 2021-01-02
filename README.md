# Crow
`crow` is a simple command-line utility that lets you run arbitrary commands when certain files change.

![Crow Banner](../assets/banner.png)

## Demo
A demonstration of crow being used to re-execute tests every time a file is saved. Also, see other [use cases](#use-cases).

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
or pipe in a list of filenames to watch from `stdin` from `fd`, `find`, `ls`, `echo`, etc...
```
filenames | crow command
```

### Use cases

Use `crow` to run tests once you save `main.go`.
```
crow -w main.go go test ./...
```
```
echo main.go | crow go test ./...
```

Automatically restart your server on changes (watches all files in the current directory).
```
crow go run main.go
```

Live preview markdown in your terminal with [glow](https://github.com/charmbracelet/glow).
```
crow -w README.md glow README.md
```
```
fd .md | crow glow README.md
```

Use `crow` with `!!` to watch files and run the last command.
```bash
crow !!
```

## Alternatives
* [entr](https://github.com/eradman/entr/)
* [reflex](https://github.com/cespare/reflex)

## Contributing
Pull requests are welcome.

## License
[MIT](https://choosealicense.com/licenses/mit/)
