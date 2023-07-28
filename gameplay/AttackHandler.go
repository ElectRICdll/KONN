package gameplay

import (
	"konn/entity/basic"
	"konn/entity/prop"
	"konn/utils"
)

func HitJudge(attacker basic.Arming, suffer prop.Substantive) bool {
	return utils.GodsJudge(attacker.Accuracy - suffer.Props().Flex)
}

func DamageHandler(attacker basic.Arming, suffer prop.Substantive) int {
	// TODO: Damage calculate
	return 0
}
