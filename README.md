# Crand
A library that proxies crypto/rand with diverse simple functions.


The library does not tolerate errors and will block until a source of entropy can fulfill a client's request.

Values returned exclude the maximum.

[Godoc](https://godoc.org/github.com/ammario/crand)

## Disclaimer
I have no formal education in cryptography.
The randomness is derived from crypto/rand and I've implemented safe guards against modulus bias.
