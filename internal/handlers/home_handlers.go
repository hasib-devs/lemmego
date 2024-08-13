package handlers

import (
	"github.com/lemmego/lemmego/api"
	"github.com/lemmego/lemmego/templates"
)

func IndexHomeHandler(ctx *api.Context) error {
	return ctx.Inertia(200, "Home/Welcome", map[string]any{
		"name": "John Doe",
	})
	return ctx.HTML(200, `
		<h1>Test Form:</h1>
		<form enctype="multipart/form-data" action="/test" method="POST">
			<input type="text" name="username" placeholder="Username" />
			<input type="password" name="password" placeholder="Password" />
			<input type="submit" value="Submit" />
		</form>
	`)
	// authUser := ctx.Get("user").(*auth.AuthUser)
	return ctx.Templ(templates.BaseLayout(templates.Hello("John Doe")))
	// return ctx.Render(200, "home.page.tmpl", &fluent.TemplateData{
	// 	StringMap: map[string]string{
	// 		"user": authUser.Username,
	// 	},
	// })
}

func StoreTestHandler(ctx *api.Context) error {
	input := &TestInput{}
	if err := ctx.Validate(input); err != nil {
		return err
	}

	return ctx.JSON(200, api.M{"input": input})
}
