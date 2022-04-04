FROM golang:1.16-alpine

WORKDIR /go/app

# go get: gopkg.in/urfave/cli.v2@none updating to
#       gopkg.in/urfave/cli.v2@v2.4.0: parsing go.mod:
#        module declares its path as: github.com/urfave/cli/v2#
#                but was required as: gopkg.in/urfave/cli.v2
# のエラーが発生するため「GO111MODULE=off」を利用
RUN apk add --no-cache git \
&& GO111MODULE=off go get -v -u github.com/oxequa/realize

COPY ../../go.mod ../../go.sum ./
RUN go mod download


CMD ["realize", "start"]

