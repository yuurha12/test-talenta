package frienddto

import "server/models"

type ResultStats struct {
	Code int                `json:"code"`
	Data models.FriendStats `json:"data"`
}

type FriendStatsResponse struct {
	TotalFriendCount int `json:"total_friend_count"`
	MaleCount        int `json:"male_count"`
	FemaleCount      int `json:"female_count"`
	Under19Count     int `json:"under_19_count"`
	Above20Count     int `json:"above_20_count"`
}
