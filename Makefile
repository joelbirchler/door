HOST=linda

all: build install run

build:
	GOARM=6 GOARCH=arm GOOS=linux go build -o ./bin/door

install: upload-bin upload-static

upload-bin:
	scp ./bin/door pi@${HOST}:~/bin

upload-static:
	rsync -azP ./static pi@${HOST}:~/

run:
	ssh -t pi@${HOST} "sudo ~/bin/door"
