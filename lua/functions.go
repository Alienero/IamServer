package lua

import (
	"html"

	"github.com/yuin/gopher-lua"
)

func initFunctions(L *lua.LState) {
	L.PreloadModule("libs.html", urlLoader)
}

// module: libs.html
func urlLoader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"html_escape": htmlEscape,
	})
	// returns the module
	L.Push(mod)
	return 1
}

func htmlEscape(L *lua.LState) int {
	str := L.ToString(1)
	L.Push(lua.LString(html.EscapeString(str)))
	return 1
}
