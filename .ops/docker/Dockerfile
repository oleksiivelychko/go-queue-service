FROM golang:1.17 as builder
COPY main.go .
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
# -gcflags="all=-N -l"
ARG SKAFFOLD_GO_GCFLAGS
# The cgo tool is enabled by default for native builds on systems where it is expected to work.
# It is disabled by default when cross-compiling.
RUN CGO_ENABLED=0 go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /app main.go

FROM alpine:3
# Define GOTRACEBACK to mark this container as using the Go language runtime
# for `skaffold debug` (https://skaffold.dev/docs/workflows/debug/).
ENV GOTRACEBACK=single
COPY --from=builder /app .
CMD ["./app"]
