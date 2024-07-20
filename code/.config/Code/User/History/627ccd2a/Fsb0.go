package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

// It reads price.txt and stores the prices into InputPrices
func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	fmt.Println("LoadData", job.InputPrices)
	return nil
}

// Apply taxes to each price and writes the results to the json files
func (job *TaxIncludedPriceJob) Process(doneChan chan bool) {
	err := job.LoadData()

	if err != nil {
		// return err
	}

	result := make(map[string]string)
	fmt.Println("Process", job.InputPrices)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price + (price * job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	// return job.IOManager.WriteResult(job)
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}

}
