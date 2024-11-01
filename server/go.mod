module github.com/thegogod/fmq/server

go 1.23.1

require (
	github.com/thegogod/fmq/logger v0.0.0
	github.com/thegogod/fmq/common v0.0.0
	github.com/thegogod/fmq/plugins/mqtt v0.0.0
)

replace github.com/thegogod/fmq/logger => ../logger

replace github.com/thegogod/fmq/common => ../common

replace github.com/thegogod/fmq/plugins/mqtt => ../plugins/mqtt
