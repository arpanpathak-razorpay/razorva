package socket

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/razorva/watson"
	"github.com/rs/cors"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

type Packet struct {
	id   string
	role string
	msg  string
}

func InitializeAndListen(assistant *assistantv2.AssistantV2) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	server, _ := socketio.NewServer(nil)

	fmt.Println("Starting...")
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Connection detected...Socket started:", s.ID())
		return nil
	})

	listen(server, assistant, mux)
}

func listen(server *socketio.Server, assistant *assistantv2.AssistantV2, mux *http.ServeMux) {

	server.OnEvent("/", "send", func(s socketio.Conn, packet Packet) {
		message := packet.msg
		s.SetContext(message)
		fmt.Println("received", message)

		response := watson.SendMessage(assistant, message, packet.id)
		s.Emit("reply", response)
	})

	//close
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	mux.Handle("/socket.io/", server)
	//http.Handle("/socket.io/", server)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"GET", "PUT", "OPTIONS", "POST", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	log.Println("Serving at localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

/*http.Handle("/socket.io/", server)
http.Handle("/", http.FileServer(http.Dir("./asset")))
log.Println("Serving at localhost:8000...")
log.Fatal(http.ListenAndServe(":8000", nil))*/
