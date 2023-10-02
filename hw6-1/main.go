package main

import (
	"fmt"
	"strings"
)

func main() {
	deliviries := make([]Package, 0)
	deliviries = append(deliviries, &envelope{
		postalInfo: postalInfo{
			srcAddr: "Kyiv, Pravdy 1a",
			dstAddr: "Kyiv, Basseynaya street, 99b",
		},
	})
	deliviries = append(deliviries, &box{
		postalInfo: postalInfo{
			srcAddr: "Kyiv, Pravdy 1a",
			dstAddr: "Kyiv, Basseynaya street, 99b",
		},
	})

	for _, delivery := range deliviries {
		fmt.Println(strings.Repeat("-", 20))
		makeDelivery(delivery)
		fmt.Printf(strings.Repeat("-", 20))
	}

}

type envelope struct {
	postalInfo postalInfo
}

func (envelope *envelope) deliver(way string) {
	fmt.Printf("Source address: %s\nDestination address: %s\nDelivery way: %s\n", envelope.postalInfo.srcAddr, envelope.postalInfo.dstAddr, way)
}

type box struct {
	postalInfo postalInfo
}

func (box *box) deliver(way string) {
	fmt.Printf("Source address: %s\nDestination address: %s\nDelivery way: %s\n", box.postalInfo.srcAddr, box.postalInfo.dstAddr, way)
}

type postalInfo struct {
	srcAddr string
	dstAddr string
}

type Package interface {
	deliver(string)
}

func makeDelivery(smthng Package) {
	smthng.deliver(chooseDeliveryWay(smthng))
}

func chooseDeliveryWay(smthng Package) string {

	switch smthng.(type) {
	case *envelope:
		return "foot"
	case *box:
		return "car"
	default:
		return "railway"
	}

}
