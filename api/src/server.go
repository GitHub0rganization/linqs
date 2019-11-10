package server

import (
	"fmt"

	"github.com/GitHub0rganization/WebApp/sample"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/justinas/alice"

	"firebase.google.com/go/auth"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
	"github.com/GitHub0rganization/WebApp/controller"
	"github.com/GitHub0rganization/WebApp/db"
	"github.com/GitHub0rganization/WebApp/middleware"
)

type Server struct {
	db         *sqlx.DB
	router     *mux.Router
	authClient *auth.Client
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init(datasource string) {

	cs := db.NewDB(datasource)
	dbcon, err := cs.Open()
	if err != nil {
		log.Fatalf("failed db init. %s", err)
	}
	s.db = dbcon
	s.router = s.Route()
}

func (s *Server) Run(addr string) {
	log.Printf("Listening on port %s", addr)
	err := http.ListenAndServe(
		fmt.Sprintf(":%s", addr),
		handlers.CombinedLoggingHandler(os.Stdout, s.router),
	)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Route() *mux.Router {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Authorization"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	})

	commonChain := alice.New(
		middleware.RecoverMiddleware,
		corsMiddleware.Handler,
	)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/public").Handler(commonChain.Then(sample.NewPublicHandler()))

	guiltController := controller.NewGuiltController(s.db)
	r.Methods(http.MethodGet).Path("/guilts").Handler(commonChain.Then(AppHandler{guiltController.All}))

	r.PathPrefix("").Handler(commonChain.Then(http.StripPrefix("/img", http.FileServer(http.Dir("./img")))))
	return r
}
