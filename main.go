package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (
	upgrader      = websocket.Upgrader{}
	connectTime   = 0
	authenticated = false
)
var socket *websocket.Conn

type msg struct {
	Num   int
	Token string
}

func main() {
	e := echo.New()
	e.GET("/", sampleAPIHandler)
	e.GET("/ws", wsHandler)
	e.Logger.Fatal(e.Start(":8080"))

}

func sampleAPIHandler(c echo.Context) error {
	fmt.Println("System started again")
	godotenv.Load()
	fmt.Println(os.Getenv("ENV_NAME"))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"val": "Hello, World 2!",
	})
}

// https://stackoverflow.com/questions/4361173/http-headers-in-websockets-client-api
// Will use protocol fields as auth
func wsHandler(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	connectTime = time.Now().Second()
	time.AfterFunc(30*time.Second, func() {
		if !authenticated {
			//Unauthenticated, close socket
			ws.Close()
		}
	})
	socket = ws
	defer ws.Close()

	for {
		m := msg{}
		err := ws.ReadJSON(&m)
		if err != nil {
			//Log error since connection got closed
			return nil
		}
		fmt.Printf("Got message: %#v\n", m)
		if m.Token != "" {
			authenticated = true
		}

		if err = ws.WriteJSON(m); err != nil {
			//Log error since connection got closed
			return nil
		}

	}
}
