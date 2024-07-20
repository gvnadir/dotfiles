package main

import (
	"fmt"

	// "example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

// For every taxRates, it reads the prices.txt file, stores the prices into the prices struct
// and for every price, it creates a json file with the tax applied
func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.5}
	doneChans := make([]chan bool, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// err := priceJob.Process()
		go priceJob.Process(doneChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for _, doneChan := range doneChans {
		// I read for every channel
		<-doneChan //This tells Go that we wait until every channel has emmited one value
	}

}
