module github.com/nats-io/nats-server/v2

go 1.16

replace github.com/hack-pad/hackpadfs => github.com/paralin/hackpadfs v0.0.0-20220810055416-aba65e99a16a

require (
	github.com/hack-pad/go-indexeddb v0.1.0
	github.com/hack-pad/hackpadfs v0.1.5
	github.com/klauspost/compress v1.14.4
	github.com/minio/highwayhash v1.0.2
	github.com/nats-io/jwt/v2 v2.2.1-0.20220330180145-442af02fd36a
	github.com/nats-io/nats.go v1.14.0
	github.com/nats-io/nkeys v0.3.0
	github.com/nats-io/nuid v1.0.1
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11
	google.golang.org/protobuf v1.28.0 // indirect
)
