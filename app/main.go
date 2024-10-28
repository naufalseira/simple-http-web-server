package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

// Middleware untuk memeriksa metode HTTP
func middleware(method string, handlerFunc http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method{
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handlerFunc(w,r)
	}
}

// Middleware untuk mencetak log request
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
    })
}

// Handler untuk menampilkan index.html
func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/index.html")
}

// Handler untuk (/contact) endpoint
func contactHandler(w http.ResponseWriter, r * http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w,"ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w, "POST request succesfull\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
	fmt.Fprintf(w, "Message = %s\n", message)
}

// Handler untuk (/jakarta) static endpoint
func jakartaHandler(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Welcome to the jakarta page!")
}

// Handler untuk (/destination/{name}) dynamic endpoint 
func destinationHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    place := vars["place"]
    fmt.Fprintf(w, "Welcome to the %s page!", place)
}

func main(){
	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	fileServer := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	r.PathPrefix("/assets/").Handler(fileServer)

	// Home handler untuk root endpoint (/)
	r.HandleFunc("/", middleware(http.MethodGet, homeHandler))
	r.HandleFunc("/contact", middleware(http.MethodPost, contactHandler))
	r.HandleFunc("/jakarta", middleware(http.MethodGet, jakartaHandler))
	r.HandleFunc("/destination/{place}", middleware(http.MethodGet, destinationHandler))

	fmt.Println("------------------------------------------------------")
	fmt.Println("|      Server running on http://localhost:8090       |")
	fmt.Println("------------------------------------------------------")
	fmt.Println("Endpoint 1 -> http://localhost:8090/")
	fmt.Println("Endpoint 2 -> http://localhost:8090/#contact")
	fmt.Println("Endpoint 3 -> http://localhost:8090/jakarta")
	fmt.Println("Endpoint 4 -> http://localhost:8090/destination/{place}")
	fmt.Println("------------------------------------------------------")
	fmt.Println("> Log request")
	fmt.Print("")

	if err := http.ListenAndServe(":8090", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}