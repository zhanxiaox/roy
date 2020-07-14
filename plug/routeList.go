package plug

var RouteList map[string][]string

func Setup() {
	RouteList = make(map[string][]string)
	RouteList["home"] = []string{"home", "index"}
}
