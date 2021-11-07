package handlers

import (
	"github.com/go-chi/chi"
)

/*

setupEndpoints is a method of type *Server

If you want to modify the receiver use a pointer like:
func (s *MyStruct) pointerMethod() { } // method on pointer

If you dont need to modify the receiver you can define the receiver as a value like:
func (s MyStruct)  valueMethod()   { } // method on value

*/
func (s *Server) setupEndpoints(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", s.registerUser())
		})
	})
}
