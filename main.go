package main

import (
	ep "github.com/ogame-ninja/extension-patcher"
)

func main() {
	const (
		webstoreURL       = "https://openuserjs.org/install/The_Stubbs/InfoCompte.user.js"
		infocompte_sha256 = "832a181d884efc688afe2e136285f1b852494862368236ff18bc4dadc67722c9"
	)

	files := []ep.FileAndProcessors{
		ep.NewFile("InfoCompte.user.js", processInfoCompte),
	}

	ep.MustNew(ep.Params{
		ExpectedSha256: infocompte_sha256,
		WebstoreURL:    webstoreURL,
		Files:          files,
	}).Start()
}

var replN = ep.MustReplaceN

func processInfoCompte(by []byte) []byte {
	by = replN(by, `// @name         InfoCompte`, `// @name         InfoCompte Ninja`, 1)
	by = replN(by, `// @match        https://*.ogame.gameforge.com/game/*`, `{old}
// @match        *://*/bots/*/browser/html/*`, 1)
	by = replN(by, `https://${ window.location.host }/game/index.php?`, `${window.location.protocol}//${window.location.host}${window.location.pathname}?`, 1)
	return by
}
