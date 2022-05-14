package main

import (
	"log"
	"net/http"
	"sync"
	"os"
)

var is_sleeping bool
var hostname string
var mutex = &sync.Mutex{}

func wake_up(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Server", hostname)
	mutex.Lock()
	is_sleeping = false
	log.Print("some made some noise and croc is woken up now")
	http.ServeFile(w, r, "croc_awake.html")
	mutex.Unlock()
}

func sleep(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Server", hostname)
	mutex.Lock()
	is_sleeping = true
	log.Print("someone sang lullaby")
	http.ServeFile(w, r, "croc_sleeping.html")
	mutex.Unlock()
}

func main() {

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Server", hostname)
		if(is_sleeping) {
			log.Print("croc is sleeping")
			http.ServeFile(w, r, "croc_sleeping.html")
		} else {
			log.Print("croc is awake")
			http.ServeFile(w, r, "croc_awake.html")
		}

	})

	http.HandleFunc("/wake_up", wake_up)
	http.HandleFunc("/sleep", sleep)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Fatal(http.ListenAndServe(":80", nil))

}
