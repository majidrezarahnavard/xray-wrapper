package vmess_maker

import (
	"fmt"
	"xray-wrapper/vmess_maker/service/builder"
	"xray-wrapper/vmess_maker/service/execute"
	"xray-wrapper/vmess_maker/service/subscribe"
	"xray-wrapper/vmess_maker/service/telegram"
)

func main() {

	executeInstance := execute.NewExecute()
	executeInstance.ExecuteCommand("./reinstall.sh")

	fmt.Println("build config ...")

	builderInstance := builder.NewBuilder().
		SetServerIP().
		SetSettingsFile().
		SetConfigurations().
		SetBlock().
		Save()

	sub := subscribe.NewSubscribe(builderInstance, executeInstance)

	tel := telegram.NewTelegramClient(builderInstance, sub)

	tel.SendVNstat()

}
