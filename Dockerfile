FROM golang:1.11-alpine
ADD . /go/src/example
RUN go install example

FROM alpine:latest
COPY --from=0 /go/bin/example .
ENV PORT 8080
CMD ["./example"]