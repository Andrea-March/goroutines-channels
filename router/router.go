package router

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"github.com/routines"
	"github.com/rs/cors"
	"net/http"
)

type Message struct {
	Text string
}


type NoCache struct {
	Handler http.Handler
}
func (m NoCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Add("Pragma", "no-cache")
	w.Header().Add("Expires", "0")
	m.Handler.ServeHTTP(w, r)
}

type Router struct {
	*httprouter.Router
	Shutdown        func()
}

func SupportCORS(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowOriginFunc: func(string) bool {
			return true
		},
		AllowedHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept"}})
	return c.Handler(handler)
}

func New() (r *Router, err error) {
	r = &Router{
		Router: httprouter.New(),
	}
	r.GET("/", serveIndex)
	r.POST("/go",executeRoutine)
	return r, nil
}

func serveIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w,r,"static/index.html")
}

func executeRoutine(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Write([]byte("Executing Routine..."))
	message := Message{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		glog.Error("Error in parsing request, exiting...")
	}
	go routines.ExecuteGoRoutine(message.Text)
}
