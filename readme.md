
# NET-CAT
* Executing 
     go run . 
     in another terminal nc localhost port number (by default 8989)
     do 2nd step at least 2 times to create several users 

OR
 * docker build -t my_golang_tcp_chat .
  * docker run -p 8000:8000 my_golang_tcp_chat





Accept connection with non-empty name
The client :

    *$ nc $IP $port

Server:

    *$ go run . 2525
    *Listening on the port :2525

Client1 (Yenlik):

    *$ nc localhost 2525
    *Welcome to TCP-Chat!
    *         _nnnn_
    *        dGGGGMMb
    *       @p~qp~~qMb
    *       M|@||@) M|
    *       @,----.JM|
    *      JS^\__/  qKL
    *     dZP        qKRb
    *    dZP          qKKb
    *   fZP            SMMb
    *   HZM            MMMM
    *   FqM            MMMM
    * __| ".        |\dS"qML
    * |    `.       | `' \Zq
    *_)      \.___.,|     .'
    *\____   )MMMMMP|   .'
    *     `-'       `--'
    *[ENTER YOUR NAME]: Yenlik
    *[2020-01-20 16:03:43][Yenlik]:hello
    *[2020-01-20 16:03:46][Yenlik]:How are you?
    *[2020-01-20 16:04:10][Yenlik]:
    *Lee has joined our chat...
    *[2020-01-20 16:04:15][Yenlik]:
    *[2020-01-20 16:04:32][Lee]:Hi everyone!
    *[2020-01-20 16:04:32][Yenlik]:
    *[2020-01-20 16:04:35][Lee]:How are you?
    *[2020-01-20 16:04:35][Yenlik]:great, and you?
    *[2020-01-20 16:04:41][Yenlik]:
    *[2020-01-20 16:04:44][Lee]:good!
    *[2020-01-20 16:04:44][Yenlik]:
    *[2020-01-20 16:04:50][Lee]:alright, see ya!
    *[2020-01-20 16:04:50][Yenlik]:bye-bye!
    *[2020-01-20 16:04:57][Yenlik]:
    *Lee has left our chat...
    *[2020-01-20 16:04:59][Yenlik]:

Client2 (Lee):

    *$ nc localhost 2525
    *Yenliks-MacBook-Air:simpleTCPChat ybokina$ nc localhost 2525
    *Yenliks-MacBook-Air:simpleTCPChat ybokina$ nc localhost 2525
    *Welcome to TCP-Chat!
    *         _nnnn_
    *        dGGGGMMb
    *       @p~qp~~qMb
    *       M|@||@) M|
    *       @,----.JM|
    *      JS^\__/  qKL
    *     dZP        qKRb
    *    dZP          qKKb
    *   fZP            SMMb
    *   HZM            MMMM
    *   FqM            MMMM
    * __| ".        |\dS"qML
    * |    `.       | `' \Zq
    *_)      \.___.,|     .'
    *\____   )MMMMMP|   .'
    *     `-'       `--'
    *[ENTER YOUR NAME]: Lee
    *[2020-01-20 16:03:43][Yenlik]:hello
    *[2020-01-20 16:03:46][Yenlik]:How are you?
    *[2020-01-20 16:04:15][Lee]:Hi everyone!
    *[2020-01-20 16:04:32][Lee]:How are you?
    *[2020-01-20 16:04:35][Lee]:
    *[2020-01-20 16:04:41][Yenlik]:great, and you?
    *[2020-01-20 16:04:41][Lee]:good!
    *[2020-01-20 16:04:44][Lee]:alright, see ya!
    *[2020-01-20 16:04:50][Lee]:
    *[2020-01-20 16:04:57][Yenlik]:bye-bye!
    *[2020-01-20 16:04:57][Lee]:^C
