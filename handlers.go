package main

import (
  "fmt"
  "net/http"
  "io"
  "io/ioutil"
  "encoding/json"
  "hash/fnv"
  "bytes"

  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintln(w, "Welcome to my key value store!")
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	var data Data
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
      panic(err)
  }
  if err := r.Body.Close(); err != nil {
      panic(err)
  }
  if err := json.Unmarshal(body, &data); err != nil {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(422) // unprocessable entity
      if err := json.NewEncoder(w).Encode(err); err != nil {
          panic(err)
      }
  }
  key := RepoGetKey(&data)
  server := GetServer(key)
  port := prefix + serverId
  if server != port {
    fmt.Println("Delegating to server at port" + server)
    jsonData := map[string]string{"key": data.Key, "value": data.Value, "encoding": data.Encoding}
    jsonValue, _ := json.Marshal(jsonData)
    http.Post("http://localhost:"+ server +"/set", "application/json", bytes.NewBuffer(jsonValue))
  } else {
    fmt.Println("Great. I will create the key")
    d := RepoCreateData(&data)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(d); err != nil {
        panic(err)
    }
  }
}

func GetData(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  key := vars["key"]
  server := GetServer(key)
  port := prefix + serverId
  if server != port {
    fmt.Println("Your key is with server at port " + server)
    response, err := http.Get("http://localhost:" + server + "/get/" + key)
    if err != nil {
        panic(err)
    } else {
      fmt.Println("Great I have your key")
      data, _ := ioutil.ReadAll(response.Body)
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusOK)
      w.Write(data)
    }
  } else {
  	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(store[key].Value); err != nil {
        panic(err)
    }
  }
}

func GetServer(key string) string  {
  id := hash(key) % n
  return "510" + fmt.Sprint(id)
}

func hash(key string) uint32 {
  h := fnv.New32a()
  h.Write([]byte(key))
  return h.Sum32()
}
