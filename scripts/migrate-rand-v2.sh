#!/bin/bash
#
# migrate-rand-v2.sh
#
# Automates the migration from math/rand to math/rand/v2 for Go 1.22+
# This script handles the following changes:
#   1. Updates import statements from "math/rand" to "math/rand/v2"
#   2. Removes "time" imports when only used for rand.Seed
#   3. Removes rand.Seed() calls (auto-seeded in v2)
#   4. Updates function calls: rand.Intn() -> rand.IntN()
#
# Usage:
#   ./scripts/migrate-rand-v2.sh [directory]
#
# Examples:
#   ./scripts/migrate-rand-v2.sh .                    # Migrate entire repository
#   ./scripts/migrate-rand-v2.sh 13-loops/            # Migrate specific directory
#   ./scripts/migrate-rand-v2.sh path/to/file.go     # Migrate single file
#
# Requirements:
#   - Go 1.22 or later
#   - sed (GNU or BSD)
#   - find, grep
#
# Notes:
#   - Creates backups with .bak extension before modifying files
#   - Skips files in vendor/ and .git/ directories
#   - Dry run mode available by setting DRY_RUN=1 environment variable
#

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
BACKUP_EXTENSION=".bak"
DRY_RUN="${DRY_RUN:-0}"

# Get the target directory or file
TARGET="${1:-.}"

if [[ ! -e "$TARGET" ]]; then
    echo -e "${RED}Error: Target '$TARGET' does not exist${NC}" >&2
    exit 1
fi

# Function to print colored messages
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to detect sed version and set appropriate flags
detect_sed() {
    if sed --version >/dev/null 2>&1; then
        # GNU sed
        echo "gnu"
    else
        # BSD sed (macOS default)
        echo "bsd"
    fi
}

SED_TYPE=$(detect_sed)

# Function to perform in-place sed replacement
sed_inplace() {
    local file="$1"
    local pattern="$2"

    if [[ "$SED_TYPE" == "gnu" ]]; then
        sed -i"$BACKUP_EXTENSION" "$pattern" "$file"
    else
        # BSD sed requires a space or empty string after -i
        sed -i "$BACKUP_EXTENSION" "$pattern" "$file"
    fi
}

# Function to migrate a single Go file
migrate_file() {
    local file="$1"
    local modified=0

    # Check if file contains math/rand import (not v2)
    if ! grep -q 'import.*"math/rand"' "$file" && ! grep -q '"math/rand"' "$file"; then
        return 0
    fi

    # Skip if already using v2
    if grep -q 'math/rand/v2' "$file"; then
        log_warning "Already using math/rand/v2: $file"
        return 0
    fi

    log_info "Migrating: $file"

    if [[ "$DRY_RUN" == "1" ]]; then
        echo "  [DRY RUN] Would modify: $file"
        return 0
    fi

    # Step 1: Update import from "math/rand" to "math/rand/v2"
    sed_inplace "$file" 's|"math/rand"|"math/rand/v2"|g'
    modified=1

    # Step 2: Remove rand.Seed() calls and related time.Now() usage
    # Pattern: rand.Seed(time.Now().UnixNano())
    sed_inplace "$file" '/rand\.Seed(time\.Now()\.UnixNano())/d'

    # Also handle variations like:
    # rand.Seed(42)
    sed_inplace "$file" '/rand\.Seed(/d'

    # Step 3: Update function calls: Intn -> IntN
    sed_inplace "$file" 's/rand\.Intn(/rand.IntN(/g'

    # Step 4: Check if time import is still needed
    # Remove time import if it's only used in a single-import statement and no other time references exist
    if ! grep -q 'time\.' "$file" && ! grep -q 'time\s' "$file"; then
        # Remove standalone time import
        sed_inplace "$file" '/^import "time"$/d'

        # Remove time from multi-import block if it's the only thing on its line
        sed_inplace "$file" '/^\s*"time"\s*$/d'
    fi

    # Step 5: Clean up empty lines (more than 2 consecutive empty lines)
    if [[ "$SED_TYPE" == "gnu" ]]; then
        sed_inplace "$file" '/^$/N;/^\n$/D'
    else
        # BSD sed approach - handled differently
        awk 'BEGIN{blank=0} /^$/{blank++; if(blank<=2) print; next} {blank=0; print}' "$file" > "$file.tmp" && mv "$file.tmp" "$file"
    fi

    if [[ $modified -eq 1 ]]; then
        log_success "Migrated: $file"
        echo "  - Updated import: math/rand -> math/rand/v2"
        echo "  - Removed rand.Seed() calls"
        echo "  - Updated rand.Intn() -> rand.IntN()"
        echo "  - Backup saved: $file$BACKUP_EXTENSION"
    fi
}

# Function to find and migrate all Go files
migrate_directory() {
    local dir="$1"
    local count=0

    log_info "Scanning for Go files in: $dir"

    # Find all .go files, excluding vendor and .git directories
    while IFS= read -r -d '' file; do
        migrate_file "$file"
        ((count++))
    done < <(find "$dir" -type f -name "*.go" ! -path "*/vendor/*" ! -path "*/.git/*" -print0)

    log_info "Processed $count Go files"
}

# Main execution
main() {
    echo "========================================"
    echo "  Go math/rand to math/rand/v2 Migration"
    echo "========================================"
    echo ""

    if [[ "$DRY_RUN" == "1" ]]; then
        log_warning "DRY RUN MODE - No files will be modified"
        echo ""
    fi

    log_info "Target: $TARGET"
    log_info "Sed type: $SED_TYPE"
    echo ""

    if [[ -f "$TARGET" ]]; then
        # Single file migration
        migrate_file "$TARGET"
    elif [[ -d "$TARGET" ]]; then
        # Directory migration
        migrate_directory "$TARGET"
    fi

    echo ""
    log_success "Migration complete!"

    if [[ "$DRY_RUN" == "0" ]]; then
        echo ""
        log_info "Next steps:"
        echo "  1. Review the changes: git diff"
        echo "  2. Run tests: go test ./..."
        echo "  3. Remove backup files: find . -name '*.go.bak' -delete"
        echo "  4. Commit changes: git add -A && git commit -m 'refactor: migrate math/rand to math/rand/v2'"
    fi
}

# Run main function
main
