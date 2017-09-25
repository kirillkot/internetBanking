FROM golang:1.9

# Install beego and the bee dev tool
RUN go get -u \
            github.com/beego/bee \
            github.com/astaxie/beego \
        && mkdir -p src/internetBanking

WORKDIR /go/src/internetBanking

COPY . .
RUN go install internetBanking

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["internetBanking"]
