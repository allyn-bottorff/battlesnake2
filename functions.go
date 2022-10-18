package main

func move(gs GameState) BattlesnakeMoveResponse {
	okMoves := []Coord{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 0, Y: -1},
	}

	okMoves = filterForSelf(gs, okMoves)
	okMoves = filterForWalls(gs, okMoves)

	move := "down"

	moveMap := map[string]Coord{
		"up":    {X: 0, Y: 1},
		"down":  {X: 0, Y: -1},
		"right": {X: 1, Y: 0},
		"left":  {X: -1, Y: 0},
	}

	for _, v := range okMoves {

		switch {
		case v.equals(moveMap["up"]):
			move = "up"
		case v.equals(moveMap["right"]):
			move = "right"
		case v.equals(moveMap["down"]):
			move = "down"
		case v.equals(moveMap["left"]):
			move = "left"
		}
	}

	return BattlesnakeMoveResponse{Move: move}

}

// Return a slice of strings of possible moves after checking for walls
func filterForWalls(gs GameState, moves []Coord) []Coord {
	head := gs.You.Head

	var okMoves []Coord
	for _, v := range moves {
		dest := head.add(v)
		if dest.X > 0 && dest.X < gs.Board.Width {
			if dest.Y > 0 && dest.Y < gs.Board.Height {
				okMoves = append(okMoves, v)
			}
		}
	}
	return okMoves

}

// up:    [1,0]
// right: [0,1]
// down:  [-1,0]
// left:  [0,-1]
func filterForSelf(gs GameState, moves []Coord) []Coord {
	head := gs.You.Head
	var okMoves []Coord
	for _, v := range moves {
		dest := head.add(v)
		if !inCoordSlice(dest, gs.You.Body) {
			okMoves = append(okMoves, v)
		}
	}
	return okMoves
}

func inCoordSlice(p Coord, ps []Coord) bool {
	for _, v := range ps {
		if p.X == v.X && p.Y == v.Y {
			return true
		}
	}
	return false
}
