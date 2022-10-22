# 07. Dial with TCP  

- Use the `net` package to establish a TCP connection
- Fetch the content of your favourite website by writing `GET / HTTP/1.0\r\n\r\n` to the connection object 
- Use the `bufio` package to read from the connection, printing out all lines received 