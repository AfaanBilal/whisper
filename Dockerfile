#
# Whisper
#
# A micro-blogging platform.
#
# @author      Afaan Bilal
# @copyright   2023 Afaan Bilal
# @link        https://afaan.dev
#

#
# Stage 1 (Build)
#

FROM golang:1.20-alpine3.17 AS build

WORKDIR /home/go/app

COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o main .

#
# Stage 2 (Run)
#

FROM alpine:3.17

WORKDIR /home/go/app

COPY --from=build /home/go/app/main ./main

EXPOSE 80

# And away we go...
CMD [ "./main" ]
