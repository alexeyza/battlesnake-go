package main

import (
	"encoding/json"
	//"fmt"
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
	data := H{
		"name":            "ZombieSnake",
		"color":           "#009b19",
		"taunt":           "bhaaaa....",
		"head_type":       "sand-worm",
		"tail_type":       "fat-rattle",
		"head_url":        "http://10.189.212.94:9000/head.jpg",
		"secondary_color": "#f26000",
	}

	json.NewEncoder(res).Encode(data)

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
	//data, err := NewMoveRequest(req)
	// _, err := NewMoveRequest(req)
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

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	respond(res, MoveResponse{
		Move:  directions[r.Intn(4)],
		Taunt: toStringPointer("Psss...psss..."),
	})
}
