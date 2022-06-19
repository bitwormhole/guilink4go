package frontend

// Station 表示前端的一个站点(aka.Client)
type Station struct {
	Host   string // 默认的主机名
	Port   int    // 默认的端口
	Client Client
}
