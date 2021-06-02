buildrun:
	go build -o bin/url_fun_thinger url_fun_thinger.go
	./bin/url_fun_thinger

buildaws:
	GOOS=linux GOARCH=amd64 go build -o bin/url_fun_thinger_aws url_fun_thinger.go

clean:
	rm bin/*