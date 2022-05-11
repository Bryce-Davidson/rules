package maps

import (
	"errors"
	"testing"

	"github.com/BattlesnakeOfficial/rules"
	"github.com/stretchr/testify/require"
)

func TestSetupBoard_NotFound(t *testing.T) {
	_, err := SetupBoard("does_not_exist", rules.Settings{}, 10, 10, []string{})

	require.EqualError(t, err, rules.ErrorMapNotFound.Error())
}

func TestSetupBoard_Error(t *testing.T) {
	testMap := StubMap{
		Id:    t.Name(),
		Error: errors.New("bad map update"),
	}
	RegisterMap(testMap.ID(), testMap)

	_, err := SetupBoard(testMap.ID(), rules.Settings{}, 10, 10, []string{})

	require.EqualError(t, err, "bad map update")
}

func TestSetupBoard(t *testing.T) {
	testMap := StubMap{
		Id: t.Name(),
		SnakePositions: map[string]rules.Point{
			"1": {X: 3, Y: 4},
			"2": {X: 6, Y: 2},
		},
		Food: []rules.Point{
			{X: 1, Y: 1},
			{X: 5, Y: 3},
		},
		Hazards: []rules.Point{
			{X: 3, Y: 5},
			{X: 2, Y: 2},
		},
	}
	RegisterMap(testMap.ID(), testMap)

	boardState, err := SetupBoard(testMap.ID(), rules.Settings{}, 10, 10, []string{"1", "2"})

	require.NoError(t, err)

	require.Len(t, boardState.Snakes, 2)

	require.Equal(t, rules.Snake{
		ID:     "1",
		Body:   []rules.Point{{X: 3, Y: 4}, {X: 3, Y: 4}, {X: 3, Y: 4}},
		Health: rules.SnakeMaxHealth,
	}, boardState.Snakes[0])
	require.Equal(t, rules.Snake{
		ID:     "2",
		Body:   []rules.Point{{X: 6, Y: 2}, {X: 6, Y: 2}, {X: 6, Y: 2}},
		Health: rules.SnakeMaxHealth,
	}, boardState.Snakes[1])
	require.Equal(t, []rules.Point{{X: 1, Y: 1}, {X: 5, Y: 3}}, boardState.Food)
	require.Equal(t, []rules.Point{{X: 3, Y: 5}, {X: 2, Y: 2}}, boardState.Hazards)
}

func TestUpdateBoard(t *testing.T) {
	testMap := StubMap{
		Id: t.Name(),
		SnakePositions: map[string]rules.Point{
			"1": {X: 3, Y: 4},
			"2": {X: 6, Y: 2},
		},
		Food: []rules.Point{
			{X: 1, Y: 1},
			{X: 5, Y: 3},
		},
		Hazards: []rules.Point{
			{X: 3, Y: 5},
			{X: 2, Y: 2},
		},
	}
	RegisterMap(testMap.ID(), testMap)

	previousBoardState := &rules.BoardState{
		Turn:    0,
		Food:    []rules.Point{{X: 0, Y: 1}},
		Hazards: []rules.Point{{X: 3, Y: 4}},
		Snakes: []rules.Snake{
			{
				ID:     "1",
				Health: 100,
				Body: []rules.Point{
					{X: 6, Y: 4},
					{X: 6, Y: 3},
					{X: 6, Y: 2},
				},
			},
		},
	}
	boardState, err := UpdateBoard(testMap.ID(), previousBoardState, rules.Settings{})

	require.NoError(t, err)

	require.Len(t, boardState.Snakes, 1)

	require.Equal(t, rules.Snake{
		ID:     "1",
		Body:   []rules.Point{{X: 6, Y: 4}, {X: 6, Y: 3}, {X: 6, Y: 2}},
		Health: rules.SnakeMaxHealth,
	}, boardState.Snakes[0])
	require.Equal(t, []rules.Point{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 5, Y: 3}}, boardState.Food)
	require.Equal(t, []rules.Point{{X: 3, Y: 4}, {X: 3, Y: 5}, {X: 2, Y: 2}}, boardState.Hazards)
}