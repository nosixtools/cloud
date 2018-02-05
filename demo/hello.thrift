namespace go hello.rpc

service HelloService {
    string Hello(1:string name),
}