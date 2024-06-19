// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Login() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex min-h-full flex-col justify-center px-6 py-12 lg:px-8\"><div class=\"sm:mx-auto sm:w-full sm:max-w-sm\"><img class=\"mx-auto h-10 w-auto\" src=\"https://tailwindui.com/img/logos/mark.svg?color=indigo&amp;shade=600\" alt=\"Your Company\"><h2 class=\"mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900\">Sign in to your account</h2></div><div class=\"mt-10 sm:mx-auto sm:w-full sm:max-w-sm\"><form class=\"space-y-6\" action=\"/login\" method=\"POST\"><div><label for=\"email\" class=\"label-primary\">Email address</label><div class=\"mt-2\"><input id=\"email\" name=\"email\" type=\"email\" autocomplete=\"email\" required class=\"input\"></div></div><div><div class=\"flex items-center justify-between\"><label for=\"password\" class=\"label-primary\">Password</label><div class=\"text-sm\"><a href=\"/forgot-password\" class=\"link-primary\">Forgot password?</a></div></div><div class=\"mt-2\"><input id=\"password\" name=\"password\" type=\"password\" autocomplete=\"current-password\" required class=\"input\"></div></div><div><button type=\"submit\" class=\"btn-primary\">Sign in</button></div></form><p class=\"mt-10 text-center text-sm text-gray-500\">Don't have an account? <a href=\"/register\" class=\"link-primary\">Register now.</a></p></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}