package main

func move(gs GameState) BattlesnakeMoveResponse {
	okMoves := map[Point]string{
		{X: 0, Y: 1}:  "up",
		{X: 0, Y: -1}: "down",
		{X: 1, Y: 0}:  "right",
		{X: -1, Y: 0}: "left",
	}

	filterForSelf(gs, okMoves)
	filterForWalls(gs, okMoves)
	filterForOthers(gs, okMoves)

	move := "down"

	for _, v := range okMoves {
		move = v
	}

	return BattlesnakeMoveResponse{Move: move}

}

// Modify the move map to filter out wall collision cases.
func filterForWalls(gs GameState, moves map[Point]string) {
	head := gs.You.Head

	for k := range moves {
		dest := head.add(k)
		if dest.X < 0 || dest.X >= gs.Board.Width {
			delete(moves, k)
			continue
		}

		if dest.Y < 0 || dest.Y >= gs.Board.Height {
			delete(moves, k)
			continue
		}
	}
}

// Modify the move map to filter out self collision cases.
func filterForSelf(gs GameState, moves map[Point]string) {
	head := gs.You.Head
	for k := range moves {
		dest := head.add(k)
		if inPointSlice(dest, gs.You.Body) {
			delete(moves, k)
		}
	}
}

func filterForOthers(gs GameState, moves map[Point]string) {
	head := gs.You.Head
	for m := range moves {
		dest := head.add(m)
		for _, s := range gs.Board.Snakes {
			if !(gs.You.ID == s.ID) {
				if gs.You.Length <= s.Length {
					if inPointSlice(dest, s.Body) {
						delete(moves, m)
					}
				}
			}
		}
	}
}

// Check to see if a Point is in a Point slice
func inPointSlice(p Point, ps []Point) bool {
	for _, v := range ps {
		if p == v {
			return true
		}
	}
	return false
}
