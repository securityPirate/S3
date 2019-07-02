# S4
simple secure secret storage

# idea
- S4 store secrets encrypted at rest.
- Secrets transmited securly over TLS.
- Secrets stored as object consist of key , secret and policy.
- Elliptic Curve Cryptography. https://tools.ietf.org/html/rfc5903
- key management ( creation , rotation , deletion ).
- main application generate client lib with the ECC keys.