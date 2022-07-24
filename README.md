# mywc

This project is an implementation of the simplest `wc` command in use of `wildcat` as an API.
Please see `goMain` function in `main.go` the details of how to count the words and etc.

Unfortunately, the API of `wildcat` was not refined well.
Moreover, the API document of `wildcat` was not written well.
I will try to resolve above problems about a month (2022-08-31).

## How to use `wildcat`

At first, create the GO project with go module.
To run a command `go get github.com/tamada/wildcat` in your project root directory, adds `wildcat` to your `go.mod` as a module.
Next, edit your go codes and use the API of wildcat on your IDE.