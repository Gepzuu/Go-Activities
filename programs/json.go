package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// Define a struct to represent a music response with an integer page and a list of fruits
type musicResponse struct {
    Page   int      `json:"page"`
    Genres []string `json:"genres"`
}

func main() {
    // Marshal different data types to JSON and print the results
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("guitar")
    fmt.Println(string(strB))

    slcD := []string{"pop", "rock", "jazz"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    // Create and marshal a map representing music genres and their popularity
    mapD := map[string]int{"pop": 5, "rock": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // Create and marshal a music response struct
    musicRes := &musicResponse{
        Page:   1,
        Genres: []string{"pop", "rock", "jazz"}}
    musicResB, _ := json.Marshal(musicRes)
    fmt.Println(string(musicResB))

    // Unmarshal JSON data into a map and access specific values
    byt := []byte(`{"popularity":6.13,"genres":["pop","rock"]}`)
    var dat map[string]interface{}
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    pop := dat["popularity"].(float64)
    fmt.Println(pop)

    genres := dat["genres"].([]interface{})
    genre1 := genres[0].(string)
    fmt.Println(genre1)

    // Unmarshal JSON data into a music response struct
    musicStr := `{"page": 1, "genres": ["pop", "rock"]}`
    musicRes2 := musicResponse{}
    json.Unmarshal([]byte(musicStr), &musicRes2)
    fmt.Println(musicRes2)
    fmt.Println(musicRes2.Genres[0])

    // Encode a map representing music genres and their popularity to JSON and print it
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"pop": 5, "rock": 7}
    enc.Encode(d)
}

/*

# Output

true
1
2.34
"guitar"
["pop","rock","jazz"]
{"pop":5,"rock":7}
{"page":1,"genres":["pop","rock","jazz"]}
map[genres:[pop rock] popularity:6.13]
6.13
pop
{1 [pop rock]}
pop
{"pop":5,"rock":7}


*/