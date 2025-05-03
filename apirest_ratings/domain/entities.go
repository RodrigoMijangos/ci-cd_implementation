package domain

type User struct {
	Id         int
	Username   string
	First_name string
	Last_name  string
	Email      string
	Gender     string
}

type Tag struct {
	Userid    string
	Movieid   string
	Tag       string
	Timestamp string
}

type Movie struct {
	Movieid string
	Title   string
	Genres  string
}

type Link struct {
	Movieid string
	Imdbid  string
	Tmdbid  string
}

type Rating struct {
	Userid    string
	Movieid   string
	Rating    string
	Timestamp string
}
