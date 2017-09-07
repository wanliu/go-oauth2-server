package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/wanliu/go-oauth2-server/util/routes"
)

// RegisterRoutes registers route handlers for the health service
func (s *Service) RegisterRoutes(router *mux.Router, prefix string) {
	subRouter := router.PathPrefix(prefix).Subrouter()

	indexRoute := s.getIndexRoute()
	if indexRoute != nil {
		handler := s.getHandler(indexRoute)
		router.Handle(prefix, handler)
	}

	routes.AddRoutes(s.GetRoutes(), subRouter)
}

func (s *Service) getHandler(route *routes.Route) http.Handler {
	var (
		handler http.Handler
		n       *negroni.Negroni
	)

	// Add any specified middlewares
	if len(route.Middlewares) > 0 {
		n = negroni.New()

		for _, middleware := range route.Middlewares {
			n.Use(middleware)
		}

		// Wrap the handler in the negroni app with middlewares
		n.Use(negroni.Wrap(route.HandlerFunc))
		handler = n
	} else {
		handler = route.HandlerFunc
	}

	return handler
}

func (s *Service) getIndexRoute() *routes.Route {
	for _, router := range s.GetRoutes() {
		router := router
		if router.Name == "index" {
			return &router
		}
	}
	return nil
}

// GetRoutes returns []routes.Route slice for the health service
func (s *Service) GetRoutes() []routes.Route {
	return []routes.Route{
		{
			Name:        "register_form",
			Method:      "GET",
			Pattern:     "/register",
			HandlerFunc: s.registerForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				// newClientMiddleware(s),
			},
		},
		{
			Name:        "register",
			Method:      "POST",
			Pattern:     "/register",
			HandlerFunc: s.register,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				// newClientMiddleware(s),
			},
		},
		{
			Name:        "login_form",
			Method:      "GET",
			Pattern:     "/login",
			HandlerFunc: s.loginForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "login",
			Method:      "POST",
			Pattern:     "/login",
			HandlerFunc: s.login,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "logout",
			Method:      "GET",
			Pattern:     "/logout",
			HandlerFunc: s.logout,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
			},
		},
		{
			Name:        "authorize_form",
			Method:      "GET",
			Pattern:     "/authorize",
			HandlerFunc: s.authorizeForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "authorize",
			Method:      "POST",
			Pattern:     "/authorize",
			HandlerFunc: s.authorize,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "index",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: s.index,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				// newClientMiddleware(s),
			},
		},
		{
			Name:        "update_user",
			Method:      "POST",
			Pattern:     "/",
			HandlerFunc: s.updateUser,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				// newClientMiddleware(s),
			},
		},
		{
			Name:        "list_clients",
			Method:      "GET",
			Pattern:     "/clients",
			HandlerFunc: s.listClient,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				// newClientMiddleware(s),
			},
		},
		{
			Name:        "new_client",
			Method:      "GET",
			Pattern:     "/clients/new",
			HandlerFunc: s.newClientForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				// newClientMiddleware(s),
			},
		},
		{
			Name:        "create_client",
			Method:      "POST",
			Pattern:     "/clients",
			HandlerFunc: s.createClient,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				// newClientMiddleware(s),
			},
		},
		{
			Name:        "client_detail",
			Method:      "GET",
			Pattern:     "/clients/{id}",
			HandlerFunc: s.clientDetail,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				// newClientMiddleware(s),
			},
		},
		{
			Name:        "delete_client",
			Method:      "DELETE",
			Pattern:     "/clients/:id",
			HandlerFunc: s.deleteClient,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				// newClientMiddleware(s),
			},
		},
	}
}
