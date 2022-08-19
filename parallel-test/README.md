# parallel-test

## parallel=off, -p 1

```bash
$ PARALLEL=off \time -h /usr/local/go/bin/go test -p 1 -count 1 ./...
?       github.com/MeteorSis/test-golang/parallel-test  [no test files]
ok      github.com/MeteorSis/test-golang/parallel-test/packa    20.926s
ok      github.com/MeteorSis/test-golang/parallel-test/packb    20.567s
        42.49s real             0.42s user              0.31s sys
```

## parallel=off, -p 30

```bash
$ PARALLEL=off \time -h /usr/local/go/bin/go test -p 30 -count 1 ./...
?       github.com/MeteorSis/test-golang/parallel-test  [no test files]
ok      github.com/MeteorSis/test-golang/parallel-test/packa    20.825s
ok      github.com/MeteorSis/test-golang/parallel-test/packb    21.227s
        21.80s real             0.50s user              0.47s sys
```

## parallel=on, -p 1

```bash
$ PARALLEL=on \time -h /usr/local/go/bin/go test -p 1 -count 1 ./...
?       github.com/MeteorSis/test-golang/parallel-test  [no test files]
ok      github.com/MeteorSis/test-golang/parallel-test/packa    11.400s
ok      github.com/MeteorSis/test-golang/parallel-test/packb    11.402s
        23.70s real             0.36s user              0.27s sys
```

## parallel=on, -p 30

```bash
$ PARALLEL=on \time -h /usr/local/go/bin/go test -p 30 -count 1 ./...
?       github.com/MeteorSis/test-golang/parallel-test  [no test files]
ok      github.com/MeteorSis/test-golang/parallel-test/packa    11.856s
ok      github.com/MeteorSis/test-golang/parallel-test/packb    11.533s
        12.40s real             0.48s user              0.43s sys
```

## 결론

- -p 옵션으로 적절한 값(CPU 개수)이 주어지면 `t.Parallel()` 호출 없이도 각각의 Test함수들은 병렬로 실행된다.
- `t.Run()`으로 실행되는 각각의 테스트 케이스에서 `t.Parallel()`을 호출하면 각 테스트 케이스들은 병렬로 실행된다.

> **Note**
> [Be Careful with Table Driven Tests and t.Parallel()](https://gist.github.com/posener/92a55c4cd441fc5e5e85f27bca008721)
