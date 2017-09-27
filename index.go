package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"html/template"
	"fmt"
)

var tpl *template.Template

type message struct {
	Message string `json:"message"`
}

func index(w http.ResponseWriter, r http.Response) {
	tpl.Execute(w, "index")
}

func socket(ws *websocket.Conn) {
	for {
		// allocate our container struct
		var m message

		// receive a message using the codec

		if err := ws.ReadJSON(&m); err != nil {
			log.Println(err)
			break
		}

		log.Println("Received message:", m.Message)

		// send a response
		m2 := message{fmt.Sprintf("I got the message: %s", m.Message)}
		if err := ws.WriteJSON(m2); err != nil {
			log.Println(err)
			break
		}
	}
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	http.HandlerFunc("/", index)
	http.Handle("/socket", websocket.)
	panic(http.ListenAndServe(":9000", nil))
}