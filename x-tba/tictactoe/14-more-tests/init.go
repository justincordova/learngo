package main

// init is another special function
// Go calls it before the main function
func init() {
	initCells()
}

// initCells initialize the played cells to empty
func initCells() {
	for i := range cells {
		cells[i] = emptyCell
	}
}
