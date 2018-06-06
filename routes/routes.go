package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moos3/golang-web-learning/api"
	"github.com/moos3/golang-web-learning/auth"
	"github.com/urfave/negroni"
)

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API) *mux.Router {

	mux := mux.NewRouter()

	// client static files
	mux.Handle("/", http.FileServer(http.Dir("./client/dist/"))).Methods("GET")
	mux.PathPrefix("/static/js").Handler(http.StripPrefix("/static/js/", http.FileServer(http.Dir("./client/dist/static/js/"))))

	// api
	a := mux.PathPrefix("/api").Subrouter()

	// users
	u := a.PathPrefix("/user").Subrouter()
	u.HandleFunc("/signup", api.UserSignup).Methods("POST")
	u.HandleFunc("/login", api.UserLogin).Methods("POST")
	u.Handle("/info", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	)).Methods("GET")
	u.Handle("/password", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.ChangeUserPassword)),
	)).Methods("POST")

	// quotes
	q := a.PathPrefix("/quote").Subrouter()
	q.HandleFunc("/random", api.Quote).Methods("GET")
	q.Handle("/protected/random", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.SecretQuote)),
	)).Methods("GET")

	// membership
	m := a.PathPrefix("/membership").Subrouter()
	m.Handle("/membership/", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.ListMemberships)),
	)).Methods("GET")
	m.Handle("/membership/add", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.AddMembership)),
	)).Methods("POST")
	m.Handle("/membership/delete", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.DeleteMembership)),
	)).Methods("POST")
	m.Handle("/membership/edit", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.EditMembership)),
	)).Methods("POST")

	return mux
}
