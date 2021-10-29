# Exercise 3 - Writing a port scanner 

In this exercise we’ll write a TCP port scanner. If we open up a TCP session by means of a TCP handshake, we know that some process has binded to the port and is open to receiving traffic. In other words: the port is open. The nmap project offers a machine which can legally be scanned. This is the host scanme.nmap.org.  

There are generally two methods to achieve this goal: 

- Easy: write a for loop for the lowest 2048 ports and setup a connection with each of them. Then print out all ports which are considered open. This method is slow as we’ll have to wait until the TCP session is established or timed out. You will probably need the “net” and “fmt” package to get this done.  

- Hard: use the power of goroutines and channels to build a program which executes concurrently. You will probably want so set up a worker function which tests whether a specific port is open. From the main function, you can collect the results and print them to the screen.
