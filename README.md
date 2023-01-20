# eazy-go

The code is a simple example of connecting to a TCP server using the net package. The main function is the entry point of the program and contains the logic for connecting to a TCP server.

The code first checks if the number of command line arguments passed to the program is not equal to 2 (i.e. the program name and the host:port of the server), if so, it prints an error message and exits the program.

Then it assigns the host:port passed as the second command line argument to the variable 'service'.

Then it uses the net.ResolveTCPAddr function to parse the address passed as a string to a TCPAddr struct. If there is an error, the program will print the error message and exit.

Then it uses the net.DialTCP function to connect to the TCP server. If there is an error, the program will print the error message and exit.

If the connection is successful, it prints a message saying "Connected to" followed by the host:port it connected to. Finally, it closes the connection.
