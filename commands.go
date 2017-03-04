package main

import (
	"encoding/json"
	_ "fmt"
	"math/rand"
	"net/http"
	"time"
)

func respond(res http.ResponseWriter, obj interface{}) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(obj)
}

func handleStart(res http.ResponseWriter, req *http.Request) {
	//data, err := NewGameStartRequest(req)
	_, err := NewGameStartRequest(req)
	// if err != nil {
	// 	respond(res, GameStartResponse{
	// 		Taunt:          toStringPointer("Me eat brains..."),
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
	if err != nil {
		respond(res, GameStartResponse{
			// Taunt:          toStringPointer("Me eat brains..."),
			// Color:          "#f26000",
			Name: "ZombieSnake", //fmt.Sprintf("%v (%vx%v)", data.GameId, data.Width, data.Height),
			// HeadUrl:        toStringPointer(fmt.Sprintf("%v://%v/head.png", scheme, req.Host)),
			// HeadType:       "pixel",
			// TailType:       "tail_type",
			// SecondaryColor: "#ffffff",
		})
	}
}

func handleMove(res http.ResponseWriter, req *http.Request) {
	//data, err := NewMoveRequest(req)
	_, err := NewMoveRequest(req)
	// if err != nil {
	// 	respond(res, MoveResponse{
	// 		Move:  "up",
	// 		Taunt: toStringPointer("You can't hide from me..."),
	// 	})
	// 	return
	// }

	directions := []string{
		"up",
		"down",
		"left",
		"right",
	}

	if err != nil {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		respond(res, MoveResponse{
			Move:  directions[r.Intn(4)],
			Taunt: toStringPointer("You can't hide from me..."),
		})
	}
}
