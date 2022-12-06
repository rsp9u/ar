package ar

var (
	replaceTargets  = make(chan string)
	printableDiffs  = make(chan [3]string)
	isDoneScanning  = false
	isDoneReplacing = false
)

func DoneScan() {
	isDoneScanning = true
}

func DoneReplace() {
	isDoneReplacing = true
}
