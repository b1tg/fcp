# fcp
a  file transfer


# usage

Scene: computer A(192.168.1.2) send handbook.pdf to computer B(192.168.1.3)

On computer A: ./fcp s handbook.pdf

On computer B: ./fcp r 192.168.1.2:1234 handbook-copy.pdf

# build & run

```
git clone https://github.com/b1tg/fcp.git
cd fcp
go build

# build for other platform

GOOS=linux GOARCH=amd64 go build 
GOOS=windows GOARCH=amd64 go build
GOOS=darwin GOARCH=amd64 go build

```
