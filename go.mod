module github.com/sean9999/go-github

replace github.com/google/go-github/v55 => ./github

require (
	github.com/ProtonMail/go-crypto v0.0.0-20230828082145-3c4c8a2d2371
	github.com/google/go-cmp v0.5.9
	github.com/google/go-querystring v1.1.0
)

require (
	github.com/cloudflare/circl v1.3.3 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
)

go 1.21.1
