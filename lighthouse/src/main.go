package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"strconv"
)

func main() {
	rc := 0
	var csvFileName string
	var perfThreshold string
	var a11yThreshold string
	var bpThreshold string
	var seoThreshold string
	var pwaThreshold string

	//read cli flags
	flag.StringVar(&csvFileName, "file", "summary.csv", "CSV file name")
	flag.StringVar(&perfThreshold, "perf", "0.8", "Performance threshold")
	flag.StringVar(&a11yThreshold, "a11y", "0.8", "Accessibility threshold")
	flag.StringVar(&bpThreshold, "bp", "0.8", "Best practice threshold")
	flag.StringVar(&seoThreshold, "seo", "0.8", "SEO threshold")
	flag.StringVar(&pwaThreshold, "pwa", "0.8", "PWA threshold")
	flag.Parse()

	// parse threshold values
	perfT, _ := strconv.ParseFloat(perfThreshold, 64)
	a11yT, _ := strconv.ParseFloat(a11yThreshold, 64)
	bpT, _ := strconv.ParseFloat(bpThreshold, 64)
	seoT, _ := strconv.ParseFloat(seoThreshold, 64)
	pwaT, _ := strconv.ParseFloat(pwaThreshold, 64)

	// read csv file
	csvFile, err := os.Open(csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	records, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// check records with thresholds
	for i, record := range records {
		if i > 0 {
			url := record[0]
			perfVal, _ := strconv.ParseFloat(record[4], 64)
			a11yVal, _ := strconv.ParseFloat(record[5], 64)
			bpVal, _ := strconv.ParseFloat(record[6], 64)
			seoVal, _ := strconv.ParseFloat(record[7], 64)
			pwaVal, _ := strconv.ParseFloat(record[8], 64)

			if (perfVal > 0 && perfVal < perfT) || (a11yVal > 0 && a11yVal < a11yT) || (bpVal > 0 && bpVal < bpT) || (seoVal > 0 && seoVal < seoT) || (pwaVal > 0 && pwaVal < pwaT) {
				log.Printf("%s - %.2f - %.2f - %.2f - %.2f - %.2f", url, perfVal, a11yVal, bpVal, seoVal, pwaVal)
				rc++
			}

		}
	}

	// return with error > 0 depending on threshold violations
	os.Exit(rc)
}
