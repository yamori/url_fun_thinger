buildrun:
	go build -o bin/url_fun_thinger url_fun_thinger.go
	./bin/url_fun_thinger

clean:
	rm bin/*