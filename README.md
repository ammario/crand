# Crand
A library that proxies crypto/rand with diverse simple functions.

Values returned exclude the maximum.

[Godoc](https://godoc.org/github.com/ammario/crand)


## Example

Generate a secure, random token

```go
fmt.Printf("Token: %v", crand.String(20))
```

## Disclaimer
I have no formal education in cryptography.
The randomness is derived from crypto/rand and I've implemented safeguards against modulus bias.
