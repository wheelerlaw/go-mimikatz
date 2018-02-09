# go-mimikatz
A Go wrapper around a pre-compiled version of the Mimikatz executable for the purpose of anti-virus evasion.

### Requirements:
	MemoryModule => https://github.com/fancycode/MemoryModule

This application utilizes encryption to encrypt the main mimikatz binary. 

### Build Process:

1. change the encryption key on both go files

2. pinpoint the SIZE constant:

```
make
$ du -bs mimikatz.exe
2637907	mimikatz.exe
```

3. encrypt the real mimikatz binary

```
./crypt mimi32.exe
```

4. build with the correct constants

```
make
```

5. put the encrypted mimikatz at the end of the unpacker

```
cat mimi32.exe >> mimikatz.exe
```

6. deploy the mimikatz.exe 

