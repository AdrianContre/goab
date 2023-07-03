# goab

There are 4 folders:
Task1 --> Includes all the tests done, for instance 100c represents a request with level 100 of concurrency. Therefore, is the folder called Task1 that explains how I did the tests and my conclusions.


Task2 --> Includes the goab program for apache benchmarks, to execute it type: go run goab -k -c X -n X url -> Where X is the value of the flag and url is where to do the request. Before this, turn on the tomcat server, opening the folder apache-tomcat/bin and execute ./startup(before chmod +x ./startup.sh)



Task3 --> Includes the HttpServer for testing our program, first open a terminal and type: go run serverHttp -> This will launch the server, then, with tomcat shutted down, executing ./shutdown.sh inside apache-tomcat/bin folder, execute goab.go as before and you will see in the terminal of the server what happened.

