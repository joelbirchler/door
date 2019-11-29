HOST=linda

all: build install run

build:
	GOARM=7 GOARCH=arm GOOS=linux go build

install:
	scp ./door pi@${HOST}:~/bin

run:
	ssh -t pi@${HOST} "sudo ~/bin/door"
