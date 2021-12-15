package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var occupied int
	for _, isCellOccupied := range cb[rank] {
		if isCellOccupied {
			occupied++
		}
	}
	return occupied
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
		return 0
	}
	file = file - 1
	var occupied int
	for _, rankCells := range cb {
		if rankCells[file] {
			occupied++
		}
	}
	return occupied
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var cellCount int
	for rankKey := range cb {
		for range cb[rankKey] {
			cellCount++
		}
	}
	return cellCount
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var totalOccupied int
	for rankKey := range cb {
		totalOccupied += CountInRank(cb, rankKey)
	}
	return totalOccupied
}
