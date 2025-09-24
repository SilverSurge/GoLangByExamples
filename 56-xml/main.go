package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	fmt.Printf("XML\n\n")

	// XML and struct
	peashooter := &Plant{Id: 11, Name: "PeaShooter"}
	peashooter.Origin = []string{"bob's backyard", "zobielane"}
	out, _ := xml.MarshalIndent(peashooter, " ", " ")
	fmt.Println(string(out))
	fmt.Println(xml.Header + string(out))

	// unmarshalling: struct
	var pea Plant
	if err := xml.Unmarshal(out, &pea); err != nil {
		panic(err)
	}
	fmt.Println(pea)

	// nesting
	sunflower := &Plant{Id: 29, Name: "Sunflower"}

	type Nursery struct {
		XMLName xml.Name `xml:"nursery"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nursery := &Nursery{}
	nursery.Plants = []*Plant{peashooter, sunflower}

	out, _ = xml.MarshalIndent(nursery, " ", "  ")
	fmt.Println(string(out))
}
