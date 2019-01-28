package funcs
import (
        "fmt"
      _ "github.com/lib/pq"
        "database/sql"
        "html/template"
        "net/http"
        "github.com/gorilla/sessions"
        //"os"
        //"github.com/robertkrimen/otto"
        )

type Reg struct{
	Id int64
	FirstName string
	LastName string
	Email string
	Password string
}

type Loginst struct{
	Email string
	Password string
}

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func Dbconn() *sql.DB {
   const (
     host = "localhost"
     port = 5432
     username = "postgres"
     password = "test123"
     dbname = "loginsystem"
       )

     psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s"+
     	" dbname=%s sslmode=disable ",host, port, username, password, dbname)

    db, err := sql.Open("postgres", psqlinfo)
    if err != nil{
      fmt.Println(">>",err)
    }
    fmt.Println("connected")
    return db
}

func Registration(w http.ResponseWriter, r *http.Request) {

    t, _ := template.ParseFiles("reg.html")
    t.Execute(w, "")
}

func Login(w http.ResponseWriter, r *http.Request) {

  // session, err := store.Get(r, "session1")
  // if err != nil {
  // 	panic (err)
  // }
  // session.Values["email"] = r.FormValue("email")
  // session.Values["password"] = r.FormValue("password")
  // session.Save(r,w)
  //fmt.Println(">?>?>?>",session1.Values["email"])
   t, _ := template.ParseFiles("login.html")
   t.Execute(w, "")
}

func Display(w http.ResponseWriter, r *http.Request) {

	  db2 := Dbconn()
    reg := new(Reg)
    reg.FirstName = r.FormValue("first_name")
    reg.LastName = r.FormValue("last_name")
    reg.Email = r.FormValue("email")
    reg.Password = r.FormValue("password")
    //fmt.Println(reg.FirstName)
    flag := true
    var em Reg
    var total []Reg
    sql := "select email from registration"
    rows, _ :=db2.Query(sql)
    for rows.Next(){
    	er := rows.Scan(&em.Email)
    	if er != nil{
    		fmt.Println(er)
    	}
    	total = append(total, em)
    }

    for i := 0; i < len(total); i++ {
    	 if total[i].Email == reg.Email{
    	 	flag = false
    	 }
    }
    if flag == false{
    	fmt.Fprintf(w, "email is duplicated")
    }else{

    insert := `insert into registration(first_name, last_name, email, password) values($1, $2, $3, $4)`
     _, err := db2.Exec(insert, reg.FirstName, reg.LastName, reg.Email, reg.Password)
    if err != nil {
    	fmt.Println(">>>>err", err)
    }
    //stmt.Query(reg.FirstName, reg.LastName, reg.Email, reg.Password)

	t, _ := template.ParseFiles("display.html")
	t.Execute(w, reg)
}
}

func Welcome(w http.ResponseWriter, r *http.Request) {
   session, _ := store.Get(r, "session")
  if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", 302)
	}
  // session, err := store.Get(r, "session1")
  // if err != nil {
  // 	panic (err)
  // }

  // session.Save(r,w)


  // session.Values["email"] = r.FormValue("email")
  // session.Values["password"] = r.FormValue("password")
  db2 := Dbconn()
  ls := new(Loginst)

  // if session.Values["email"] != ""{
    // http.Redirect(w, r, "/login", 302)
  // }
  ls.Email = r.FormValue("email")
  ls.Password = r.FormValue("password")

  // fmt.Println(session.Values["email"],">>>", r.FormValue("email"))
  // fmt.Println(session.Values["email"]== r.FormValue("email"))

  //fmt.Println( "aniket->",session.Values["email"])
  var Email string
  var Password string
  auth := "select email, password from registration where email = $1"
  rows,err:= db2.Query(auth, ls.Email)
  if err != nil {
  	panic (err)
  }
  for rows.Next(){
  	err := rows.Scan(&Email, &Password)
  	if err != nil{
  		panic (err)
  	}
  }

  // fmt.Println(ls.Email == Email && ls.Password == Password)
   //fmt.Println(Email, ls.Email ,  ls.Password , Password)

  if ls.Email == Email && ls.Password == Password {

  t, _ := template.ParseFiles("welcome.html")
  t.Execute(w, ls)
  }else{
     // fmt.Println("abcd")
    http.Redirect(w, r, "/login", 302)
  }

  //fmt.Println("mayank",Email, Password)
  //fmt.Println(">>>>?????",ls.Email, ls.Password)
}

func SessionLogin(w http.ResponseWriter, r *http.Request) {
      session, _ := store.Get(r, "session")
      session.Values["authenticated"] = true
      session.Values["email"] = r.FormValue("email")
      session.Values["password"] = r.FormValue("password")
      session.Save(r, w)
      http.Redirect(w, r, "/welcome", 302)
}

func SessionLogout(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "session")
    session.Values["authenticated"] = false
    session.Save(r, w)
    http.Redirect(w, r, "/login", 302)
}



