# Remaining go vet Issues

This document catalogs the remaining `go vet` issues found in the repository after automated fixes.

**Generated:** 2026-01-30
**Total Issues:** 23

## Summary by Category

| Category | Count | Severity | Auto-fixable |
|----------|-------|----------|--------------|
| Non-constant format strings | 2 | Medium | No |
| Example functions with unknown identifiers | 20 | Low | No |
| Malformed struct tags | 1 | High | No |

---

## 1. Non-constant Format Strings (2 issues)

### Issue Description
Using non-constant format strings in `fmt.Printf`/`fmt.Fprintf` can lead to security vulnerabilities and runtime errors if the format string contains user input or dynamic content.

### Affected Files

#### 1.1 logparser/functional/textwriter.go:41:18
```
logparser/functional/textwriter.go:41:18: non-constant format string in call to fmt.Fprintf
```

**Severity:** Medium
**Recommendation:** Review the format string usage. If the format is truly dynamic, consider using `fmt.Fprint` instead, or validate/sanitize the format string.

#### 1.2 x-tba/tictactoe-experiments/06-refactor/main.go:44:15
```
x-tba/tictactoe-experiments/06-refactor/main.go:44:15: non-constant format string in call to fmt.Printf
```

**Severity:** Medium
**Recommendation:** Same as above. Review if the dynamic format is necessary or if it can be replaced with a constant format string.

---

## 2. Example Functions with Unknown Identifiers (20 issues)

### Issue Description
Example functions in test files reference identifiers that don't exist in the package. This typically happens in tutorial/learning code where examples are written before the actual implementation, or when code is refactored and examples aren't updated.

### Affected Files

All issues are in `x-tba/tictactoe/*/board_test.go` files:

| File | Line | Function | Missing Identifier |
|------|------|----------|-------------------|
| x-tba/tictactoe/05-testing/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/06-if-switch/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/07-loop/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/08-multi-loop/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/09-slices/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/09-slices/board_test.go | 32 | ExamplePrintBoardCells | PrintBoardCells |
| x-tba/tictactoe/10-arrays/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/10-arrays/board_test.go | 32 | ExamplePrintBoardCells | PrintBoardCells |
| x-tba/tictactoe/11-randomization/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/11-randomization/board_test.go | 31 | ExamplePrintBoardCells | PrintBoardCells |
| x-tba/tictactoe/12-infinite-loop/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/12-infinite-loop/board_test.go | 31 | ExamplePrintBoardCells | PrintBoardCells |
| x-tba/tictactoe/13-detect-winning/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/13-detect-winning/board_test.go | 31 | ExamplePrintBoardCells | PrintBoardCells |
| x-tba/tictactoe/14-more-tests/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/14-more-tests/board_test.go | 31 | ExamplePrintBoardCells | PrintBoardCells |
| x-tba/tictactoe/15-os-args/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/15-os-args/board_test.go | 31 | ExamplePrintBoardCells | PrintBoardCells |
| x-tba/tictactoe/16-types/board_test.go | 14 | ExamplePrintBoard | PrintBoard |
| x-tba/tictactoe/16-types/board_test.go | 31 | ExamplePrintBoardCells | PrintBoardCells |

**Severity:** Low
**Impact:** These issues don't affect runtime behavior but prevent `go test` from running these examples and may confuse learners.

**Recommendation Options:**
1. **Add the missing functions** to each package (if they're meant to exist as part of the tutorial progression)
2. **Remove or comment out the example functions** if they're placeholders
3. **Update the examples** to reference functions that actually exist in each stage

**Note:** Since this is in the `x-tba` (to-be-announced) directory, these appear to be work-in-progress tutorial materials. The decision on how to fix these should align with the pedagogical goals of each tutorial stage.

---

## 3. Malformed Struct Tags (1 issue)

### Issue Description
Struct field tags must follow a specific format: `key:"value"`. The tag `json:name` is missing quotes around the value.

### Affected Files

#### 3.1 x-tba/wizards-structs/unmarshal/main.go:21:2
```
x-tba/wizards-structs/unmarshal/main.go:21:2: struct field tag `json:name` not compatible with reflect.StructTag.Get: bad syntax for struct tag value
```

**Severity:** High
**Impact:** This tag will be silently ignored at runtime, causing JSON unmarshaling to fail or behave incorrectly.

**Fix:** Change `json:name` to `json:"name"`

**Example:**
```go
// Wrong:
type Person struct {
    Name string `json:name`
}

// Correct:
type Person struct {
    Name string `json:"name"`
}
```

---

## Resolution Strategy

### Immediate Action Required (High Priority)
1. **Fix malformed struct tag** in `x-tba/wizards-structs/unmarshal/main.go` - This is a clear bug

### Medium Priority
2. **Review non-constant format strings** - Determine if they're security concerns or intentional design choices

### Low Priority (Design Decision Needed)
3. **Resolve example function issues** - Requires understanding the tutorial structure and pedagogical intent

---

## Notes

- All issues are in experimental or tutorial directories (`x-tba/`)
- None of the issues are in production or core library code
- The `x-tba` directory appears to contain work-in-progress educational materials
- Some issues may be intentional for teaching purposes (showing broken code that students fix)

---

## Next Steps

1. Review each issue category with the repository maintainer
2. Determine which issues are intentional (for teaching) vs. actual bugs
3. Create individual tasks/issues for fixes that should be made
4. Update this document as issues are resolved
