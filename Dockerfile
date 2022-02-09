FROM golang as build

ENV GOPROXY=https://goproxy.cn

ADD . /camp-course-selection

WORKDIR /camp-course-selection

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7

EXPOSE 3000
MAINTAINER huangzc 540955198@qq.com
VOLUME /tmp

ENV REDIS_ADDR="180.184.74.124:6379"
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN="root:bytedancecamp@tcp(180.184.74.124:3306)/camp_base?charset=utf8mb4&parseTime=True&loc=Local"
ENV GIN_MODE="release"

COPY --from=build /camp-course-selection/api_server /usr/bin/api_server

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]