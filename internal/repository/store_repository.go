package repository

import "github.com/ISSuh/msago-sample/internal/domain/model"

type StoreRepository interface {
	StoreById(int) model.Store
}
