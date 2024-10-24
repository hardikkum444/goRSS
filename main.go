package main

import(
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/go-chi/chi"
    "net/http"
)

func main(){

    godotenv.Load(".env")

    portString := os.Getenv("PORT")

    router := chi.NewRouter()

    server := &http.Server{
        Handler: router,
        Addr: ":"+portString,
    }

    fmt.Println("server will now run on -> " + portString)


    err := server.ListenAndServe()

    if err != nil{
        log.Fatal(err.Error())
    }

}
