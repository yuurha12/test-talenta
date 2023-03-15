package models

type Friend struct {
	ID     int   `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}


type FriendStats struct {
	TotalFriendCount int `json:"total_friend_count"`
	MaleCount        int `json:"male_count"`
	FemaleCount      int `json:"female_count"`
	Under19Count     int `json:"under_19_count"`
	Above20Count     int `json:"above_20_count"`
}
