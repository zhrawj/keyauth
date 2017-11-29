package router

import (
	"openauth/api/http/handler"
)

// RouteToDomain use to deal message
func (r *MyRouter) RouteToDomain() {
	// validate message
	r.HandlerFunc("GET", "/domains/", handler.ListDomain)
	r.HandlerFunc("GET", "/domains/:did/", handler.GetDomain)
	r.HandlerFunc("DELETE", "/domains/:did/", handler.DeleteDomain)
	r.HandlerFunc("POST", "/domains/:did/", handler.CreateDomain)
}