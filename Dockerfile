#https://docs.docker.com/language/golang/build-images/
#https://hub.docker.com/_/golang

# Specify go base image
FROM golang:1.19.3-alpine

RUN apk add --update --no-cache gcc libc-dev  
ENV GO111MODULE=on

# Tạo và nhảy vào thư mục working directory, Tất cả lệnh RUN, CMD, ADD, COPY, or ENTRYPOINT command sẽ được thực thi tại thư mục này. Nếu chưa có thì thư mục tự dc tạo
WORKDIR /app

# Copy file tại thư mục ngang hàng với Dockerfile vào thư mục app (.)
COPY go.mod .
COPY go.sum .
# RUN go mod download && go mod verify

#Copy tất cả file ở thư mục của project vào trong thư mục app
COPY . .

#expose port 8080 (default is 80)
EXPOSE 8080

# compile our go app and output it to app với binary file tên là app
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# start the application (chạy đúng cái app binary vừa build, dùng lệnh ./ chạy bình thường như linux command)
CMD ["./server"]

# Giờ chỉ cần chạy lệnh docker run docker run --name mygolang -dit -p 1323:1323 my-golang-app
# Hoặc dùng chạy cùng với docker-compose