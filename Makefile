ifneq ("$(shell which x86_64-w64-mingw32-gcc)","")
compiler = x86_64-w64-mingw32-gcc
else
compiler = amd64-mingw32msvc-gcc
endif
arch = amd64

# Build the dependencies first (subdirs), then move onto the meat and potatoes.
mimikatz.exe: MemoryModule mimikatz.go
	CC=$(compiler) CGO_ENABLED=1 GOOS=windows GOARCH=$(arch) go build -x mimikatz.go

# Dependency build. 
SUBDIRS = MemoryModule
subdirs: $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@
# Override default subdir build behavior (make) with cmake. 
MemoryModule:
	[ "`ls -A MemoryModule`" ] || git submodule update --init
	$(MAKE) -C $@


# Packing the binary
pack: crypt download
	7z e -so mimikatz_trunk.7z x64/mimikatz.exe | ./crypt > mimi_encrypted

crypt: crypt.go
	go build crypt.go

download:
	curl -OL https://github.com/gentilkiwi/mimikatz/releases/download/2.1.1-20171220/mimikatz_trunk.7z


# Clean target. 
CLEANDIRS = $(SUBDIRS:%=clean-%)
clean: $(CLEANDIRS)
	rm -f crypt mimikatz.exe mimikatz_trunk.7z mimi_encrypted
$(CLEANDIRS): 
	$(MAKE) -C $(@:clean-%=%) clean

test:
	$(MAKE) -C tests test

.PHONY: subdirs $(INSTALLDIRS) $(SUBDIRS) clean test pack download
