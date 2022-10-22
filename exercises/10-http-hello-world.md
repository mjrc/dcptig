# 10. Hello World Web Server 

- Use the `net/http` package 
- Create `handler(w http.ResponseWriter, req *http.Request)` function, the handler should write `Hello World` to the `http.ResponseWriter` 
- Register an endpoint like `/mick` which is handled by the handler function 
- Run the server by `ListenAndServe` on port `8080`
