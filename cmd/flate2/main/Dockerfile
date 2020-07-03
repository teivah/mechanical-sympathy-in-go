FROM ubuntu:18.04
RUN apt update && apt install -y linux-tools-common && apt install -y linux-tools-generic
COPY flate2 main
COPY Isaac.Newton-Opticks.txt newton.txt
CMD ["sleep","3600000"]
# env GOOS=linux GOARCH=amd64 go build .
# docker run -d --cap-add SYS_ADMIN flate2
# /usr/lib/linux-tools-4.15.0-109/perf stat -B -e cache-references,cache-misses,cycles,instructions,branches,faults,migrations ./main newton.txt 10000