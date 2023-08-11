package main



import (



    "encoding/json"

    "fmt"

    "net/http"

    "github.com/go-redis/redis"



)    



func main() {

    http.HandleFunc("/scores", httpHandler)        

    http.ListenAndServe(":8080", nil)

}



func httpHandler(w http.ResponseWriter, req *http.Request) {            



    redisClient := redis.NewClient(&redis.Options{

    Addr: "localhost:6379",

    Password: "",

    DB: 0,

    })   



    params := map[string]interface{}{}



    resp := map[string]interface{}{}

    var err error



    if req.Method == "GET" {



        for k, v := range req.URL.Query() {

            params[k] = v[0]

        }    



        resp, err = getScores(redisClient, params)





    } else if req.Method == "POST" {



        err = json.NewDecoder(req.Body).Decode(&params)



        resp, err = addScore(redisClient, params)

    }



    enc := json.NewEncoder(w)

    enc.SetIndent("", "  ")



    if err != nil {

        resp = map[string]interface{}{

                   "error": err.Error(),

               }

    } else {



        if encodingErr := enc.Encode(resp); encodingErr != nil {

            fmt.Println("{ error: " + encodingErr.Error() + "}")

        }   

    }



}