ifneq ("$(shell which i686-w64-mingw32-gcc)","")
compiler = i686-w64-mingw32-gcc
else
compiler = i586-mingw32msvc-gcc
endif

# Build the dependencies first (subdirs), then move onto the meat and potatoes.
all: subdirs
	CC=$(compiler) CGO_ENABLED=1 GOOS=windows GOARCH=386 go build mimikatz.go
	go build crypt.go

# Dependency build. 
SUBDIRS = MemoryModule
subdirs: $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@

# Clean targed. 
CLEANDIRS = $(SUBDIRS:%=clean-%)
clean: $(CLEANDIRS)
	rm -f crypt mimikatz.exe
$(CLEANDIRS): 
	$(MAKE) -C $(@:clean-%=%) clean

test:
	$(MAKE) -C tests test

test2:
	CGO_ENABLED=0 go build mimikatz.go

.PHONY: subdirs $(INSTALLDIRS) $(SUBDIRS) clean test
