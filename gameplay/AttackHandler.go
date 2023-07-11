package gameplay

import (
	"konn/entity/basic"
	"konn/utils"
)

func hitJudge(attacker basic.Arming, suffer basic.Substance) bool {
	return utils.GodsJudge(attacker.Accuracy - suffer.Props().Flex)
}

func damageHandler(attacker basic.Arming, suffer basic.Substance) int {
	// TODO: Damage calculate
	return 0
}
