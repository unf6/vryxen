package browsers

import (
	"path/filepath"

	_ "modernc.org/sqlite"
)

func (c *Chromium) GetBookmarks(path string) (bookmarks []Bookmark, err error) {
	db, err := GetDBConnection(filepath.Join(path, "Default", "Web Data"))
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT name, value FROM autofill")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var name, value string

		if err := rows.Scan(&name, value); err != nil {
			continue
		}
		if name == "" || value == "" {
			continue
		}

		bookmarks = append(bookmarks, Bookmark{
			Name:          name,
			Value:         value,
		})

	}

		return bookmarks, nil	
}
