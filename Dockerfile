FROM golang:1.21-alpine

RUN apk add --no-cache go poetry just
RUN apk add --no-cache gcc musl-dev

ENV CGO_ENABLED=1

WORKDIR /task

COPY cucumber ./cucumber
RUN cd cucumber && poetry install
COPY . .

CMD ["just", "test"]
