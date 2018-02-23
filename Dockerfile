FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD corald,.env /
EXPOSE 8080
CMD ["/corald"]

