package main

// Give us some seed data
func init() {
  RepoCreateData(&Data{Key: "a", Value: "1", Encoding: "integer"})
  RepoCreateData(&Data{Key: "b", Value: "2", Encoding: "integer"})
}

func RepoFindData(key string) *Data {
    return store[key]
}

func RepoCreateData(d *Data) *Data {
	key := RepoGetKey(d)
	store[key] = d
  return d
}

func RepoGetKey(d *Data) string  {
	return d.Key
}
