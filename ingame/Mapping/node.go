package mapping

import (
	. "konn/ingame/basic"
)

type Node struct {
	no         int
	contains   map[Substantive]int
	curTerrain *Terrain
	curWeather *Weather
}

func IsConnected(n1 *Node, n2 *Node) bool {
	return false
}

type Position interface {
	Pos() Position
}

type Terrain struct {
	Name       string
	Properties *PropertiesOffset
}

func NewDesertTerrain() *Terrain {
	return &Terrain{
		Name: "desert",
		Properties: &PropertiesOffset{
			Health: 0,
			Damage: 0,
			Armor: 0,
			AntiArmor: 0,
			Scout: 30,
			AntiScout: -30,
		},
	}
}

func NewForestTerrain() *Terrain {
	return &Terrain{
		Name: "forest",
		Properties: &PropertiesOffset{
			Health: 0,
			Damage: 0,
			Armor: 0,
			AntiArmor: 0,
			Scout: -10,
			AntiScout: 30,
		},
	}
}

func NewMoutainTerrain() *Terrain {
	return &Terrain {
		Name: "moutain",
		Properties: &PropertiesOffset{
			Health: 0,
			Damage: 0,
			Armor: 0,
			AntiArmor: 0,
			Scout: 30,
			AntiScout: 0,
		},
	}
}

func NewPlainTerrain() *Terrain {
	return &Terrain{
		Name: "plain",
		Properties: &PropertiesOffset{
			Health: 0,
			Damage: 0,
			Armor: 0,
			AntiArmor: 0,
			Scout: 20,
			AntiScout: -20,
		},
	}
}

func NewCoastalTerrain() *Terrain {
	return &Terrain{
		Name: "coastal",
		Properties: &PropertiesOffset{
			Health: 0,
			Damage: -40,
			Armor: 0,
			AntiArmor: 0,
			Scout: 0,
			AntiScout: 10,
		},
	}
}

type Weather struct {
	Name       string
	Properties *PropertiesOffset
}

func NewCommonWeather() *Weather {
	return &Weather{
		Name: "common",
		Properties: &PropertiesOffset{
			Health: 0,
			Damage: 0,
			Armor: 0,
			AntiArmor: 0,
			Scout: 0,
			AntiScout: 0,
		},
	}
}

func NewRainyWeather() *Weather {
	return &Weather{
		Name: "rainy",
		Properties: &PropertiesOffset{
			Health: 0,
			Damage: 0,
			Armor: 0,
			AntiArmor: 0,
			Scout: -10,
			AntiScout: 10,
		},
	}
}

func NewSnowyWeather() *Weather {
	return &Weather{
		Name: "common",
		Properties: &PropertiesOffset{
			Health: 0,
			Damage: 10,
			Armor: 0,
			AntiArmor: 0,
			Scout: -20,
			AntiScout: 0,
		},
	}
}
