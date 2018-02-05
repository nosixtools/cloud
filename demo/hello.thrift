namespace go hello.rpc

service HelloService {
    string hello(1:string name),
}