default:build run
build:
	@@cd ./cmd/main; go build main.go
	@@echo "build success"
run:
	@@cls;echo "starting app...."
	@@echo ----------------
	@@cd ./cmd/main; ./main.exe
	@@echo ----------------
clean:
	@@cd ./cmd/main;rm main.exe