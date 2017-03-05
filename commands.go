package main

import (
	"encoding/json"
	//	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func respond(res http.ResponseWriter, obj interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
}

type H map[string]interface{}

func handleStart(res http.ResponseWriter, req *http.Request) {
	output_data := H{
		"name":            "ZombieSnake",
		"color":           "#009b19",
		"taunt":           "bhaaaa....",
		"head_type":       "sand-worm",
		"tail_type":       "fat-rattle",
		"head_url":        "https://zombiesnake.herokuapp.com/head.jpg",
		"secondary_color": "#f26000",
	}

	json.NewEncoder(res).Encode(output_data)

	//data, err := NewGameStartRequest(req)
	// _, err := NewGameStartRequest(req)
	// if err != nil {
	// 	respond(res, GameStartResponse{
	// 		Taunt:          toStringPointer("I eat brains..."),
	// 		Color:          "#f26000",
	// 		Name:           "ZombieSnake", //fmt.Sprintf("%v (%vx%v)", data.GameId, data.Width, data.Height),
	// 		HeadUrl:        toStringPointer(fmt.Sprintf("%v://%v/head.png")),
	// 		HeadType:       "pixel",
	// 		TailType:       "tail_type",
	// 		SecondaryColor: "#ffffff",
	// 	})
	// }

	// scheme := "http"
	// if req.TLS != nil {
	// 	scheme = "https"
	// }
	// respond(res, GameStartResponse{
	// 	Taunt:          toStringPointer("Me eat brains..."),
	// 	Color:          "#f26000",
	// 	Name:           "ZombieSnake", //fmt.Sprintf("%v (%vx%v)", data.GameId, data.Width, data.Height),
	// 	HeadUrl:        toStringPointer(fmt.Sprintf("%v://%v/head.png", scheme, req.Host)),
	// 	HeadType:       "pixel",
	// 	TailType:       "tail_type",
	// 	SecondaryColor: "#ffffff",
	// })
}

func handleMove(res http.ResponseWriter, req *http.Request) {
	data, _ := NewMoveRequest(req)
	// _, err := NewMoveRequest(req)
	// if err != nil {
	// 	respond(res, MoveResponse{
	// 		Move:  "up",
	// 		Taunt: toStringPointer("You can't hide from me..."),
	// 	})
	// 	return
	// }

	//var board = make([]int, data.Height*data.Width)
	var head Point
	for _, snake := range data.Snakes {
		if data.You == snake.Id {
			head = snake.Coords[0]
		}
	}
	//my_id := data.You

	directions := []string{
		"up",
		"down",
		"left",
		"right",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var move_direction string

	var my_snake Snake
	for _, snake := range data.Snakes {
		if data.You == snake.Id {
			my_snake = snake
		}
	}

	var foundpath bool
	for foundpath = false; foundpath == false; {
		possible_direction := r.Intn(4)
		switch possible_direction {
		case 0:
			if head.Y > 1 && my_snake.Coords[1].Y != head.Y-1 {
				move_direction = directions[0]
				foundpath = true
			}
		case 1:
			if head.Y < data.Height-2 && my_snake.Coords[1].Y != head.Y+1 {
				move_direction = directions[1]
				foundpath = true
			}
		case 2:
			if head.X > 1 && my_snake.Coords[1].X != head.X-1 {
				move_direction = directions[2]
				foundpath = true
			}
		default:
			if head.X < data.Height-2 && my_snake.Coords[1].X != head.X+1 {
				move_direction = directions[3]
				foundpath = true
			}
		}
	}
	// if head.X < data.Width-1 {
	// 	move_direction = directions[3]
	// } else {
	// 	if head.Y < data.Height-1 {
	// 		move_direction = directions[1]
	// 	} else {
	// 		if head.X > 0 {
	// 			move_direction = directions[2]
	// 		}
	// 	}
	// }
	respond(res, MoveResponse{
		//Move:  directions[r.Intn(4)],
		Move:  move_direction,
		Taunt: toStringPointer("Psss...psss..."),
	})

}
