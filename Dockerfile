FROM golang:1.16

WORKDIR /go/src/github.com/montybeatnik/tutorial

COPY . .

CMD ["go", "run", "."]

EXPOSE 8000:8000


# TO BUILD
# docker build -t tutorial .
# TO RUN
# docker run -dp 8000:8000 tutorial
