FROM golang:1.20-alpine as builder

# The cgo tool is enabled by default for native builds on systems where it is expected to work.
# It is disabled by default when cross-compiling.
ENV CGO_ENABLED 0

WORKDIR /build
COPY . /build

# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /build/app main.go

FROM alpine:3

# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single

WORKDIR /build
COPY --from=builder /build/app .

CMD ["./app"]
