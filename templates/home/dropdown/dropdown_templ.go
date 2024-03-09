// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"strings"
)

var Subreddits []string = []string{"memes", "dankmemes", "wholesomememes", "Animemes", "artmemes", "holesome", "2meirl4meirl", "shitposting", "ProgrammerHumor", "PrequelMemes"}
var x_data = fmt.Sprintf(`{options:["%s"], open: false, filter: ''}`, strings.Join(Subreddits, `","`))

// Inspired by https://mistralui.com/components/multiselect/
func Dropdown() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-data=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(x_data))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full relative\"><div @click=\"open = !open\" class=\"text-lg p-2 rounded-lg flex gap-2 w-full border cursor-pointer truncate h-12 bg-black text-[#faebd7]\" x-text=\"&#39;Choose subreddits: &#39; + options.length + &#39; selected&#39;\"></div><div class=\"overflow-y-auto h-[250px] md:bottom-[60px] px-3 pb-3 rounded-lg flex gap-3 w-full shadow-lg x-50 absolute flex-col bg-black mt-3\" x-show=\"open\" x-trap=\"open\" @click.outside=\"open = false\" @keydown.escape.window=\"open = false\" x-transition:enter=\"ease-[cubic-bezier(.3,2.3,.6,1)] duration-200\" x-transition:enter-start=\"!opacity-0 !mt-0\" x-transition:enter-end=\"!opacity-1 !mt-3\" x-transition:leave=\" ease-out duration-200\" x-transition:leave-start=\"!opacity-1 !mt-3\" x-transition:leave-end=\"!opacity-0 !mt-0\"><input x-model=\"filter\" placeholder=\"Search subreddit\" class=\"border-b outline-none p-3 -mx-3 bg-black text-slate-100 sticky top-0 z-10\" type=\"text\"><p x-show=\"! $el.parentNode.innerText.toLowerCase().includes(filter.toLowerCase())\" class=\"text-[#faebd7] text-center font-bold text-2xl\"></p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, subreddit := range Subreddits {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div x-show=\"$el.innerText.toLowerCase().includes(filter.toLowerCase())\" class=\"flex items-center\"><input x-model=\"options\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(subreddit))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" name=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(subreddit))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"checkbox\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(subreddit))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-4 h-4 text-slate-200 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 text-[14px]\"> <label for=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(subreddit))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"ml-2 text-[1rem] font-medium text-white flex-grow tracking-wider\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(subreddit)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/home/dropdown/dropdown.templ`, Line: 23, Col: 114}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</label></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
