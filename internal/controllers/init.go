package controllers

import "github.com/rishad004/SimpleBookManagementAPI/internal/service"

type ApiHandler struct {
	svc service.Svc
}

func NewControllers(svc service.Svc) *ApiHandler {
	return &ApiHandler{
		svc: svc,
	}
}
