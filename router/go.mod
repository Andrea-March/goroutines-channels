module router

replace github.com/config => ../config

replace github.com/routines => ../routines

go 1.13

require (
	github.com/golang/glog v1.0.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/routines v0.0.0-00010101000000-000000000000
	github.com/rs/cors v1.8.0
)
