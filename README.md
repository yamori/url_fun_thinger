# URL Fun Thinger

URL shortener with a twist.

## running

`make buildrun`

--> [localhost:8080](http://localhost:8080/)

# Things I Learned

- Starting a go module: `go mod init (module name e.g. domain/app_name`)
- Exported identifiers (variable, function, struct key, etc.) are not exported (to another package) if it begins with a lowercase letter.
- Random values, it's necessary to initialize the source within the `rand` module.  `rand.Seed(time.Now().UnixNano())` or read more [here](https://stackoverflow.com/a/39529428)