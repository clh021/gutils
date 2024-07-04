# Exec

## GetProgramPath

```bash
# lee @ lianghongpc in ~/Projects/gutils/examples/ufunc/exec on git:master x [13:09:05]
$ go run main.go
GetProgramPath: /tmp/go-build2945755718/b001/exe

# lee @ lianghongpc in ~/Projects/gutils/examples/ufunc/exec on git:master x [13:09:12]
$ go build main.go

# lee @ lianghongpc in ~/Projects/gutils/examples/ufunc/exec on git:master x [13:09:29]
$ ls -lah
总计 2.0M
drwxr-xr-x 2 lee lee 4.0K Jul 4日 13:09 ./
drwxr-xr-x 3 lee lee 4.0K Jul 4日 12:49 ../
-rwxr-xr-x 1 lee lee 2.0M Jul 4日 13:09 main*
-rw-r--r-- 1 lee lee  152 Jul 4日 12:51 main.go

# lee @ lianghongpc in ~/Projects/gutils/examples/ufunc/exec on git:master x [13:09:32]
$ ./main
GetProgramPath: /home/lee/Projects/gutils/examples/ufunc/exec
```