# Go Clean Architecture

Here I structure my Go project like I usually did in NodeJS. 
Maybe not really optimal, because in some part there still some repetitif code. Also one component that I miss here due to deadline is ClassMapper, I usually use mapper to convert between DTO -> entity -> DPO

The problem with this setup is I can't run the code by executing `go run main.go` but if I compiled the code and execute the binary file, it's run as expected

With this architecture we can easily swap database, create mock (altough i still didnt implement it)

Cheers