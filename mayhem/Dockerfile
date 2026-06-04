FROM golang:1.25 as builder

COPY . /spectrum
WORKDIR /spectrum/mayhem

ENV GOTOOLCHAIN=auto
RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN apt update && apt install -y clang

# Create go.mod for the fuzz package with proper dependencies
ENV GOFLAGS=-mod=mod
RUN go mod init fuzzOpenapi2Module && \
    go mod edit -require github.com/grokify/spectrum@v0.0.0 && \
    go mod edit -replace github.com/grokify/spectrum=../ && \
    go get github.com/dvyukov/go-fuzz/go-fuzz-dep && \
    go mod tidy

RUN go-fuzz-build -libfuzzer -o fuzzOpenapi2.a . && \
    clang -fsanitize=fuzzer fuzzOpenapi2.a -o fuzzOpenapi2.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /spectrum/mayhem/fuzzOpenapi2.libfuzzer /
