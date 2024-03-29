package delivery

import "net/http"

type InterfaceContactDelivery interface {
	CreateContact(http.ResponseWriter, *http.Request)
	ViewContact(http.ResponseWriter, *http.Request)
	UpdateContact(http.ResponseWriter, *http.Request)
	DeleteContact(http.ResponseWriter, *http.Request)

	CreateGroup(http.ResponseWriter, *http.Request)
	DeleteGroup(http.ResponseWriter, *http.Request)
	UpdateGroup(http.ResponseWriter, *http.Request)
	GetGroupByID(http.ResponseWriter, *http.Request)
}