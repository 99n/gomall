namespace go hello.world

service HelloService {
    string Hello(1: string name) (api.get="/hello");
    string Bye(1: string name) (api.get="/bye");
}
