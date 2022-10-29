module github.com/ncarlier/feedpushr/v3/contrib/kafka

require (
	github.com/ncarlier/feedpushr/v3 v3.0.0-00010101000000-000000000000
	github.com/segmentio/kafka-go v0.4.35
)

require (
	github.com/klauspost/compress v1.15.7 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
)

replace github.com/ncarlier/feedpushr/v3 => ../..

go 1.17
