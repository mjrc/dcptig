# Exercise 4 - Writing a Fuzzer

In this exercise we’re going to write a fuzzer for a vulnerable FTP server. The goal of a fuzzer is to randomly test input variables and see if the program handles these correctly. In this exercise, the vulnerability has to do with a stack overflow buffer on the variable which handles the login field. You know that you are successful when the program crashes, this means that we have interfered with the regular execution process by overwriting certain pointers.   

Make sure you download the .ova image and run it in either VirtualBox or VMware. There is a desktop icon which opens the FTP server. Copy the ip address to your host machine.  Make sure that the windows box is reachable from your testing box. 

There are two ways to complete this: 

- Hard: start from scratch. you will need the "bufio" and the “net” package. Also, familiarize yourself with the Reader and Writer interface.  
- Easy: Use the following Gist as a starting point. This Gist contains the incomplete code which can be modified. https://gist.github.com/dgiampouris/f5e600b8600d28b44bb3f86a49e6e370. 
