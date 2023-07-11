package basic

type Weather struct {
	Name       string
	Properties *Properties
}

func NewCommonWeather() Weather {
	return Weather{
		Name: "common",
		Properties: &Properties{
			Health:    0,
			Damage:    0,
			Armor:     0,
			AntiArmor: 0,
			Scout:     0,
			AntiScout: 0,
		},
	}
}

func NewRainyWeather() Weather {
	return Weather{
		Name: "rainy",
		Properties: &Properties{
			Health:    0,
			Damage:    0,
			Armor:     0,
			AntiArmor: 0,
			Scout:     -10,
			AntiScout: 10,
		},
	}
}

func NewSnowyWeather() Weather {
	return Weather{
		Name: "common",
		Properties: &Properties{
			Health:    0,
			Damage:    10,
			Armor:     0,
			AntiArmor: 0,
			Scout:     -20,
			AntiScout: 0,
		},
	}
}
