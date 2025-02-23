/*
A man has a rather old car being worth $2000. He saw a secondhand car being worth $8000. He wants to keep his old car until he can buy the secondhand one.

He thinks he can save $1000 each month but the prices of his old car and of the new one decrease of 1.5 percent per month. Furthermore this percent of loss increases of 0.5 percent at the end of every two months. Our man finds it difficult to make all these calculations.
*/

package kata

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	month := 0
	priceOld := float64(startPriceOld)
	priceNew := float64(startPriceNew)

	for ; priceOld+float64(month*savingperMonth) < priceNew; month++ {
		if month%2 == 1 {
			percentLossByMonth += 0.5
		}
		priceOld -= priceOld * percentLossByMonth / 100.0
		priceNew -= priceNew * percentLossByMonth / 100.0
	}
	return [2]int{month, int(priceOld + float64(month*savingperMonth) - priceNew + 0.5)}
}
