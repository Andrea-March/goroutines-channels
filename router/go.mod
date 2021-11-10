module router

replace github.com/config => ../config

go 1.13

require (
	github.com/config v0.0.0-00010101000000-000000000000
	github.com/golang/glog v1.0.0 // indirect
	github.com/julienschmidt/httprouter v1.3.0
	github.com/rs/cors v1.8.0
)
