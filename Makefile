all:
	go build -o music-waybar-module cmd/music-waybar-module/main.go

run:
	go run cmd/music-waybar-module/main.go

tidy:
	go mod tidy
	go mod vendor

clean:
	rm music-waybar-module

