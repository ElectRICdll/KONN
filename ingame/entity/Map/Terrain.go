package Map

import (
	"konn/ingame/entity/BasicClass"
)

type Terrain struct {
	Name       string
	Properties *BasicClass.Properties
}

func NewDesertTerrain() Terrain {
	return Terrain{
		Name: "desert",
		Properties: &BasicClass.Properties{
			Health:    0,
			Damage:    0,
			Armor:     0,
			AntiArmor: 0,
			Scout:     30,
			AntiScout: -30,
		},
	}
}

func NewForestTerrain() Terrain {
	return Terrain{
		Name: "forest",
		Properties: &BasicClass.Properties{
			Health:    0,
			Damage:    0,
			Armor:     0,
			AntiArmor: 0,
			Scout:     -10,
			AntiScout: 30,
		},
	}
}

func NewMoutainTerrain() Terrain {
	return Terrain{
		Name: "moutain",
		Properties: &BasicClass.Properties{
			Health:    0,
			Damage:    0,
			Armor:     0,
			AntiArmor: 0,
			Scout:     30,
			AntiScout: 0,
		},
	}
}

func NewPlainTerrain() Terrain {
	return Terrain{
		Name: "plain",
		Properties: &BasicClass.Properties{
			Health:    0,
			Damage:    0,
			Armor:     0,
			AntiArmor: 0,
			Scout:     20,
			AntiScout: -20,
		},
	}
}

func NewCoastalTerrain() Terrain {
	return Terrain{
		Name: "coastal",
		Properties: &BasicClass.Properties{
			Health:    0,
			Damage:    -40,
			Armor:     0,
			AntiArmor: 0,
			Scout:     0,
			AntiScout: 10,
		},
	}
}