build:
	ruby bin/build.rb calendar.csv | gofmt > source.go

test:
	go test -race -cover ./...
