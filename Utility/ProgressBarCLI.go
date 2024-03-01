package Utility

import "github.com/schollz/progressbar/v3"

type PB struct {
	progress *progressbar.ProgressBar
}

func StartProgressBar(count int) PB {
	bar := progressbar.NewOptions(count,
		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan] Getting from API..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	return PB{
		progress: bar,
	}
}

func (bar PB) UpdateStatus(count int) {
	bar.progress.Add(count)
}

func (bar PB) Finish() {
	bar.progress.Finish()
}
