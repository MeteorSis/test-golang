# gRPC-interceptor

## Server

```bash
$ go run main.go server 8080
22:22:31.116586 server interceptor 1_before
22:22:31.116892 server interceptor 2_before
22:22:31.116898 server interceptor 3_before
22:22:31.116905 server handler
22:22:31.116911 server interceptor 3_after
22:22:31.116916 server interceptor 2_after
22:22:31.116922 server interceptor 1_after
```

## Client

```bash
$ go run main.go client 8080
22:22:31.112799 client interceptor 1_before
22:22:31.113130 client interceptor 2_before
22:22:31.113137 client interceptor 3_before
22:22:31.117141 client interceptor 3_after
22:22:31.117160 client interceptor 2_after
22:22:31.117168 client interceptor 1_after
22:22:31.117178 client response
```

## Conclusion

Order

  1. client interceptor 1_before
  2. client interceptor 2_before
  3. client interceptor 3_before
  4. server interceptor 1_before
  5. server interceptor 2_before
  6. server interceptor 3_before
  7. server handler
  8. server interceptor 3_after
  9. server interceptor 2_after
  10. server interceptor 1_after
  11. client interceptor 3_after
  12. client interceptor 2_after
  13. client interceptor 1_after
  14. client response
