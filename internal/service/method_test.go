package service

import "testing"

func TestGetLicitSoldier(t *testing.T) {
	soldier := GetLicitSoldier("1", "0", "1000")
	t.Log(soldier)
}

func TestGetRarity(t *testing.T) {
	rarity := GetRarity("10102")
	t.Log(rarity)
}

func TestGetCombatPoints(t *testing.T) {
	points := GetCombatPoints("10102")
	t.Log(points)
}

func TestGetCvcLicitSoldier(t *testing.T) {
	soldier := GetCvcLicitSoldier("1000")
	t.Log(soldier)
}

func TestGetUnkSoldier(t *testing.T) {
	soldier := GetUnkSoldier()
	t.Log(soldier)
}
