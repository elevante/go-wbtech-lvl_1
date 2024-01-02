<h1 align="center">WB Tech Level 1</h1>

<p align="center">
  <a href="https://golang.org/"><img alt="Golang" src="https://badgen.net/badge/language/Golang/blue"></a>
</p>

This repository contains solutions for Level 1 WB Tech assignments, implemented in the Golang programming language. The tasks cover various aspects of programming, such as working with structures, concurrent programming, working with channels, sets, etc.

For each task, there is a folder. In addition to the task, there are unit tests in each folder.

## Running Tests

To run the tests for any task, navigate to the corresponding task folder within the `src` directory and use the `go test` command. For example:

```bash
cd src/<task folder name>
go test
```

Replace `<task folder name>` with the name of the task folder for which you want to run the tests. This command will run all tests in the specified folder.

## Output Examples

Here's an example of what you'll see in the terminal after running the tests:

```bash
$ cd src/task_2
$ go test
ok      go-wbtech-lvl_1/src/task_2      4.031s
```

And here's an example of the test coverage output:

```bash
$ go test -cover
PASS
coverage: 88.9% of statements
ok      go-wbtech-lvl_1/src/task_2      4.037s
```
