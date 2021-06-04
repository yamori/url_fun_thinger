# URL Fun Thinger

URL shortener (into three word string), and fun NES CSS styling.

Persistence is simplistic local JSON file read/write.

## running

`make buildrun`

--> [localhost:8080](http://localhost:8080/)

Also a `make buildaws` for running on an Amazon (ec2) AMI.  Just deploy the binary `bin/url_fun_thinger_aws`!

## CSS library

[NES.css](https://nostalgic-css.github.io/NES.css/)

# Things I Learned

- Starting a go module: `go mod init (module name e.g. domain/app_name`)
- Exported identifiers (variable, function, struct key, etc.) are not exported (to another package) if it begins with a lowercase letter.
- Random values, it's necessary to initialize the source within the `rand` module.  `rand.Seed(time.Now().UnixNano())` or read more [here](https://stackoverflow.com/a/39529428)
- How to get a value from a `map` object: `objmap[keyString].(string)`

## TODO / Fixes

- Lacking a routing framework like `mux`, as an unsuccessful POST to `/lookup` needs an error message to be displayed but the URL ends up as `lookup`.
- The two html template should be abstracted somehow.  But the 'redirect' isn't meant to be displayed anyways.

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