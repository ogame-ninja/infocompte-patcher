package main

import (
	ep "github.com/ogame-ninja/extension-patcher"
)

func main() {
	const (
		webstoreURL       = "https://openuserjs.org/install/The_Stubbs/InfoCompte.user.js"
		infocompte_sha256 = "cc12f9c7f17f09fc2fb947779ece471f79256db4bc9d7f882ecc99e6650b24ab"
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
// @match        *127.0.0.1*/bots/*/browser/html/*
// @match        *.ogame.ninja/bots/*/browser/html/*`, 1)
	return by
}
