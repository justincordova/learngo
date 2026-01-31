// ---------------------------------------------------------
// EXERCISE: Wrap File Errors
//
//  Create a program that reads multiple files and wraps
//  any errors with context about which file failed.
//
//  1- Create a function that attempts to read a file
//     and wraps any errors with the filename
//
//  2- Create a function that processes multiple files
//     and wraps errors with additional context
//
//  3- In main, try to read these files:
//     - "config.json" (doesn't exist)
//     - "data.txt" (doesn't exist)
//     - "settings.yaml" (doesn't exist)
//
//  4- Print each error showing the full error chain
//
//
// EXPECTED OUTPUT (similar to):
//
//  Processing files...
//
//  Error processing config.json:
//  failed to process configuration: failed to read file "config.json": open config.json: no such file or directory
//
//  Error processing data.txt:
//  failed to process data file: failed to read file "data.txt": open data.txt: no such file or directory
//
//  Error processing settings.yaml:
//  failed to process settings: failed to read file "settings.yaml": open settings.yaml: no such file or directory
//
// ---------------------------------------------------------

package main

func main() {
}
