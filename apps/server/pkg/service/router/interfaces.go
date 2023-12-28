package router

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Monitor
type Monitor interface {
	GetCpuUsage() (float32, error)
	GetMemUsage() (float32, error)
}
