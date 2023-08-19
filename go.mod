module github.com/mmadfox/gpsgend

go 1.21

toolchain go1.21.0

require (
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/fasthttp/websocket v1.5.3
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/lucasb-eyer/go-colorful v1.2.0
	github.com/mmadfox/go-gpsgen v0.2.1-0.20230819121240-143682c9c64a
	github.com/stretchr/testify v1.8.4
	github.com/valyala/fasthttp v1.48.0
	google.golang.org/protobuf v1.30.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/icholy/utm v1.0.1 // indirect
	github.com/klauspost/compress v1.16.6 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/paulmach/orb v0.10.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/savsgio/gotils v0.0.0-20230208104028-c358bd845dee // indirect
	github.com/tkrajina/gpxgo v1.3.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
	go.mongodb.org/mongo-driver v1.11.7 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gotest.tools/v3 v3.0.3 // indirect
)

replace github.com/docker/docker => github.com/docker/docker v20.10.3-0.20221013203545-33ab36d6b304+incompatible // 22.06 branch
