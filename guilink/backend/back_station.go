package backend

// Station 表示后端的一个站点(aka.Server)
type Station struct {
	Host   string
	Port   int
	Server Server
}
