package generator

import log "github.com/sirupsen/logrus"

const WeaselWordsURL = "https://raw.githubusercontent.com/btford/weasel-words/master/weasel.js"

func GenerateWeaselWords() error {
	lines, err := RemoteFileLines(WeaselWordsURL)
	if err != nil {
		return err
	}

	weasels, err := LineSliceExclusive(lines, "^var weasels = \\[", "^\\];")
	if err != nil {
		log.Info("Got here")
		return err
	}
	TrimList(weasels)

	log.Info("Weasel count: ", len(weasels))
	log.Info("Weasels: ", weasels)

	exceptions, err := LineSliceExclusive(lines, "^var exceptions = \\[", "^\\]")
	if err != nil {
		return err
	}
	TrimList(exceptions)

	log.Info("Exception count: ", len(exceptions))
	log.Info("Exceptions: ", exceptions)

	return nil
}
