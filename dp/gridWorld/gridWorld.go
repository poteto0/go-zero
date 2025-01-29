package gridworld

type Grid struct {
	X int
	Y int
}

func (g *Grid) Add(other Grid) Grid {
	return Grid{
		X: g.X + other.X,
		Y: g.Y + other.Y,
	}
}

func (g *Grid) Equal(other Grid) bool {
	return g.X == other.X && g.Y == other.Y
}

var ActionMap = []Grid{
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: -1, Y: 0},
	{X: 1, Y: 0},
}

type GridWorld struct {
	actionSpace   []int
	ActionMeaning map[int]string
	RewardMap     [][]float64
	GoalState     Grid
	WallStates    []Grid
	StartState    Grid
	AgentState    Grid
}

var DefaultGridWorld = &GridWorld{
	actionSpace: []int{0, 1, 2, 3},
	ActionMeaning: map[int]string{
		0: "UP",
		1: "DOWN",
		2: "LEFT",
		3: "RIGHT",
	},
	RewardMap: [][]float64{
		{0.0, 0.0, 0.0, 1.0},
		{0.0, 0.0, 0.0, -1.0},
		{0.0, 0.0, 0.0, 0.0},
	},
	GoalState: Grid{X: 3, Y: 0},
	WallStates: []Grid{
		{X: 1, Y: 1},
	},
	StartState: Grid{X: 0, Y: 2},
	AgentState: Grid{X: 0, Y: 2},
}

func (gw *GridWorld) Height() int {
	return len(gw.RewardMap)
}

func (gw *GridWorld) Width() int {
	return len(gw.RewardMap[0])
}

func (gw *GridWorld) Shape() (int, int) {
	return gw.Height(), gw.Width()
}

func (gw *GridWorld) Actions() []int {
	return gw.actionSpace
}

func (gw *GridWorld) Run(V [][]float64, gamma float64) {
	for y := 0; y < gw.Height(); y++ {
		for x := 0; x < gw.Width(); x++ {
			state := Grid{X: x, Y: y}
			if state.Equal(gw.GoalState) {
				V[y][x] = 0
				continue
			}

			newV := 0.0
			for _, action := range gw.Actions() {
				nextState := gw.Move(state, action)
				reward := gw.Reward(nextState)
				newV += 0.25 * (reward + gamma*V[nextState.Y][nextState.X])
			}
			V[state.Y][state.X] = newV
		}
	}
}

func (gw *GridWorld) Move(state Grid, action int) Grid {
	moving := ActionMap[action]
	nextState := state.Add(moving)

	if !gw.canMove(nextState) {
		return state
	}

	return nextState
}

func (gw *GridWorld) canMove(state Grid) bool {
	if state.X < 0 {
		return false
	}

	if state.X >= gw.Width() {
		return false
	}

	if state.Y < 0 {
		return false
	}

	if state.Y >= gw.Height() {
		return false
	}

	if gw.isWall(state) {
		return false
	}

	return true
}

func (gw *GridWorld) isWall(state Grid) bool {
	for _, wallState := range gw.WallStates {
		if state.Equal(wallState) {
			return true
		}
	}
	return false
}

func (gw *GridWorld) Reward(state Grid) float64 {
	return gw.RewardMap[state.Y][state.X]
}
