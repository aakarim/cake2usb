package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func main() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Hello and welcome to...")
	fmt.Println()
	fmt.Println()
	myFigure := figure.NewFigure("Cake2USB", "", true)
	myFigure.Print()
	fmt.Println("																										v0.1")
	fmt.Println()

	color.Blue("Step 1. Please remove any wrapping, tissues or party favours from the cake")
	fmt.Println()
	fmt.Println("Press Enter to proceed...")
	fmt.Println()
	awaitConfirmation()

	color.Blue("Step 2. Insert the cake into your USB drive.")
	fmt.Println()
	fmt.Println("Tips:")
	fmt.Println()
	fmt.Println("	-- USB 3.0 slots are preferred, USB-C slots are too small to support cake.")
	fmt.Println("	-- Coffee cake may cause your machine's fans to spin at full speed. This is normal.")
	fmt.Println("	-- If your cake contains nuts or popcorn, please ensure that it is free of kernels. If you must upload the kernels please ensure")
	fmt.Println("	   That your system's kernel is compatible with the cake's kernel.")
	fmt.Println()
	fmt.Println("Press Enter to upload")
	awaitConfirmation()
	var wg sync.WaitGroup
	// pass &wg (optional), so p will wait for it eventually
	p := mpb.New(mpb.WithWaitGroup(&wg))
	barNames := []string{"Cake2Vec", "Decompiling frosting", "Redistributing Cake Tensors", "Recompiling frosting", "Sampling for flavour", "Reticulating splines", "Resampling", "Normalising cake"}
	total, numBars := 100, len(barNames)
	wg.Add(numBars)

	for i := 0; i < numBars; i++ {
		name := fmt.Sprintf("%v:", barNames[i])
		bar := p.AddBar(int64(total),
			mpb.PrependDecorators(
				// simple name decorator
				decor.Name(name),
				// decor.DSyncWidth bit enables column width synchronization
				decor.Percentage(decor.WCSyncSpace),
			),
			mpb.AppendDecorators(
				// replace ETA decorator with "done" message, OnComplete event
				decor.OnComplete(
					// ETA decorator with ewma age of 60
					decor.EwmaETA(decor.ET_STYLE_GO, 60), "done",
				),
			),
		)
		// simulating some work
		go func() {
			defer wg.Done()
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			max := 100 * time.Millisecond
			for i := 0; i < total; i++ {
				start := time.Now()
				time.Sleep(time.Duration(rng.Intn(10)+1) * max / 10)
				// since ewma decorator is used, we need to pass time.Since(start)
				bar.Increment(time.Since(start))
			}
		}()
	}
	// Waiting for passed &wg and for all bars to complete and flush
	p.Wait()

	cake := `.------------------------------------------------------------------------------.
|  _   _                           ____  _      _   _         _             _  |
| | | | | __ _ _ __  _ __  _   _  | __ )(_)_ __| |_| |__   __| | __ _ _   _| | |
| | |_| |/ _ ` + "`" + `| '_ \| '_ \| | | | |  _ \| | '__| __| '_ \ / _` + "`" + ` |/ _` + "`" + ` | | | | | |
| |  _  | (_| | |_) | |_) | |_| | | |_) | | |  | |_| | | | (_| | (_| | |_| |_| |
| |_| |_|\__,_| .__/| .__/ \__, | |____/|_|_|   \__|_| |_|\__,_|\__,_|\__, (_) |
|             |_|   |_|    |___/         .                            |___/    |
|                                     .'.:                                     |
|                                     '::'      ,                              |
|                                     .--.    ,` + "` `" + `.                            |
|                                     |  |   : .-` + "`" + `',-,                         |
|                                     |  |*'` + "` `" + `''.'.-` + "`" + `;                        |
|                                ,.-*'|  |       '.,,.,-,                      |
|                          ,.-*'` + "`" + `     '--'          .'.-` + "`" + `'                     |
|                    ,.-*'` + "`" + `                         '.,,-'                     |
|                 .'` + "`" + `-----------------------------------:                      |
|                 | ::::::::::::::::::::::::::::::::::: |                      |
|                 | ----------------------------------- |                      |
|                 |                                     |                      |
|                 | ----------------------------------- |                      |
|                 | ::::::::::::::::::::::::::::::::::: |                      |
|                 '-------------------------------------'                      |
|                                                                              |
'------------------------------------------------------------------------------'`
	cakeSlice := strings.Split(cake, "\n")
	for _, str := range cakeSlice {
		fmt.Println(str)
		time.Sleep(time.Millisecond * 400)
	}
	f, err := os.Create("birthday.cake")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	f.Write([]byte(cake))
	color.Red("You can now share cake")

}

func awaitConfirmation() {
	_, err := fmt.Scanln()
	if err != nil {
		panic(err)
	}
}
