# Battlesnake-go

This is my first attempt at the [battlesnake competition](http://battlesnake.io) which I attended in 2017. The battlesnake is written in Go and follows the game server [API documentation](https://stembolthq.github.io/battle_snake/). 

### Running the AI locally

1) [Fork this repo](https://github.com/alexeyza/battlesnake-go/fork).

2) Clone repo to your development environment:
```
git clone git@github.com:USERNAME/battlesnake-go.git $GOPATH/github.com/USERNAME/battlesnake-go
cd $GOPATH/github.com/USERNAME/battlesnake-go
```

3) Compile the battlesnake-go server.
```
go build
```
This will create a `battlesnake-go` executable.

4) Run the server.
```
./battlesnake-go
```

5) Test the client in your browser: [http://127.0.0.1:9000](http://127.0.0.1:9000). I recommend using [Postman](https://www.getpostman.com/) or [Insomnia](https://insomnia.rest/).

Example start game request:
```
{
  "width": 20,
  "height": 20,
  "game_id": "b1dadee8-a112-4e0e-afa2-2845cd1f21aa"
}
:ok
```

Example move request:
```
{
  "you": "2c4d4d70-8cca-48e0-ac9d-03ecafca0c98",
  "width": 2,
  "turn": 0,
  "snakes": [
    {
      "taunt": "git gud",
      "name": "my-snake",
      "id": "2c4d4d70-8cca-48e0-ac9d-03ecafca0c98",
      "health_points": 93,
      "coords": [
        [
          0,
          0
        ],
        [
          0,
          0
        ],
        [
          0,
          0
        ]
      ]
    },
    {
      "taunt": "gotta go fast",
      "name": "other-snake",
      "id": "c35dcf26-7f48-492c-b7b5-94ae78fbc713",
      "health_points": 50,
      "coords": [
        [
          1,
          0
        ],
        [
          1,
          0
        ],
        [
          1,
          0
        ]
      ]
    }
  ],
  "height": 2,
  "game_id": "a2facef2-b031-44ba-a36c-0859c389ef96",
  "food": [
    [
      1,
      1
    ]
  ],
  "dead_snakes": [
    {
      "taunt": "gotta go fast",
      "name": "other-snake",
      "id": "83fdf2b9-c8d0-44f4-acb2-0c506139079e",
      "health_points": 50,
      "coords": [
        [
          5,
          0
        ],
        [
          5,
          0
        ],
        [
          5,
          0
        ]
      ]
    }
  ]
}
:ok
```

### Deploying to Heroku

1. Create a new Go Heroku app using Go buildpack.
```
heroku create [APP_NAME]
```

2. Add a buildpack for Go.
```
heroku buildpacks:set heroku/go
```

3. Push code to Heroku servers. Make sure you have a `vendor/vendor.json`, otherwise Heroku will fail building.
```
git push heroku master
```

4. Open Heroku app in browser.
```
heroku open
```
Or go directly via http://APP_NAME.herokuapp.com

5. View/stream server logs.
```
heroku logs --tail
```

### Questions?

[Email](mailto:battlesnake@sendwithus.com), [Twitter](http://twitter.com/send_with_us)
