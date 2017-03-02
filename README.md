# Benami

*Benami in Hindi refers to a transaction made in the name of another person.*

The project does simple [domain fronting](https://www.bamsoftware.com/papers/fronting/) on a
server under your control.

```
$ git clone https://github.com/theju/benami
$ cd benami
$ go build server.go
$ ./server -debug
# On another shell
$ curl -H "Host: duckduckgo.com" "http://localhost:8080/"
```

## Status

The project is still under development.

## Roadmap

- A browser extension (WebExtension) to manipulate the request to redirect to 
the fronting domain transparently
- Load test the fronting server and see how it performs

## License

The project is licensed under the MIT License. Please check the `License.md` file for more details.

## Alternatives

There are more mature alternatives out there:
- [Meek](https://trac.torproject.org/projects/tor/wiki/doc/meek)
- [Psiphon](https://bitbucket.org/psiphon/psiphon-circumvention-system) (does more than domain fronting)
