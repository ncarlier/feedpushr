module github.com/ncarlier/feedpushr/v3/contrib/kafka

require (
	github.com/ncarlier/feedpushr/v3 v3.0.0-00010101000000-000000000000
	github.com/segmentio/kafka-go v0.3.6
)

replace github.com/ncarlier/feedpushr/v3 => ../..

go 1.13
