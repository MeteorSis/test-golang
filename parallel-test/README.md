# parallel-test

## func parallel=off, case parallel=off, -p 1

```bash
$ FUNC_PARALLEL=off CASE_PARALLEL=off \time -h /usr/local/go/bin/go test -p 1 -count 1 ./...
?       github.com/MeteorSis/test-golang/parallel-test  [no test files]
ok      github.com/MeteorSis/test-golang/parallel-test/packa    20.597s
ok      github.com/MeteorSis/test-golang/parallel-test/packb    40.408s
        1m1.85s real            0.37s user              0.27s sys
```

## func parallel=off, case parallel=off, -p 40

```bash
$ FUNC_PARALLEL=off CASE_PARALLEL=off \time -h /usr/local/go/bin/go test -p 40 -count 1 ./...
?       github.com/MeteorSis/test-golang/parallel-test  [no test files]
ok      github.com/MeteorSis/test-golang/parallel-test/packa    20.384s
ok      github.com/MeteorSis/test-golang/parallel-test/packb    40.751s
        41.29s real             0.52s user              0.47s sys
```

## func parallel=off, case parallel=on, -p 1

```bash
$ FUNC_PARALLEL=off CASE_PARALLEL=on \time -h /usr/local/go/bin/go test -p 1 -count 1 ./...
?       github.com/MeteorSis/test-golang/parallel-test  [no test files]
ok      github.com/MeteorSis/test-golang/parallel-test/packa    11.189s
ok      github.com/MeteorSis/test-golang/parallel-test/packb    22.169s
        34.15s real             0.36s user              0.26s sys
```

## func parallel=off, case parallel=on, -p 40

```bash
$ FUNC_PARALLEL=off CASE_PARALLEL=on \time -h /usr/local/go/bin/go test -p 40 -count 1 ./...
?       github.com/MeteorSis/test-golang/parallel-test  [no test files]
ok      github.com/MeteorSis/test-golang/parallel-test/packa    11.503s
ok      github.com/MeteorSis/test-golang/parallel-test/packb    22.332s
        22.74s real             0.48s user              0.44s sys
```

## func parallel=on, case parallel=on, -p 40

```bash
$ FUNC_PARALLEL=on CASE_PARALLEL=on \time -h /usr/local/go/bin/go test -p 40 -count 1 ./...
?       github.com/MeteorSis/test-golang/parallel-test  [no test files]
ok      github.com/MeteorSis/test-golang/parallel-test/packa    12.102s
ok      github.com/MeteorSis/test-golang/parallel-test/packb    12.724s
        13.23s real             0.49s user              0.44s sys
```

## ??????

- -p ???????????? ????????? ???(CPU ??????)??? ???????????? `t.Parallel()` ?????? ????????? ??? ????????? ??????????????? ????????? ????????????.
  - ?????? ???????????? ????????? ???????????? ????????? ???????????? ????????? ????????? ???????????? `t.Parallel()` ????????? ????????????.
- ??? ????????? ??????????????? ????????? ???????????? ???????????? `t.Run()`?????? ???????????? ????????? ????????? ??????????????? `t.Parallel()`??? ???????????? ??????.

> **Note**
>
> [Be Careful with Table Driven Tests and t.Parallel()](https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721)
