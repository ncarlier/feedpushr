module github.com/ncarlier/feedpushr/v3/contrib/twitter

require (
	github.com/ChimeraCoder/anaconda v2.0.0+incompatible
	github.com/ncarlier/feedpushr/v3 v3.0.0-00010101000000-000000000000
)

require (
	github.com/ChimeraCoder/tokenbucket v0.0.0-20131201223612-c5a927568de7 // indirect
	github.com/azr/backoff v0.0.0-20160115115103-53511d3c7330 // indirect
	github.com/dustin/go-jsonpointer v0.0.0-20160814072949-ba0abeacc3dc // indirect
	github.com/dustin/gojson v0.0.0-20160307161227-2e71ec9dd5ad // indirect
	github.com/garyburd/go-oauth v0.0.0-20180319155456-bca2e7f09a17 // indirect
	golang.org/x/net v0.4.0 // indirect
)

replace github.com/ncarlier/feedpushr/v3 => ../..

go 1.17
