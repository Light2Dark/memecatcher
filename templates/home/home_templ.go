// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Light2Dark/memecatcher/templates/home/dropdown"

func Index() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Memes</title><link href=\"templates/css/styles.css\" rel=\"stylesheet\"><link href=\"templates/css/output.css\" rel=\"stylesheet\"><link rel=\"preconnect\" href=\"https://fonts.googleapis.com\"><link rel=\"preconnect\" href=\"https://fonts.gstatic.com\" crossorigin><link href=\"https://fonts.googleapis.com/css2?family=Amatic+SC:wght@400;700&amp;display=swap\" rel=\"stylesheet\"><link rel=\"icon\" type=\"image/x-icon\" href=\"templates/assets/doge.png\"><link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css\"><script src=\"templates/assets/htmx.min.js\"></script><script src=\"templates/assets/loading-states.js\"></script><script defer src=\"templates/assets/alpine.min.js\"></script><!-- Google tag (gtag.js) --><script async src=\"https://www.googletagmanager.com/gtag/js?id=G-9SZVRLJ6X8\"></script><script>\n\t\t\t\twindow.dataLayer = window.dataLayer || [];\n\t\t\t\tfunction gtag(){dataLayer.push(arguments);}\n\t\t\t\tgtag('js', new Date());\n\n\t\t\t\tgtag('config', 'G-9SZVRLJ6X8');\n\t\t\t</script></head><body class=\"amatic-sc-bold bg-gradient-to-r from-violet-200 to-pink-200 w-screen\"><h1 class=\"text-6xl text-slate-900 mt-14 mb-10 text-center\">Find A Meme</h1><div class=\"flex flex-col lg:flex-row gap-12 lg:gap-8 lg:mx-16 justify-center items-center lg:items-start\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = Main().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = RecentMemes().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Main() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-5/6 lg:w-full\"><div class=\"mb-10 flex flex-col gap-4 items-center\"><div class=\"bg-slate-800 h-[300px] w-full lg:h-[500px] rounded-lg shadow-lg flex justify-center items-center\"><svg id=\"spinner\" viewBox=\"0 0 50 50\" class=\"htmx-indicator transition-all\"><circle class=\"path\" cx=\"25\" cy=\"25\" r=\"20\" fill=\"none\" stroke-width=\"5\"></circle></svg><div id=\"imageContainer\" class=\"w-full h-full flex justify-center items-center\"></div></div><p id=\"textContent\" class=\"w-2/3 text-center text-2xl\"></p></div><form id=\"memeForm\" hx-post=\"/fetchMeme\" hx-target=\"#imageContainer\" hx-indicator=\"#spinner\" class=\"flex flex-col gap-4 md:w-[540px] justify-center mx-auto\"><div class=\"flex flex-row gap-4\"><input required type=\"search\" name=\"search\" id=\"search\" placeholder=\"keywords\" class=\"text-2xl p-2 border-2 rounded-lg w-4/6 lg:w-4/5\"> <button type=\"submit\" class=\"border-2 border-black rounded-lg bg-slate-900 text-[#faebd7] px-3 hover:bg-slate-800 transition-all text-2xl w-2/6 lg:w-1/5\">Search</button></div><div class=\"flex flex-col\"><div class=\"flex flex-row gap-4 text-2xl\" x-data=\"{ memeVal : 30 }\"><p>Memes searched</p><input type=\"range\" name=\"numMemes\" id=\"numMemes\" min=\"1\" max=\"50\" x-model=\"memeVal\"><p x-text=\"memeVal\"></p></div><div class=\"flex flex-row gap-2\"><label for=\"nsfw\" class=\"text-xl\">NSFW</label> <input type=\"checkbox\" name=\"nsfw\" id=\"nsfw\" checked></div></div><div class=\"flex flex-col gap-1\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templates.Dropdown().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-row gap-2\"><label for=\"extraSubreddits\" class=\"text-lg\">Extra Subreddits:</label> <input type=\"text\" name=\"extraSubreddits\" id=\"extraSubreddits\" placeholder=\"Separate by comma\" class=\"font-extralight text-xs font-mono rounded-lg bg-black px-2 py-1.5 flex-1 text-white placeholder:text-[#faebd7]\"></div></div></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func RecentMemes() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-5/6\"><div class=\"flex flex-row gap-4 justify-center items-center mb-4\"><h3 class=\"text-2xl\">Recent Memes</h3><button type=\"button\" hx-get=\"/getAllMemes\" hx-target=\"#allMemes\" hx-trigger=\"load,submit from:#memeForm delay:4s,click\"><i class=\"fa fa-refresh refresh\" style=\"font-size:20px\"></i></button></div><div id=\"allMemes\" class=\"grid md:grid-cols-2 gap-2\"><div class=\"h-48 bg-blue-50 hover:border-2 hover:border-lime-50\"></div><div class=\"h-48 bg-red-50\"></div><div class=\"h-48 bg-green-50\"></div><div class=\"h-48 bg-yellow-50\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
