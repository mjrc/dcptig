# 01. Concurrent port scanner

- Use a worker function for determining if a port is open
- Use 100 workers to check for the first 1024 ports on `scanme.nmap.org`
- The main function should run workers as goroutines
- The main function should retreive whether a port is closed or open from channels
