package main
import (
        //"fmt"
        "login-system/funcs"
        //"html/template"
        "github.com/zenazn/goji"
        //"github.com/zenazn/goji/web"
        "flag"

       )

func main() {

    goji.Get("/reg", funcs.Registration)
    goji.Get("/login", funcs.Login)
    goji.Post("/display", funcs.Display)
    goji.Get("/welcome", funcs.Welcome)
    goji.Get("/slogin", funcs.SessionLogin)
    goji.Get("/slogout", funcs.SessionLogout)
    flag.Set("bind" , ":8080")
    goji.Serve()

}



