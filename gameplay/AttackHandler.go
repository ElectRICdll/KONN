package gameplay

import (
	"fmt"
	"konn/entity/basic"
	"konn/entity/prop"
	"konn/utils"
)

func Hit(attacker basic.Arming, suffer prop.Substantive) {
	if attacker.IsDisable == true {
		defer utils.Logger.Debug("Arming Disabled.")
		return
	} else if hitJudge(attacker, suffer) == true {
		damage := damage(attacker, suffer)
		utils.Logger.Info(fmt.Sprintf("Hit confirmed. Damage is %d", damage))
	} else {
		utils.Logger.Info("Hit failed.")
	}
}

func hitJudge(attacker basic.Arming, suffer prop.Substantive) bool {
	return utils.GodsJudge(attacker.Accuracy - suffer.Props().Flex)
}

func damage(attacker basic.Arming, suffer prop.Substantive) int {
	var result int
	// TODO: Damage calculate
	return result
}
