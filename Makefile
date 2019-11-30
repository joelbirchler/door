HOST=linda

all: build certs install run

build:
	GOARM=6 GOARCH=arm GOOS=linux go build -o ./bin/door

certs:
	mkdir -p certs
	openssl req -x509 -nodes \
		-days 365 \
		-newkey rsa:2048 \
		-subj "/C=US/ST=Oregon/L=Eugene/CN=${HOST}" \
		-keyout certs/tls-key.pem \
		-out certs/tls-cert.pem

install: upload-bin upload-static

upload-bin:
	scp ./bin/door pi@${HOST}:~/bin

upload-static:
	rsync -azP ./certs pi@${HOST}:~/
	rsync -azP ./static pi@${HOST}:~/

run:
	ssh -t pi@${HOST} "sudo ~/bin/door"
