https://github.com/carlmjohnson/requests

https://golang.cafe/blog/golang-context-with-timeout-example.html


time.Duration(time.Millisecond*100
res, err := http.DefaultClient.Do(req)

log.Fatalln(err)
    2021/12/28 06:23:35 Get "https://jsonplaceholder.typicode.com": context deadline exceeded
    exit status 1

panic(err)
    panic: Get "https://jsonplaceholder.typicode.com": context deadline exceeded

    goroutine 1 [running]:
    main.main()
        /Users/..../gobasic/web/get.go:28 +0x2bc
    exit status 2

time.Duration(time.Millisecond*300
    no error