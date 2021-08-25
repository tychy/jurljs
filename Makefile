test:
	go test -v ./...
chap2:
	go run chapter2/chap2.go 
bench_chap2:
	go test -bench=. ./chapter2/popcount
chap4:
	go run chapter4/chap4.go 
# chapter 8
run_clock:
	go build -o bin/clock chapter8/clock/clock.go
	./bin/clock -p 8012 &
	TZ=US/Eastern ./bin/clock -p 8013
	
