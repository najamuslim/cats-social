package usecase

import (
	"cats-social/model/database"
	"cats-social/model/dto"
)

type CatUsecaseInterface interface {
	GetCatById(id int) (err error)
	AddCat(request dto.RequestCreateCat) (id int64, err error)
	GetCat(request dto.RequestGetCat) (cats []dto.CatDetail, err error)
	UpdateCat(request dto.RequestCreateCat, id int64) (err error)
	CheckHasMatch(id int) (hasMatched bool, err error)
	DeleteCat(id int) (err error)
}

type AuthUsecaseInterface interface {
	Register(request dto.RequestCreateUser) (token string, err error)
	Login(request dto.RequestAuth) (token string, user database.User, err error)
	GetUserByEmail(email string) (exists bool, err error)
}

type MatchUsecaseInterface interface {
	CreateMatch(request dto.RequestCreateMatch, reqUserId int) error
	GetMatch(userId int) ([]dto.ResponseGetMatch, error)
	GetMatchById(id int) (err error)
	DeleteMatch(id int) (err error)
	ApproveMatch(id int, matchCatId int, userCatId int) (err error)
	RejectMatch(id int) (err error)
	GetCatIdByMatchId(id int) (matchCatId int, userCatIs int, err error)
}
