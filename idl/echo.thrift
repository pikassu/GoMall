namespace go api

struct Requset{
    1:string message
}

struct Response{
    1:string message
}

service Echo{
    Response Echo(1:Requset req)
}