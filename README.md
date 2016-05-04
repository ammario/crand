# Crand
A library that proxies crypto/rand with diverse simple functions.

Crand does not tolerate rand.Reader errors and will block until enough entropy is available.

Values returned exclude the maximum.

[Godoc](https://godoc.org/github.com/ammario/crand)

## Disclaimer
I have no formal education in cryptography.
The randomness is derived from crypto/rand and I've implemented safe guards against modulus bias.
