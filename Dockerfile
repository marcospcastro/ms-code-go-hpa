  
FROM golang:1.14-alpine as builder

WORKDIR /go/src/hpa
COPY ./src/hpa/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o hpa
RUN ls

FROM scratch
COPY --from=builder  /go/src/hpa .

CMD [ "./hpa" ]