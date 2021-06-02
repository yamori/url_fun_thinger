# URL Fun Thinger

URL shortener with a twist.

## running

`make buildrun`

--> [localhost:8080](http://localhost:8080/)

## CSS library

[NES.css](https://nostalgic-css.github.io/NES.css/)

# Things I Learned

- Starting a go module: `go mod init (module name e.g. domain/app_name`)
- Exported identifiers (variable, function, struct key, etc.) are not exported (to another package) if it begins with a lowercase letter.
- Random values, it's necessary to initialize the source within the `rand` module.  `rand.Seed(time.Now().UnixNano())` or read more [here](https://stackoverflow.com/a/39529428)

# EC2 hosting

```
 # build then SCP the binary up to the EC2
make buildaws
scp...
```

```
 # Amazon linux locks down port 80, instead redirect
sudo iptables -t nat -A PREROUTING -p tcp --dport 80 -j REDIRECT --to-ports 8080
```

```
nohup ./url_fun_thinger_aws &
```