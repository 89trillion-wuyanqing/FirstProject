package handler

import (
	"reflect"
	"testing"
)

func TestGetArmyAtk(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"test1", "10101", "100"},
		{"test2", "10502", "10502"},
		{"test3", "10502", "10502"},
	}
	armyhandler := ArmyHandler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := armyhandler.GetArmyAtk(tt.args); !reflect.DeepEqual(got.Data, tt.want) {
				t.Errorf("GetCiftCodes() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestGetArmys(t *testing.T) {
	tests := []struct {
		name        string
		rarity      string
		unlockArena int
		cvc         string
		want        string
	}{
		{"test1", "2", 2, "1000", "100"},
		{"test2", "3", 3, "1000", "100"},
		{"test3", "1", 1, "1000", "100"},
	}
	armyhandler := ArmyHandler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := armyhandler.GetArmys(tt.rarity, tt.unlockArena, tt.cvc); !reflect.DeepEqual(got.Data, tt.want) {
				t.Errorf("GetCiftCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArmyRarity(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"test1", "10101", "3"},
		{"test2", "10502", "2"},
		{"test3", "10502", "1"},
	}
	armyhandler := ArmyHandler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := armyhandler.GetArmyRarity(tt.args); !reflect.DeepEqual(got.Data, tt.want) {
				t.Errorf("GetCiftCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetArmysByCvc(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"test1", "1000", "100"},
	}
	armyhandler := ArmyHandler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := armyhandler.GetArmysByCvc(tt.args); !reflect.DeepEqual(got.Data, tt.want) {
				t.Errorf("GetArmyAtk() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetArmysByStage(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"test1", "100"},
	}
	armyhandler := ArmyHandler{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := armyhandler.GetArmysByStage(); !reflect.DeepEqual(got.Data, tt.want) {
				t.Errorf("GetCiftCodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
