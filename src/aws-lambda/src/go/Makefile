build:
	GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

zip:
	zip lambda_function.zip bootstrap

clean:
	rm -f bootstrap lambda_function.zip