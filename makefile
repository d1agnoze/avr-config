default:build run
build:
	@@echo "build starting"
	@@cd ./cmd/main; go build main.go
	@@echo "build success"
run:
	@@echo "starting app...."
	@@echo ----------------
	@@cd ./cmd/main; ./main
clean:
	@@cd ./cmd/main;rm main