package storage

var m = make(map[string]string)

func StoreData(name, data string){
	m[name] = data
}

func RetrieveData(name string) string{
	return m[name]
}