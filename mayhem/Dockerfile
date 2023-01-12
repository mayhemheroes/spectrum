FROM golang:1.18 as builder

COPY . /spectrum
WORKDIR /spectrum/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

ENV GOFLAGS -mod=mod

RUN go-fuzz-build -libfuzzer -o fuzzOpenapi2.a && \
    clang -fsanitize=fuzzer fuzzOpenapi2.a -o fuzzOpenapi2.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /spectrum/mayhem/fuzzOpenapi2.libfuzzer /