module github.com/thegogod/fmq/server

go 1.23.1

require (
	github.com/go-chi/chi/v5 v5.1.0
	github.com/go-chi/render v1.0.3
	github.com/thegogod/fmq/async v0.0.0
	github.com/thegogod/fmq/common v0.0.0
	github.com/thegogod/fmq/logger v0.0.0
	github.com/thegogod/fmq/plugins/mqtt v0.0.0
)

require github.com/ajg/form v1.5.1 // indirect

replace github.com/thegogod/fmq/logger => ../logger

replace github.com/thegogod/fmq/common => ../common

replace github.com/thegogod/fmq/plugins/mqtt => ../plugins/mqtt

replace github.com/thegogod/fmq/async => ../async
