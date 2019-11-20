package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/UlisesCampos/myspace/commons"
	"github.com/UlisesCampos/myspace/configuration"
	"github.com/UlisesCampos/myspace/models"
)

// VoteRegister controlador para registrar un voto
func VoteRegister(w http.ResponseWriter, r *http.Request) {
	vote := models.Vote{}
	user := models.User{}
	currentVote := models.Vote{}
	m := models.Message{}

	user, _ = r.Context().Value("user").(models.User)
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		m.Message = fmt.Sprintf("Error al leer el voto a registrar: %s", err)
		m.Code = http.StatusBadRequest
		commons.DisplayMessage(w, m)
		return
	}
	vote.UserID = user.ID

	db := configuration.GetConnection()
	defer db.Close()

	db.Where("comment_id = ? and user_id = ?", vote.CommentID, vote.UserID).First(&currentVote)

	// Si no existe
	if currentVote.ID == 0 {
		db.Create(&vote)
		err := updateCommentVotes(vote.CommentID, vote.Value, false)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}
		m.Message = "Voto registrado"
		m.Code = http.StatusCreated
		commons.DisplayMessage(w, m)
		return
	} else if currentVote.Value != vote.Value {
		currentVote.Value = vote.Value
		db.Save(&currentVote)
		updateCommentVotes(vote.CommentID, vote.Value, true)
		if err != nil {
			m.Message = err.Error()
			m.Code = http.StatusBadRequest
			commons.DisplayMessage(w, m)
			return
		}
		m.Message = "Voto actualizado"
		m.Code = http.StatusOK
		commons.DisplayMessage(w, m)
		return
	}
	m.Message = "Este voto ya esta registrado"
	m.Code = http.StatusBadRequest
	commons.DisplayMessage(w, m)
}

//Actualiza la cantidad de votos en la tabla comentarios
//isUpdate indica si es un voto para actualizar
func updateCommentVotes(commentID uint, vote bool, isUpdate bool) (err error) {
	comment := models.Comment{}

	db := configuration.GetConnection()
	defer db.Close()

	rows := db.First(&comment, commentID).RowsAffected

	if rows > 0 {
		if vote {
			comment.Votes++
			if isUpdate {
				comment.Votes++
			}
		} else {
			comment.Votes--
			if isUpdate {
				comment.Votes--
			}
		}
		db.Save(&comment)
	} else {
		err = errors.New("No se encontro un registro de comentario para asignarle el voto")
	}
	return
}
