# Talks

This repository contains a collection of talk slides and other related materials.

## Content

1. [Introduction to DDD](introduction-to-ddd.slide)
1. [Course Go](course-go.slide)

## Tooling

The slides are made using the [go present syntax](https://pkg.go.dev/golang.org/x/tools/present).
You can run the slides locally using the go present CLI command.

### Installation

You can install the CLI command via go install like so:

```sh
go install golang.org/x/tools/cmd/present@latest
```

This installs the **present** executable into your
**$GOPATH/bin** directory if **GOPATH** is set or the
**$HOME/go/bin** directory otherwise.

### Usage

Make sure that the directory with the **present** binary is in your $PATH.
After that, running the present tool is simple:

```sh
present
```

A webserver is run on localhost where you can preview
the slides using your favourite web browser.
