FROM golang:latest 
ENV PORT  8080
ENV GIN_MODE release

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

RUN go get github.com/tools/godep
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/static
RUN go install github.com/tools/godep
RUN go install github.com/gin-gonic/gin
RUN go install github.com/gin-contrib/static

RUN go build -o main . 
EXPOSE 8080
CMD ["/app/main"]