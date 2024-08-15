package service

import (
	"github.com/ggorockee/airbnb-backend/data/request"
	"github.com/ggorockee/airbnb-backend/data/response"
)

type NoteService interface {
	Create(note request.CreateNoteRequest)
	Update(note request.UpdateNoteRequest)
	Delete(noteId int)
	FindById(noteId int) response.NoteResponse
	FindAll() []response.NoteResponse
}
