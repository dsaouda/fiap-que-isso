package main

import (
	"fmt"
	"github.com/dsaouda/fiap-que-isso/admin/fixture/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Image struct {
	Name string
	Filename string
	Group string
}

type Slide struct {
	Name string
	Type string
	Value string
	Group string
}


func main() {

	slides := []Slide{
		//alfabeto
		Slide{"Letra A", "text", "A", "alfabeto"},
		Slide{"Letra B", "text", "B", "alfabeto"},
		Slide{"Letra C", "text", "C", "alfabeto"},
		Slide{"Letra D", "text", "D", "alfabeto"},
		Slide{"Letra E", "text", "E", "alfabeto"},
		Slide{"Letra F", "text", "F", "alfabeto"},
		Slide{"Letra G", "text", "G", "alfabeto"},
		Slide{"Letra H", "text", "H", "alfabeto"},
		Slide{"Letra I", "text", "I", "alfabeto"},
		Slide{"Letra J", "text", "J", "alfabeto"},
		Slide{"Letra K", "text", "K", "alfabeto"},
		Slide{"Letra L", "text", "L", "alfabeto"},
		Slide{"Letra M", "text", "M", "alfabeto"},
		Slide{"Letra N", "text", "N", "alfabeto"},
		Slide{"Letra O", "text", "O", "alfabeto"},
		Slide{"Letra P", "text", "P", "alfabeto"},
		Slide{"Letra Q", "text", "Q", "alfabeto"},
		Slide{"Letra R", "text", "R", "alfabeto"},
		Slide{"Letra S", "text", "S", "alfabeto"},
		Slide{"Letra T", "text", "T", "alfabeto"},
		Slide{"Letra U", "text", "U", "alfabeto"},
		Slide{"Letra V", "text", "V", "alfabeto"},
		Slide{"Letra W", "text", "W", "alfabeto"},
		Slide{"Letra X", "text", "X", "alfabeto"},
		Slide{"Letra Y", "text", "Y", "alfabeto"},
		Slide{"Letra Z", "text", "Z", "alfabeto"},

		//números
		Slide{"Número 0", "text", "0", "numero"},
		Slide{"Número 1", "text", "1", "numero"},
		Slide{"Número 2", "text", "2", "numero"},
		Slide{"Número 3", "text", "3", "numero"},
		Slide{"Número 4", "text", "4", "numero"},
		Slide{"Número 5", "text", "5", "numero"},
		Slide{"Número 6", "text", "6", "numero"},
		Slide{"Número 7", "text", "7", "numero"},
		Slide{"Número 8", "text", "8", "numero"},
		Slide{"Número 9", "text", "9", "numero"},
		Slide{"Número 10", "text", "10", "numero"},
		Slide{"Número 11", "text", "11", "numero"},
		Slide{"Número 12", "text", "12", "numero"},
		Slide{"Número 13", "text", "13", "numero"},
		Slide{"Número 14", "text", "14", "numero"},
		Slide{"Número 15", "text", "15", "numero"},
		Slide{"Número 16", "text", "16", "numero"},
		Slide{"Número 17", "text", "17", "numero"},
		Slide{"Número 18", "text", "18", "numero"},
		Slide{"Número 19", "text", "19", "numero"},
		Slide{"Número 20", "text", "20", "numero"},
		Slide{"Número 21", "text", "21", "numero"},
		Slide{"Número 22", "text", "22", "numero"},
		Slide{"Número 23", "text", "23", "numero"},
		Slide{"Número 24", "text", "24", "numero"},
		Slide{"Número 25", "text", "25", "numero"},
		Slide{"Número 26", "text", "26", "numero"},
		Slide{"Número 27", "text", "27", "numero"},
		Slide{"Número 28", "text", "28", "numero"},
		Slide{"Número 29", "text", "29", "numero"},
		Slide{"Número 30", "text", "30", "numero"},
	}


	images := []Image{
		//bancos
		Image{"Banco do Brasil", "fixture/images/bancos/banco-do-brasil.png", "bancos"},
		Image{"Bradesco", "fixture/images/bancos/bradesco.png", "bancos"},
		Image{"Caixa", "fixture/images/bancos/caixa.png", "bancos"},
		Image{"Santander", "fixture/images/bancos/santander.png", "bancos"},

		//desenho
		Image{"Carros", "fixture/images/desenho/carros.png", "desenho"},
		Image{"Peppa Pig", "fixture/images/desenho/peppa.png", "desenho"},
		Image{"Show da Luna", "fixture/images/desenho/show-da-luna.png", "desenho"},
		Image{"Thomas e seus amigos", "fixture/images/desenho/thomas-e-seus-amigos.png", "desenho"},

		//refrigerante
		Image{"Coca-Cola", "fixture/images/guarana/coca-cola.png", "refrigerante"},
		Image{"Fanta", "fixture/images/guarana/fanta.png", "refrigerante"},
		Image{"Fanta Guaraná", "fixture/images/guarana/fanta-guarana.png", "refrigerante"},
		Image{"Guaraná Antartica", "fixture/images/guarana/guarana-antartica.png", "refrigerante"},
		Image{"Pepsi", "fixture/images/guarana/pepsi.png", "refrigerante"},

		//geral
		Image{"Boticário", "fixture/images/marcas/boticario.png", "geral"},
		Image{"Piracanjuba", "fixture/images/marcas/piracanjuba.png", "geral"},
		Image{"Piraquê", "fixture/images/marcas/piraque.png", "geral"},
		Image{"Play Doh", "fixture/images/marcas/play-doh.png", "geral"},
		Image{"Ri Happy", "fixture/images/marcas/ri-happy.png", "geral"},

		//futebol
		Image{"Corinthians", "fixture/images/times/corinthians.png", "futebol"},
		Image{"Brasil", "fixture/images/times/brasil.png", "futebol"},

		//globo
		Image{"Bom dia Brasil", "fixture/images/globo/bom-dia-brasil.png", "globo"},
		Image{"Bom dia São Paulo", "fixture/images/globo/bom-dia-sp.png", "globo"},
		Image{"Caldeirão do Huck", "fixture/images/globo/caldeirao-do-huck.png", "globo"},
		Image{"Como Será", "fixture/images/globo/como-sera.png", "globo"},
		Image{"É de casa", "fixture/images/globo/e-de-casa.png", "globo"},
		Image{"Faustão", "fixture/images/globo/faustao.png", "globo"},
		Image{"Faustão Ding Dong", "fixture/images/globo/faustao-ding-dong.png", "globo"},
		Image{"Globo Esporte", "fixture/images/globo/globo-esporte.png", "globo"},
		Image{"Jornal Hoje", "fixture/images/globo/jornal-hoje.png", "globo"},
		Image{"Jornal Nacional", "fixture/images/globo/jornal-nacional.png", "globo"},
		Image{"Mais Você", "fixture/images/globo/mais-voce.png", "globo"},
		Image{"Novo Mundo", "fixture/images/globo/novo-mundo.png", "globo"},
		Image{"Os dias eram assim", "fixture/images/globo/os-dias-eram-assim.png", "globo"},
		Image{"Pega-Pega", "fixture/images/globo/pega-pega.png", "globo"},
		Image{"Pequenas empresas grandes negócios", "fixture/images/globo/pequenas-empresas-grandes-negocios.png", "globo"},
		Image{"SPTV 1", "fixture/images/globo/sptv-1.png", "globo"},
		Image{"SPTV 2", "fixture/images/globo/sptv-2.png", "globo"},
		Image{"Video Show", "fixture/images/globo/video-show.png", "globo"},
	}

	for _, image := range images {
		base64 := util.ImageToBase64(image.Filename)

		slide := Slide{image.Name, "image", base64, image.Group}
		slides = append(slides, slide)
	}

	session, _ := mgo.Dial("192.168.33.10")
	defer session.Close()

	c := session.DB("fiap-que-isso").C("slide")

	for _, slide := range slides {

		result := Slide{}
		c.Find(bson.M{"name": slide.Name, "group": slide.Group}).One(&result)

		if (Slide{}) != result {
			continue
		}

		fmt.Println("Criando slide", slide.Name, "grupo", slide.Group)
		c.Insert(&slide)
	}
}


