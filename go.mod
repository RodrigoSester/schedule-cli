module example/hello

go 1.22.5

replace schedule.com/schedule => ./schedule

require schedule.com/schedule v0.0.0-00010101000000-000000000000

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
