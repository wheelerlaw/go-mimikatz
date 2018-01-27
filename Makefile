ifneq ("$(shell which i686-w64-mingw32-gcc)","")
compiler=i686-w64-mingw32-gcc
else
compiler=i586-mingw32msvc-gcc
endif

all:
	CC=$(compiler) CGO_ENABLED=1 GOOS=windows GOARCH=386 go build mimikatz.go
	go build crypt.go
test:
	CGO_ENABLED=0 go build mimikatz.go
clean:
	rm crypt mimikatz.exe