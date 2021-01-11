# gm

This is a trivial gm library which support both CBC and ECB with padding.

## features

provide:

- [x] sm3
- [x] sm4

support:

- [x] CBC
- [x] ECB

## tools

```shell
$go run tools/main.go -h
Usage of /tmp/go-build806812568/b001/exe/main:
  -cipher string
        cipher text in hex format, e.g. '1234' means []byte{0x12, 0x34}
  -f1 string
        factor1 for sm3 to generate protect Key which used by sm4 ecb later
  -f2 string
        factor2 for sm3 to generate protect Key which used by sm4 ecb later
  -plain string
        plain text in normal string, e.g. 'hello world'
exit status 2
$
```

```shell
# encryption
$go run tools/main.go -f1 '12345' -f2 '54321' -plain 'hello world'
[SM4 ECB encrypt] hello world --> 24653acc38e317c8379846a2cf0c32f9

# decryption
$go run tools/main.go -f1 '12345' -f2 '54321' -cipher '24653acc38e317c8379846a2cf0c32f9'
[SM4 ECB decrypt] 24653acc38e317c8379846a2cf0c32f9 --> hello world

# confirm whether plaintext and ciphertext is match under f1 and f2
$go run tools/main.go -f1 '12345' -f2 '54321' -plain 'hello world' -cipher '24653acc38e317c8379846a2cf0c32f9'
[Compare]: hello world <--> 24653acc38e317c8379846a2cf0c32f9  match
```

both f1 and f2 can be omitted.

```shell
$go run tools/main.go -plain 'hello world'
[SM4 ECB encrypt] hello world --> 327afb512955ce72719b5fe9f361d528

$go run tools/main.go -cipher '327afb512955ce72719b5fe9f361d528'
[SM4 ECB decrypt] 327afb512955ce72719b5fe9f361d528 --> hello world

$go run tools/main.go -plain 'hello world' -cipher '327afb512955ce72719b5fe9f361d528'
[Compare]: hello world <--> 327afb512955ce72719b5fe9f361d528  match
```
