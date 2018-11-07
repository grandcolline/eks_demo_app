FROM golang:latest as build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/grandcolline/ec2_info_app
COPY . .
RUN go install

CMD ["/go/bin/ec2_info_app"]

## Application Image
#FROM gcr.io/distroless/base
#
#COPY --from=build /go/bin/ec2_info_app /ec2_info_app
#ENV PORT=8080
#EXPOSE 8080
#
#CMD ["/ec2_info_app"]

