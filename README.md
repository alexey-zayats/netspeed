# netspeed

Small Golang library that tests the download and upload speeds by using 
- Ookla's https://www.speedtest.net/ 
- Netflix's https://fast.com/.

For the Ookla's measurements used `github.com/showwin/speedtest-go` library as dependency.

For the Netflix's measurements used implementation from `https://github.com/adhocore/fast/blob/main/internal/fast/fast.go`
This is means that the fast.com usage require installed Chrome on target computer  

## Usage example
```go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alexey-zayats/netspeed"
)

func main() {

	var pkind string

	flag.StringVar(&pkind, "provider", "", "Choose a provider to measure net speed with: ookla, netflix")
	flag.Parse()

	result, err := netspeed.Measure(netspeed.KindFromString(pkind))
	if err != nil {
		fmt.Printf("unable measure network speed with provider %s: %s\n", pkind, err)
		os.Exit(1)
	}

	fmt.Printf("Provider: %s\n", result.Provider())
	fmt.Printf("Download: %f Mbps\n", result.Download())
	fmt.Printf("Upload: %f Mbps\n", result.Upload())
	fmt.Printf("Time: %s\n", result.Latency().String())
}

```

Test
```
make test
```

Test coverage
```
make test-cover
```

## Dockerfile example
```dockerfile
FROM golang:alpine AS builder

LABEL stage=netspeed-intermediate

RUN apk -U --no-cache add git make

WORKDIR /go/src/github.com/alexey-zayats/netspeed
COPY ../.. .

RUN go build -o /go/src/github.com/alexey-zayats/netspeed/bin/netspeed /go/src/github.com/alexey-zayats/netspeed/cmd/netspeed

FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /go/src/github.com/alexey-zayats/netspeed/bin/netspeed /app/netspeed

RUN apk -U --no-cache add bash ca-certificates chromium \
    && chmod +x /app/netspeed

ENTRYPOINT ["/app/netspeed"]

```

## Existing docker image
### netflix
```shell
docker run aazayats/netspeed -provider netflix

Provider: netflix
Download: 92.000000 Mbps
Upload: 87.000000 Mbps
Time: 40.873942982s
```
### ookla
```shell
docker run aazayats/netspeed -provider ookla

Provider: ookla
Download: 83.244657 Mbps
Upload: 87.105972 Mbps
Time: 21.497108358s
```

## Build options
```
make help
```
