# Learn Go Programming - Modern Go 1.25 Edition

**A comprehensive collection of Go examples, exercises, and quizzes - modernized for Go 1.25**

The best way to learn is by doing. Inside this repository, you will find hundreds of Go examples, exercises, and quizzes covering everything from basics to advanced topics including the latest Go 1.25 features.

---

## What's Inside

- 1,000+ lines of example code across 40+ topics
- Hands-on exercises with detailed solutions
- Modern Go practices using Go 1.25 features
- Real-world patterns for concurrency, error handling, and more
- Progressive learning path from beginner to advanced

---

## About This Repository

This repository is based on the excellent [learngo](https://github.com/inancgumus/learngo) repository by [Inanc Gumus](https://github.com/inancgumus), expanded and modernized for Go 1.25 using Claude Code.

---

## Getting Started

### Prerequisites

- Go 1.25+ installed ([Download here](https://go.dev/dl/))
- A code editor (VS Code, GoLand, Vim, etc.)
- Basic programming knowledge (helpful but not required)

### Quick Start

```bash
# Clone the repository
git clone https://github.com/yourusername/learngo.git
cd learngo

# Start with the basics
cd 01-get-started
go run main.go

# Or jump to modern features
cd 29-concurrency/01-goroutines-basics
go run main.go
```

### Running Examples

Each section contains runnable examples:

```bash
cd <section-name>
go run main.go
```

### Running Tests

Some sections include tests (especially concurrency examples):

```bash
cd 29-concurrency/07-testing-concurrent-code
go test -v
```

### Using the Race Detector

For concurrency examples, use the race detector:

```bash
go run -race main.go
go test -race -v
```

---

## Learning Path

### Fundamentals (Sections 1-10)
Get started with Go basics: variables, types, constants, and control flow.

### Core Concepts (Sections 11-20)
Master arrays, slices, strings, and complete hands-on projects.

### Advanced Topics (Sections 21-26)
Deep dive into maps, structs, functions, and pointers.

### Modern Go (Sections 27-31)
Learn error handling, generics, concurrency, context, and Go 1.25 features.

---

## Table of Contents

### Basics (1-10)
- 01-get-started
- 02-write-your-first-program
- 03-packages-and-scopes
- 04-statements-expressions-comments
- 05-write-your-first-library-package
- 06-variables
- 07-printf
- 08-numbers-and-strings
- 09-go-type-system
- 10-constants

### Control Flow & Projects (11-21)
- 11-if
- 12-switch
- 13-loops
- 14-arrays
- 15-project-retro-led-clock
- 16-slices
- 17-project-empty-file-finder
- 18-project-bouncing-ball
- 19-strings-runes-bytes
- 20-project-spam-masker
- 21-project-text-wrapper

### Data Structures & Functions (22-26)
- 22-maps
- 23-input-scanning
- 24-structs
- 25-functions
- 26-pointers

### Modern Go Features (27-31)
- **27-error-handling** - Error wrapping, inspection, custom errors
- **28-generics** - Type parameters, constraints, generic types
- **29-concurrency** - Goroutines, channels, patterns, Go 1.25 features
- **30-context** - Cancellation, timeouts, request-scoped values
- **31-modern-stdlib** - Go 1.25 stdlib features (json/v2, CSRF, reflection)

---

## How to Use This Repository

### For Beginners
1. Start with **01-get-started** and work through sequentially
2. Complete each section's exercises before moving on
3. Don't skip the projects - they solidify your learning
4. Read the README in each section for context

### For Experienced Programmers
1. Jump to **27-error-handling** for modern Go patterns
2. Explore **28-generics** if you're familiar with other generic languages
3. Master **29-concurrency** for Go's killer feature
4. Check out **31-modern-stdlib** for the latest Go 1.25 innovations

### For Go 1.25 Features
1. **29-concurrency/06-waitgroup-go-method** - New WaitGroup.Go() method
2. **29-concurrency/07-testing-concurrent-code** - testing/synctest package
3. **31-modern-stdlib/** - All Go 1.25 features

---


## Additional Resources

### Official Go Resources
- [Official Go Website](https://go.dev/)
- [Go Tour](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)

### Go 1.25 Specific
- [Go 1.25 Release Notes](https://go.dev/doc/go1.25)
- [Go 1.25 Blog Post](https://go.dev/blog/go1.25)
- [Go 1.25 Interactive Tour](https://antonz.org/go-1-25/)

### Original Repository
- [Inanc Gumus's learngo](https://github.com/inancgumus/learngo)

---

## Acknowledgments

- **[Inanc Gumus](https://github.com/inancgumus)** - Original learngo repository and excellent Go teaching
- **[Anthropic](https://anthropic.com)** - Claude Code used for modernization and expansion
- **Go Team** - For creating an amazing language and continuously improving it

---

## Repository Stats

- Original sections: 26
- New sections added: 5
- Total sections: 31
- Examples: 40+
- Exercises: 15+
- Lines of code: 5,000+
- Go version: 1.25.6

---

**Happy Learning!**

*Remember: The best way to learn Go is by writing Go. Start coding today!*
