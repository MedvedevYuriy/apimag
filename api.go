package main

import (
        "encoding/json"
        "log"
        "net/http"
        "database/sql"
        "github.com/lib/pq"

        "fmt"
)

var db *sql.DB

type Product struct {
        Id int
        Name string
        Description string
        Balance int
        Discount int
        Category int
}

type User struct {
        Id int
        Email string
        Password string
        Role int
        Status bool
}

type Products struct {
        Products []Product
}

type Users struct {
        Users []User
}

func main () {
        var err error

        db, err = sql.Open("postgess", "host=127.0.0.1 user=api password 123456 dbname=api sslmode=disable")
        if err != nil {
                panic(err)
        }

        defer db.Close()

        fmt.Println("starting server...")

        http.handleFunc("/v1/products/add", addProduct)
        http.handleFunc("/v1/products/", getProducts)
        http.handleFunc("/v1/user", addUser)
        http.handleFunc("/v1/user/search", searchUser)
        log.Fatal(http.ListenAndServe(":8080", nill))
}

func addUser (w http.ResponseWritter, r *http.Request) {
        if r.Method != "POST" {
                http.Error(w, "Method Not Allowed", 405)
        } else {
                decoder := json.NewDecoder(r.Body)
                var g_user User

                err := decoder.Decode(&g_user)
                if err != nil {
                        panic(err)
                }

                query := fmt.Sprintf("INSERT INTO users (email, password, role, status) VALUES ('%s', '%s', %d, %t) RETURNING id", g_user.Email, g_user.Password, g_user.Rol
e, g_user.Status)

                fmtPrintln("# INSERT QUERY: %s", query)

                rows, err := db.Query(query)
                if err != nil {
                        panic(err)
                }

                for rows.Next() {
                        var id int
                        err = rows.Scan(&id)
                        if err != nil {
                                panic(err)
                        }
                        fmt.Fprintf(w, "{\"id\":%d}", id)
                }
        }
}

func searchUser (w http.ResponseWritter, r *http.Request) {
        if r.Method != "GET" {
                http.Error(w, "Method Not Allowed", 405)
        } else {
                decoder := json.NewDecoder(r.Body)
                var g_user User

                err := decoder.Decode(&g_userid)
                if err !=nil {
                        panic(err)
                }

                fmt.Println("# Querying")
                rows, err := db.Query("SELECT from roles where id=%d", g_userid.Id)
                if err != nil {
                        panic(err)
                }

                for rows.Next() {
                        w_user := User{}

                        err = rows.Scan(&w_user.Email,&w_user.Role,&w_user.Status)
                        if err != nil {
                                panic (err)
                        }
                        w_array.Users = append(w_array.Users, w_user)
                }

                json.NewEncoder(w).Encode(w_array)
        }
}


func addProduct (w http.ResponseWritter, r *http.Request) {
        if r.Method != "POST" {
                http.Error(w, "Method Not Allowed", 405)
        } else {
                decoder := json.NewDecoder(r.Body)
                var g_product Product

                err := decoder.Decode(&g_product)
                if err != nil {
                        panic(err)
                }

                query := fmt.Sprintf("INSERT INTO products(name, description, balance, discount, category) VALUES('%s', '%s', %d, %d, %d) RETURNING id", g_product.Name, g_p
roduct.Description, g_product.Balance, g_product.Discount, g_product.Category)

                fmt.Println("# INSERT QUERY: %s", query)

                rows, err := db.Query(query)
                if err != nil {
                        panic(err)
                }

                for rows.Next() {
                        var id int
                        err = rows.Scan(&id)
                        if err != nil {
                                panic(err)
                        }
                        fmt.Fprintf(w, "{\"id\":%d}", id)
                }

        }
}

func getProducts(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
                http.Error(w, "Method Not Allowed", 405)
        } else {
                w_array := Products{}

                fmt.Println("# Querying")
                rows, err := db.Query("SELECT id,name,description,discount,category from products")
                if err != nil {
                        panic(err)
                }

                for rows.Next() {
                        w_product := Product{}

                        err = rows.Scan(&w_product.Id,&w_product.Name,&w_product.Description,&w_product.Discount,&w_product.Category)
                        if err != nil {
                                panic(err)
                        }
                        w_array.Products = append(w_array.Products, w_product)

                }

                json.NewEncoder(w).Encode(w_array)
        }
}
