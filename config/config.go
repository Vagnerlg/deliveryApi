package config

import "os"

type db struct {
	Uri  string
	Name string
}

func DB() db {
	return db{
		Uri:  os.Getenv("DATABASE_URI"),
		Name: os.Getenv("DATABASE_NAME"),
	}
}
