package data

import (
	"github.com/jmoiron/sqlx"
)

// DB is the current database backend
var DB dbBackend

// dbBackend represents the database backend that the program will connect to
type dbBackend interface {
	Open() (*sqlx.DB, error)
	Setup() error
	DSN(string)

	AllArtists() ([]Artist, error)
	PurgeOrphanArtists() (int, error)
	DeleteArtist(*Artist) error
	LoadArtist(*Artist) error
	SaveArtist(*Artist) error

	AllAlbums() ([]Album, error)
	AlbumsForArtist(int) ([]Album, error)
	PurgeOrphanAlbums() (int, error)
	DeleteAlbum(*Album) error
	LoadAlbum(*Album) error
	SaveAlbum(*Album) error

	AllSongs() ([]Song, error)
	SongsForAlbum(int) ([]Song, error)
	SongsForArtist(int) ([]Song, error)
	SongsInPath(string) ([]Song, error)
	SongsNotInPath(string) ([]Song, error)
	DeleteSong(*Song) error
	LoadSong(*Song) error
	SaveSong(*Song) error

	DeleteUser(*User) error
	LoadUser(*User) error
	SaveUser(*User) error

	DeleteSession(*Session) error
	LoadSession(*Session) error
	SaveSession(*Session) error
	UpdateSession(*Session) error
}
