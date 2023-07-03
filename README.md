# goab
Tasks done by Adrián Contreras Martín


There are 4 folders:


Task1 --> Includes all the tests done, for instance 100c represents a request with level 100 of concurrency. In addition, there is a file called Task1 where my conclusions are explained and how I did the tests.


Task2 --> Includes the goab program for apache benchmarks, to execute it type: go run goab -k -c X -n X url -> Where X is the value of the flag and url is where to do the request. Before this, turn on the tomcat server, opening the folder apache-tomcat/bin and execute ./startup(before chmod +x ./startup.sh)



Task3 --> Includes the HttpServer for testing our program, first open a terminal and type: go run serverHttp -> This will launch the server, then, with tomcat shutted down, executing ./shutdown.sh inside apache-tomcat/bin folder, execute goab.go as before and you will see in the terminal of the server what happened.

apache-tomcat --> Is the server that I used for the first tests. Inside the folder, there is the folder bin that has two important shell scripts:



        -- startup.sh and shutdown.sh, the first is required to turn on the server and the second is needed to turn down the server.

