FROM golang:1.12.7-stretch as builder

COPY go.mod /go/src/github.com/siangyeh8818/golang.exporter.tencentcloud/go.mod
COPY go.sum /go/src/github.com/siangyeh8818/golang.exporter.tencentcloud/go.sum

# Run golang at any directory, not neccessary $GOROOT, $GOPATH
ENV GO111MODULE=on
WORKDIR /go/src/github.com/siangyeh8818/golang.exporter.tencentcloud

# RUN go mod init github.com/pnetwork/sre.monitor.metrics
RUN go mod download
COPY cmd /go/src/github.com/siangyeh8818/golang.exporter.tencentcloud/cmd
COPY internal /go/src/github.com/siangyeh8818/golang.exporter.tencentcloud/internal
#COPY pkg /go/src/github.com/pnetwork/sre.monitor.metrics/pkg

# Build the Go app
RUN env GOOS=linux GOARCH=amd64 go build -o tencentcloud-exporter -v -ldflags "-s" github.com/siangyeh8818/golang.exporter.tencentcloud/cmd/

##### To reduce the final image size, start a new stage with alpine from scratch #####
FROM alpine:3.9
RUN apk --no-cache add ca-certificates libc6-compat

# Run as root
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/siangyeh8818/golang.exporter.tencentcloud/tencentcloud-exporter /usr/local/bin/tencentcloud-exporter

# EXPOSE 8081

ENTRYPOINT [ "tencentcloud-exporter" ] 
