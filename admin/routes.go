package admin

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/wanliu/go-oauth2-server/util/routes"
)

// RegisterRoutes registers route handlers for the health service
func (s *Service) RegisterRoutes(router *mux.Router, prefix string) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	routes.AddRoutes(s.GetRoutes(), subRouter)
}

// GetRoutes returns []routes.Route slice for the health service
func (s *Service) GetRoutes() []routes.Route {
	return []routes.Route{
		{
			Name:        "list_users",
			Method:      "GET",
			Pattern:     "/users",
			HandlerFunc: s.listUsers,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "find_users",
			Method:      "GET",
			Pattern:     "/users/query",
			HandlerFunc: s.queryUsers,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "create_user",
			Method:      "POST",
			Pattern:     "/users",
			HandlerFunc: s.createUser,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "edit_user_form",
			Method:      "GET",
			Pattern:     "/users/:id",
			HandlerFunc: s.userForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "edit_user",
			Method:      "PUT",
			Pattern:     "/users/:id",
			HandlerFunc: s.updateUser,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "change_user_password",
			Method:      "PUT",
			Pattern:     "/users/:id/password",
			HandlerFunc: s.changePassword,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "remove_user",
			Method:      "DELETE",
			Pattern:     "/users/:id",
			HandlerFunc: s.removeUser,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
			},
		},
		{
			Name:        "list_clients",
			Method:      "GET",
			Pattern:     "/clients",
			HandlerFunc: s.listClients,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "find_clients",
			Method:      "GET",
			Pattern:     "/clients/query",
			HandlerFunc: s.findClients,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},

		{
			Name:        "create_client_form",
			Method:      "GET",
			Pattern:     "/clients",
			HandlerFunc: s.createClientForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
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
				newClientMiddleware(s),
			},
		},
		{
			Name:        "edit_client_form",
			Method:      "GET",
			Pattern:     "/clients/:id",
			HandlerFunc: s.editClientForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "edit_client",
			Method:      "PUT",
			Pattern:     "/clients/:id",
			HandlerFunc: s.editClient,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "edit_client",
			Method:      "PUT",
			Pattern:     "/clients/:id",
			HandlerFunc: s.editClient,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "remove_client",
			Method:      "DELETE",
			Pattern:     "/clients/:id",
			HandlerFunc: s.deleteClient,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "list_tokens",
			Method:      "GET",
			Pattern:     "/tokens",
			HandlerFunc: s.listTokens,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "update_token_scope_form",
			Method:      "GET",
			Pattern:     "/tokens",
			HandlerFunc: s.changeTokenForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "update_token_scope",
			Method:      "GET",
			Pattern:     "/tokens",
			HandlerFunc: s.changeTokenScopes,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "remove_token",
			Method:      "DELETE",
			Pattern:     "/tokens/:token",
			HandlerFunc: s.deleteTokens,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "list_scopes",
			Method:      "GET",
			Pattern:     "/scopes",
			HandlerFunc: s.listScopes,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "create_scope_form",
			Method:      "GET",
			Pattern:     "/scopes",
			HandlerFunc: s.createScopeForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "create_scope",
			Method:      "POST",
			Pattern:     "/scopes",
			HandlerFunc: s.createScope,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "edit_scope_form",
			Method:      "GET",
			Pattern:     "/scopes/:id",
			HandlerFunc: s.editScopeForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "update_scope",
			Method:      "GET",
			Pattern:     "/scopes/:id",
			HandlerFunc: s.updateScope,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "remove_scope",
			Method:      "DELETE",
			Pattern:     "/scopes/:id",
			HandlerFunc: s.deleteScope,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "list_roles",
			Method:      "GET",
			Pattern:     "/roles",
			HandlerFunc: s.listRoles,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "create_role_form",
			Method:      "GET",
			Pattern:     "/roles/form",
			HandlerFunc: s.createRoleForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "create_role",
			Method:      "POST",
			Pattern:     "/roles",
			HandlerFunc: s.createRole,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "edit_role_form",
			Method:      "GET",
			Pattern:     "/roles/:id",
			HandlerFunc: s.editRoleForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "edit_role",
			Method:      "PUT",
			Pattern:     "/roles/:id",
			HandlerFunc: s.updateRole,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "remove_role",
			Method:      "DELETE",
			Pattern:     "/roles/:id",
			HandlerFunc: s.deleteRole,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newLoggedInMiddleware(s),
				newClientMiddleware(s),
			},
		},
	}
}
