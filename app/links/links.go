package links

import (
	"errors"
	"fmt"
	"net/http"
	"shortener/customtypes"
	"shortener/repository"
)

func AddLink(link customtypes.NewLink, conn customtypes.Conn) (int, error) {
	fmt.Println("Frtom add link", link.Alias, "Alias", link.Url, "Url")
	success, err := repository.SetExKey(conn, link.Alias, link.Url)

	if !success && err != nil {
		return http.StatusInternalServerError, errors.New("Unable to insert link")
	}

	if !success {
		return http.StatusConflict, errors.New("A link already exists with given alias")
	}

	return http.StatusCreated, nil
}

func GetLink(alias string, conn customtypes.Conn) (int, string, error) {
	value, err := repository.GetKey(conn, alias)

	if err != nil {
		return http.StatusInternalServerError, "", err
	}

	if value == "" {
		return http.StatusNotFound, "", errors.New("No url found for given alias")
	}

	return http.StatusOK, value, nil
}
