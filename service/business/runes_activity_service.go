package business

import (
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	"github.com/Pump-Elf-Ranch/per_apollo/database/business"
)

type RunesActivityService interface {
	RunesActivityList(activityType []int, pageNum, pageSize int) ([]business.RunesActivity, int64)
	RunesActivityInfo(orderId string) business.RunesActivity
}

type runesActivityService struct {
	db *database.DB
}

func NewRunesActivityService(db *database.DB) RunesActivityService {
	return &runesActivityService{db: db}
}

func (r runesActivityService) RunesActivityList(activityType []int, pageNum, pageSize int) ([]business.RunesActivity, int64) {
	return r.db.RunesActivity.ListRunesActivity(activityType, pageNum, pageSize)
}

func (r runesActivityService) RunesActivityInfo(orderId string) business.RunesActivity {
	//TODO implement me
	panic("implement me")
}
