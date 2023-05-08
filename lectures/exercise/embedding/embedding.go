//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

// * Create functions to calculate averages for each dashboard component
func (b *BandwidthUsage) Average() float32 {
	var sum Bytes = 0
	for i := 0; i < len(b.amount); i++ {
		sum += b.amount[i]
	}
	return float32(sum) / float32(len(b.amount))
}

func (t *CpuTemp) AverageCPUtemp() float32 {
	var sum Celcius = 0
	for i := 0; i < len(t.temp); i++ {
		sum += t.temp[i]
	}
	return float32(sum) / float32(len(t.temp))
}

func (m *MemoryUsage) Average() float32 {
	var sum Bytes = 0
	for i := 0; i < len(m.amount); i++ {
		sum += m.amount[i]
	}
	return float32(sum) / float32(len(m.amount))
}

type Averager interface {
	Average() float32
}

//   - Using struct embedding, create a Dashboard structure that contains
//     each dashboard component
type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dash := Dashboard{bandwidth, temp, memory}

	fmt.Println("Avg Bandwidth usage:", dash.BandwidthUsage.Average()) //using to showcase interface
	fmt.Println(dash.AverageCPUtemp())                                 // promoted function to Dashboard
	fmt.Println(dash.MemoryUsage.Average())                            //using to showcase interface, I am slowly understanding interfaces.

	var Stats Averager = &BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	fmt.Println("Avg Bandwidth usage using interface:", Stats.Average()) //Attempting to get the average using interface.
}
