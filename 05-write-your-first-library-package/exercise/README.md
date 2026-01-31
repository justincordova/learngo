[Check out the exercise and its solution here.](https://github.com/inancgumus/learngo/tree/master/05-write-your-first-library-package/exercise)

---

# EXERCISE
1. Create a new library
2. In it, create a function that returns the Go version
3. Create a command and import your library
4. Call your function that returns Go version
5. Run your program

## HINTS
**Create your package function like this:**

```go
func Version() string {
    return runtime.Version()
}
```

## EXPECTED OUTPUT
It should print the current Go version on your system.

## WARNING

Create this package in your own module with its own `go.mod` file, not in the `github.com/inancgumus/learngo` folder.

VS Code may auto-import a similarly-named library if it exists in your dependencies. Verify the import path in your code matches your module name declared in `go.mod`.
