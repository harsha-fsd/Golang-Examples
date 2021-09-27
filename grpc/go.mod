module main

replace apiproto => ./protobuf/

go 1.16

require (
	apiproto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.41.0
)
