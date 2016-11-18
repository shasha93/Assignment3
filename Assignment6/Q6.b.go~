package main

import (
	"fmt"
)
type Planet interface{

 Name()string
 Mass()int64
}
type MyPlanet struct{
 NameOfPlanet string
 MassOfPlanet int64
}
func(p *MyPlanet) Name()string{
 return p.NameOfPlanet
}
func(m *MyPlanet)Mass()int64{
return m.MassOfPlanet
}

func (p MyPlanet) String() string {
	
return fmt.Sprintf("Name:%s  Mass: %d Ton  \n",p.NameOfPlanet,p.MassOfPlanet)
}
func main() {
	marsh:=MyPlanet{NameOfPlanet:"Marsh",MassOfPlanet:254343434}
	jupiter:=MyPlanet{NameOfPlanet:"jupiter",MassOfPlanet:305432}
        pluto:=MyPlanet{NameOfPlanet:"Pluto",MassOfPlanet:2599000}
        neptune:=MyPlanet{NameOfPlanet:"Neptune",MassOfPlanet:423455}
	fmt.Printf("%s",marsh)
	fmt.Printf("%s",jupiter)
        fmt.Printf("%s",pluto)
        fmt.Printf("%s",neptune)
}

