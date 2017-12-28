package hubx

import (
	"fmt"
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "home.html")
}

type User struct {
	Name string
}

func startServer() {
	hub := NewHub(1)
	// hub.BeforeJoin(func(client IClient) error {
	// 	return nil
	// })
	hub.Use(func(m *ClientHubMessage, next func()) {
		fmt.Printf("use before %+v\n", m)
		next()
		fmt.Printf("user after %+v\n", m)
	})
	hub.Use(func(m *ClientHubMessage, next func()) {
		fmt.Printf("use before2 %+v\n", m)
		next()
		fmt.Printf("user after2 %+v\n", m)
	})
	hub.On("im", func(m *ClientHubMessage) {
		var str string
		err := m.Decode(&str)
		fmt.Println(str, err)
		m.Client.Send("im", m.Client.Get("user").(*User).Name+" say:"+str)
		if "close" == str {
			fmt.Println("close client")
			go func() { hub.UnregisterChan() <- m.Client }()
		}
	})

	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		clientProp := map[interface{}]interface{}{"user": &User{Name: "jim"}}
		ServeWs(hub, w, r, clientProp)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
