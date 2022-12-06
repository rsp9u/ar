package ar

var (
	replaceTargets = make(chan string)
	printableDiffs = make(chan [3]string)
	scanDone       = false
	replaceDone    = false
)
