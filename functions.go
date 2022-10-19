package main

func move(gs GameState) BattlesnakeMoveResponse {
	okMoves := map[Coord]string{
		{X: 0, Y: 1}:  "up",
		{X: 0, Y: -1}: "down",
		{X: 1, Y: 0}:  "right",
		{X: -1, Y: 0}: "left",
	}

	filterForSelf(gs, okMoves)
	filterForWalls(gs, okMoves)

	move := "down"

	for _, v := range okMoves {
		move = v
	}

	return BattlesnakeMoveResponse{Move: move}

}

// Return a slice of strings of possible moves after checking for walls
func filterForWalls(gs GameState, moves map[Coord]string) {
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

// Modify the move map to filter out collision cases

func filterForSelf(gs GameState, moves map[Coord]string) {
	head := gs.You.Head
	for k := range moves {
		dest := head.add(k)
		if inCoordSlice(dest, gs.You.Body) {
			delete(moves, k)
		}
	}
}

func inCoordSlice(p Coord, ps []Coord) bool {
	for _, v := range ps {
		if p.X == v.X && p.Y == v.Y {
			return true
		}
	}
	return false
}
