module main.go

replace github.com/config => ./config

replace github.com/server => ./server

replace github.com/router => ./router

go 1.13

require (
	github.com/config v0.0.0-00010101000000-000000000000
	github.com/router v0.0.0-00010101000000-000000000000
	github.com/server v0.0.0-00010101000000-000000000000
	github.com/spf13/viper v1.9.0
)
