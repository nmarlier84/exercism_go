package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	var res int

	for _, rank := range cb[file] {
		if rank {
			res += 1
		}
	}

	return res
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var res int
	if rank < 1 || rank > 8 {
		return 0
	}

	for _, file := range cb {
		if file[rank-1] {
			res += 1
		}
	}

	return res
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var res int
	for _, file := range cb {
		for range file {
			res += 1
		}
	}
	return res
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var res int
	for _, file := range cb {
		for _, v := range file {
			if v {
				res += 1
			}
		}
	}
	return res
}
