module github.com/ncarlier/feedpushr/v3/contrib/wallabag

require (
	github.com/ncarlier/feedpushr/v3 v3.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20211201190559-0a0e4e1bb54c // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/ncarlier/feedpushr/v3 => ../..

go 1.17
