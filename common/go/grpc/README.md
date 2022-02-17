# GRPC

## Usage

Create a new default server with `grpc.NewServer`. Then run it with `grpc.StartServer`.

## Caveats

! Important note about usage of middlewares:
When you chain middlewares with `grpc_middleware.ChainUnaryServer`, they will be chained such as `grpc_middleware.ChainUnaryServer(one, two, three)` will be performed in the following order: one, two, three, endpoint. This means that they will be wrapped in such a way that the final call will be equivalent to:

```go
three(two(one(handler())))
```

In case your middlewares are of the shape:

```go
func UnaryServerInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
 resp, err := handler(ctx, req)

  // Print the number of the middleware (eg. one, two, or three)

 return resp, err
}
```

then your output will look like:

```plaintext
three
two
one
```
