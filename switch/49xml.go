package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id int `xml:"id,attr"`
	Name string `xml:"name"`
	Origin []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plact id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{
		Id:      27,
		Name:    "Coffee",
	}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	/*
	 <plant id="27">
	   <name>Coffee</name>
	   <origin>Ethiopia</origin>
	   <origin>Brazil</origin>
	 </plant>
	*/

	fmt.Println(xml.Header + string(out))
	/*
	<?xml version="1.0" encoding="UTF-8"?>
	 <plant id="27">
	   <name>Coffee</name>
	   <origin>Ethiopia</origin>
	   <origin>Brazil</origin>
	 </plant>
	*/

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p) //
	/*
		Plact id=27, name=Coffee, origin=[Ethiopia Brazil]
	*/

	tomato := &Plant{
		Id:      81,
		Name:    "tomato",
	}
	tomato.Origin = []string{"mexico", "california"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
	/*
	 <nesting>
	   <parent>
	     <child>
	       <plant id="27">
	         <name>Coffee</name>
	         <origin>Ethiopia</origin>
	         <origin>Brazil</origin>
	       </plant>
	       <plant id="81">
	         <name>tomato</name>
	         <origin>mexico</origin>
	         <origin>california</origin>
	       </plant>
	     </child>
	   </parent>
	 </nesting>

	*/
}
