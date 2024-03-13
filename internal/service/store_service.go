package service

import "github.com/ISSuh/msago-sample/internal/domain/model"

type StoreService interface {
	StoreById(int) model.Store
}
