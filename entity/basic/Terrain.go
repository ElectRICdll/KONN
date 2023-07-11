package basic

type Terrain struct {
	name       string
	properties *Properties
}

func NewDesertTerrain() Terrain {
	return Terrain{
		name: "desert",
		properties: &Properties{
			Health:    0,
			Damage:    0,
			Armor:     0,
			AntiArmor: 0,
			Scout:     10,
			AntiScout: 30,
		},
	}
}

func NewForestTerrain() Terrain {
	return Terrain{
		name: "forest",
		properties: &Properties{
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
		name: "moutain",
		properties: &Properties{
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
		name: "plain",
		properties: &Properties{
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
		name: "coastal",
		properties: &Properties{
			Health:    0,
			Damage:    -40,
			Armor:     0,
			AntiArmor: 0,
			Scout:     0,
			AntiScout: 10,
		},
	}
}
