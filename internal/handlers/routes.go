package handlers

import (
	"fmt"
	"lemmego/api"
)

func Routes(r *api.Router) {
	//r.Get("/", func(c *api.Context) error {
	//	return c.Send(200, []byte(c.Query("code")))
	//})
	r.Get("/oauth/clients/create", OauthClientCreateHandler)
	r.Post("/oauth/clients", OauthClientStoreHandler)
	r.Get("/oauth/authorize", AuthorizeIndexHandler)
	r.Post("/register", RegistrationStoreHandler)
	r.Get("/api/v1/test1", func(c *api.Context) error {
		fmt.Println("inside test1")
		return c.Send(200, []byte("test1"))
	}).UseBefore(func(next api.Handler) api.Handler {
		return func(c *api.Context) error {
			fmt.Println("before test1")
			return next(c)
		}
	}).UseAfter(func(next api.Handler) api.Handler {
		return func(c *api.Context) error {
			fmt.Println("after test1")
			return next(c)
		}
	})

	//r.Group("/api", func(r *api.Router) {
	//	r.Group("/v1", func(r *api.Router) {
	//		r.Get("/test1", func(c *api.Context) error {
	//			fmt.Println("inside test1")
	//			return c.Send(200, []byte("test1"))
	//		})
	//	}).UseBefore(func(next api.Handler) api.Handler {
	//		return func(c *api.Context) error {
	//			fmt.Println("before v1")
	//			return next(c)
	//		}
	//	}).UseAfter(func(next api.Handler) api.Handler {
	//		return func(c *api.Context) error {
	//			fmt.Println("after v1")
	//			return next(c)
	//		}
	//	})
	//}).UseBefore(func(next api.Handler) api.Handler {
	//	return func(c *api.Context) error {
	//		fmt.Println("before api")
	//		return next(c)
	//	}
	//}).UseAfter(func(next api.Handler) api.Handler {
	//	return func(c *api.Context) error {
	//		fmt.Println("after api")
	//		return next(c)
	//	}
	//})
}
