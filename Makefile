test:
	go test -v ./...
chap2:
	go run chapter2/chap2.go 
bench_chap2:
	go test -bench=. ./chapter2/popcount