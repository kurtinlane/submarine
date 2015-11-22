package models

type Key struct {
  Email string //we will actually not want to store their email at all, just the hash
  Sha256 string
  DO_NOT_STORE_DO_NOT_LOG string
  App int64 //Id for which application this key belongs to
}

type App struct {
  SECRET_API_KEY string 
  Name string `datastore:",noindex"`
  Id int64 `datastore:"-"`
}