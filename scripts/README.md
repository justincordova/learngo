# Scripts Directory

This directory contains utility scripts for maintaining and updating the Learn Go Programming repository.

## Available Scripts

### migrate-rand-v2.sh

Automates the migration from `math/rand` to `math/rand/v2` for Go 1.22+.

**What it does:**
- Updates import statements from `"math/rand"` to `"math/rand/v2"`
- Removes `"time"` imports when only used for seeding
- Removes `rand.Seed()` calls (auto-seeded in v2)
- Updates function calls: `rand.Intn()` → `rand.IntN()`
- Creates backups before modifying files

**Usage:**

```bash
# Migrate entire repository
./scripts/migrate-rand-v2.sh .

# Migrate specific directory
./scripts/migrate-rand-v2.sh 13-loops/

# Migrate single file
./scripts/migrate-rand-v2.sh path/to/file.go

# Dry run mode (preview changes without modifying files)
DRY_RUN=1 ./scripts/migrate-rand-v2.sh .
```

**Requirements:**
- Go 1.22 or later
- sed (GNU or BSD)
- find, grep

**Example Migration:**

Before:
```go
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    n := rand.Intn(10)
    fmt.Println(n)
}
```

After:
```go
import (
    "fmt"
    "math/rand/v2"
)

func main() {
    n := rand.IntN(10)
    fmt.Println(n)
}
```

**Post-Migration Steps:**

1. Review changes:
   ```bash
   git diff
   ```

2. Run tests to ensure everything works:
   ```bash
   go test ./...
   ```

3. Remove backup files:
   ```bash
   find . -name '*.go.bak' -delete
   ```

4. Commit changes:
   ```bash
   git add -A
   git commit -m 'refactor: migrate math/rand to math/rand/v2'
   ```

**Background:**

Go 1.22 introduced `math/rand/v2` with several improvements:
- Automatic seeding (no need for `rand.Seed()`)
- Better randomness with ChaCha8 algorithm
- Improved API consistency (`IntN` instead of `Intn`)
- Thread-safe by default

The old `math/rand` package is now considered legacy and will eventually be deprecated.

**References:**
- [Go 1.22 Release Notes](https://go.dev/doc/go1.22#math/rand/v2)
- [math/rand/v2 Package Documentation](https://pkg.go.dev/math/rand/v2)
- [Proposal: math/rand/v2](https://github.com/golang/go/issues/61716)

## Contributing

When adding new scripts:

1. Include a header comment explaining the script's purpose
2. Add usage examples in the script's help text
3. Document the script in this README
4. Make the script executable: `chmod +x scripts/your-script.sh`
5. Use proper error handling with `set -euo pipefail`
6. Support dry-run mode when applicable
7. Create backups before modifying files
8. Use colored output for better UX

## Script Conventions

- **Naming:** Use kebab-case for script names
- **Extensions:** `.sh` for bash scripts
- **Shebang:** Always start with `#!/bin/bash`
- **Error handling:** Use `set -euo pipefail` at the top
- **Colors:** Define standard color variables for consistent output
- **Backups:** Always create backups before modifying files
- **Dry run:** Support `DRY_RUN=1` environment variable
- **Documentation:** Include inline comments and usage examples
