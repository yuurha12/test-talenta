package repositories

import (
	"gorm.io/gorm"
	"server/models"
)

type FriendRepository interface {
	FindFriends() ([]models.Friend, error)
	GetFriend(ID int) (models.Friend, error)
	CreateFriend(friend models.Friend) (models.Friend, error)
	UpdateFriend(friend models.Friend, ID int) (models.Friend, error)
	DeleteFriend(ID int) error
	GetFriendStats() (models.FriendStats, error)
}

type friendRepository struct {
	db *gorm.DB
}

func NewFriendRepository(db *gorm.DB) FriendRepository {
	return &friendRepository{db}
}

func (r *friendRepository) FindFriends() ([]models.Friend, error) {
	var friends []models.Friend
	err := r.db.Find(&friends).Error

	return friends, err
}

func (r *friendRepository) GetFriend(ID int) (models.Friend, error) {
	var friend models.Friend
	err := r.db.First(&friend, ID).Error

	return friend, err
}

func (r *friendRepository) CreateFriend(friend models.Friend) (models.Friend, error) {
	err := r.db.Create(&friend).Error

	return friend, err
}

func (r *friendRepository) UpdateFriend(friend models.Friend, ID int) (models.Friend, error) {
	err := r.db.Model(&models.Friend{}).
		Where("id = ?", ID).
		Updates(map[string]interface{}{
			"name":   friend.Name,
			"gender": friend.Gender,
			"age":    friend.Age,
		}).Error

	return friend, err
}

func (r *friendRepository) DeleteFriend(ID int) error {
	err := r.db.Delete(&models.Friend{}, ID).Error

	return err
}

func (r *friendRepository) GetFriendStats() (models.FriendStats, error) {
	var stats models.FriendStats

	err := r.db.Model(&models.Friend{}).
    Select("COUNT(*) as TotalFriendCount, "+
        "SUM(CASE WHEN gender = 'male' THEN 1 ELSE 0 END) as MaleCount, "+
        "SUM(CASE WHEN gender = 'female' THEN 1 ELSE 0 END) as FemaleCount, "+
        "SUM(CASE WHEN age < 19 THEN 1 ELSE 0 END) as Under19Count, "+
        "SUM(CASE WHEN age >= 20 THEN 1 ELSE 0 END) as Above20Count").
    Scan(&stats).Error
	return stats, err
}
