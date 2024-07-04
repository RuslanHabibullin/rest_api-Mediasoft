package repository

import (
	"fmt"
	restapimediasoft "tgtest/Desktop/Univercity/go_proj/rest_api_mediasoft/rest_api-Mediasoft"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type FurnituresListPostgres struct {
	db *sqlx.DB
}

func NewFurnituresListPostgres(db *sqlx.DB) *FurnituresListPostgres {
	return &FurnituresListPostgres{db: db}
}
func (r *FurnituresListPostgres) Create(list restapimediasoft.Furnitures_list) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (name, manufacturer, height, width, lenght) VALUES ($1, $2, $3, $4, $5) RETURNING id", furnituresTable)
	row := tx.QueryRow(createListQuery, list.Name, list.Manufacturer, list.Height, list.Lenght, list.Width)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (r *FurnituresListPostgres) GetAll() ([]restapimediasoft.Furnitures_list, error) {
	var lists []restapimediasoft.Furnitures_list
	query := fmt.Sprintf("SELECT tl.name, tl.manufacturer, tl.height, tl.lenght, tl.width FROM %s tl",
		furnituresTable)
	err := r.db.Select(&lists, query)
	return lists, err
}

func (r *FurnituresListPostgres) GetByID(listID int) (restapimediasoft.Furnitures_list, error) {
	var list restapimediasoft.Furnitures_list
	query := fmt.Sprintf("SELECT tl.name, tl.manufacturer, tl.height, tl.lenght, tl.width FROM %s tl WHERE tl.id = $1",
		furnituresTable)
	err := r.db.Get(&list, query, listID)
	return list, err
}

func (r *FurnituresListPostgres) DeleteList(listID int) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.id = $1",
		furnituresTable)
	_, err := r.db.Exec(query, listID)
	return err
}

func (r *FurnituresListPostgres) UpdateAll(listID int, input restapimediasoft.UpdateListInput) error {
	setQuery, argId, args := UpdateChecker(input)
	if argId == 6 {
		query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d",
			furnituresTable, setQuery, argId)
		args = append(args, listID)
		logrus.Debugf("updateQuery: %s", query)
		logrus.Debugf("args: %s", args)
		_, err := r.db.Exec(query, args...)
		return err
	} else {
		return nil
	}

}

func (r *FurnituresListPostgres) Update(listID int, input restapimediasoft.UpdateListInput) error {

	setQuery, argId, args := UpdateChecker(input)
	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.id = $%d",
		furnituresTable, setQuery, argId)
	args = append(args, listID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err

}
