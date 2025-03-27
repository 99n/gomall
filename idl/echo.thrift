namespace go api

struct Request {
    1: string message
}

struct Responce {
    1: string message
}

service Echo {
    Responce echo(1: Request req )
}