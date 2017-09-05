package main

import (
	"fmt"
	"github.com/dsaouda/fiap-que-isso/admin/models"
	"gopkg.in/mgo.v2/bson"
	"crypto/md5"
	"encoding/hex"
	"github.com/dsaouda/fiap-que-isso/admin/db"
)

type Image struct {
	Name string
	Filename string
	Group string
}

func main() {

	logins:= []models.Login{
		models.NewLogin("dsaouda@gmail.com", "dsaouda"),
		models.NewLogin("diegosaouda@gmail.com", "dsaouda"),
		models.NewLogin("admin@admin", "dsaouda"),
		models.NewLogin("root@root", "dsaouda"),
	}

	slides := []models.Slide{
		//alfabeto
		models.Slide{Name: "Letra A", Type: "text", Value: "A", Group: "alfabeto", Order: 1 },
		models.Slide{Name: "Letra B", Type: "text", Value: "B", Group: "alfabeto", Order: 2 },
		models.Slide{Name: "Letra C", Type: "text", Value: "C", Group: "alfabeto", Order: 3 },
		models.Slide{Name: "Letra D", Type: "text", Value: "D", Group: "alfabeto", Order: 4 },
		models.Slide{Name: "Letra E", Type: "text", Value: "E", Group: "alfabeto", Order: 5 },
		models.Slide{Name: "Letra F", Type: "text", Value: "F", Group: "alfabeto", Order: 6 },
		models.Slide{Name: "Letra G", Type: "text", Value: "G", Group: "alfabeto", Order: 7 },
		models.Slide{Name: "Letra H", Type: "text", Value: "H", Group: "alfabeto", Order: 8 },
		models.Slide{Name: "Letra I", Type: "text", Value: "I", Group: "alfabeto", Order: 9 },
		models.Slide{Name: "Letra J", Type: "text", Value: "J", Group: "alfabeto", Order: 10},
		models.Slide{Name: "Letra K", Type: "text", Value: "K", Group: "alfabeto", Order: 11},
		models.Slide{Name: "Letra L", Type: "text", Value: "L", Group: "alfabeto", Order: 12},
		models.Slide{Name: "Letra M", Type: "text", Value: "M", Group: "alfabeto", Order: 13},
		models.Slide{Name: "Letra N", Type: "text", Value: "N", Group: "alfabeto", Order: 14},
		models.Slide{Name: "Letra O", Type: "text", Value: "O", Group: "alfabeto", Order: 15},
		models.Slide{Name: "Letra P", Type: "text", Value: "P", Group: "alfabeto", Order: 16},
		models.Slide{Name: "Letra Q", Type: "text", Value: "Q", Group: "alfabeto", Order: 17},
		models.Slide{Name: "Letra R", Type: "text", Value: "R", Group: "alfabeto", Order: 18},
		models.Slide{Name: "Letra S", Type: "text", Value: "S", Group: "alfabeto", Order: 19},
		models.Slide{Name: "Letra T", Type: "text", Value: "T", Group: "alfabeto", Order: 20},
		models.Slide{Name: "Letra U", Type: "text", Value: "U", Group: "alfabeto", Order: 21},
		models.Slide{Name: "Letra V", Type: "text", Value: "V", Group: "alfabeto", Order: 22},
		models.Slide{Name: "Letra W", Type: "text", Value: "W", Group: "alfabeto", Order: 23},
		models.Slide{Name: "Letra X", Type: "text", Value: "X", Group: "alfabeto", Order: 24},
		models.Slide{Name: "Letra Y", Type: "text", Value: "Y", Group: "alfabeto", Order: 25},
		models.Slide{Name: "Letra Z", Type: "text", Value: "Z", Group: "alfabeto", Order: 26},

		//números
		models.Slide{Name:"Número 0",  Type:"text", Value:"0",   Group: "numero", Order: 1 },
		models.Slide{Name:"Número 1",  Type:"text", Value:"1",   Group: "numero", Order: 2 },
		models.Slide{Name:"Número 2",  Type:"text", Value:"2",   Group: "numero", Order: 3 },
		models.Slide{Name:"Número 3",  Type:"text", Value:"3",   Group: "numero", Order: 4 },
		models.Slide{Name:"Número 4",  Type:"text", Value:"4",   Group: "numero", Order: 5 },
		models.Slide{Name:"Número 5",  Type:"text", Value:"5",   Group: "numero", Order: 6 },
		models.Slide{Name:"Número 6",  Type:"text", Value:"6",   Group: "numero", Order: 7 },
		models.Slide{Name:"Número 7",  Type:"text", Value:"7",   Group: "numero", Order: 8 },
		models.Slide{Name:"Número 8",  Type:"text", Value:"8",   Group: "numero", Order: 9 },
		models.Slide{Name:"Número 9",  Type:"text", Value:"9",   Group: "numero", Order: 10},
		models.Slide{Name:"Número 10", Type: "text",Value: "10", Group: "numero", Order: 11},
		models.Slide{Name:"Número 11", Type: "text",Value: "11", Group: "numero", Order: 12},
		models.Slide{Name:"Número 12", Type: "text",Value: "12", Group: "numero", Order: 13},
		models.Slide{Name:"Número 13", Type: "text",Value: "13", Group: "numero", Order: 14},
		models.Slide{Name:"Número 14", Type: "text",Value: "14", Group: "numero", Order: 15},
		models.Slide{Name:"Número 15", Type: "text",Value: "15", Group: "numero", Order: 16},
		models.Slide{Name:"Número 16", Type: "text",Value: "16", Group: "numero", Order: 17},
		models.Slide{Name:"Número 17", Type: "text",Value: "17", Group: "numero", Order: 18},
		models.Slide{Name:"Número 18", Type: "text",Value: "18", Group: "numero", Order: 19},
		models.Slide{Name:"Número 19", Type: "text",Value: "19", Group: "numero", Order: 20},
		models.Slide{Name:"Número 20", Type: "text",Value: "20", Group: "numero", Order: 21},
		models.Slide{Name:"Número 21", Type: "text",Value: "21", Group: "numero", Order: 22},
		models.Slide{Name:"Número 22", Type: "text",Value: "22", Group: "numero", Order: 23},
		models.Slide{Name:"Número 23", Type: "text",Value: "23", Group: "numero", Order: 24},
		models.Slide{Name:"Número 24", Type: "text",Value: "24", Group: "numero", Order: 25},
		models.Slide{Name:"Número 25", Type: "text",Value: "25", Group: "numero", Order: 26},
		models.Slide{Name:"Número 26", Type: "text",Value: "26", Group: "numero", Order: 27},
		models.Slide{Name:"Número 27", Type: "text",Value: "27", Group: "numero", Order: 28},
		models.Slide{Name:"Número 28", Type: "text",Value: "28", Group: "numero", Order: 29},
		models.Slide{Name:"Número 29", Type: "text",Value: "29", Group: "numero", Order: 30},
		models.Slide{Name:"Número 30", Type: "text",Value: "30", Group: "numero", Order: 31},
	}

	images := []Image{
		//bancos
		Image{"Banco do Brasil", "bancos/banco-do-brasil.png", "bancos"},
		Image{"Bradesco", "bancos/bradesco.png", "bancos"},
		Image{"Caixa", "bancos/caixa.png", "bancos"},
		Image{"Santander", "bancos/santander.png", "bancos"},

		//desenho
		Image{"Carros", "desenho/carros.png", "desenho"},
		Image{"Peppa Pig", "desenho/peppa.png", "desenho"},
		Image{"Show da Luna", "desenho/show-da-luna.png", "desenho"},
		Image{"Thomas e seus amigos", "desenho/thomas-e-seus-amigos.png", "desenho"},

		//refrigerante
		Image{"Coca-Cola", "guarana/coca-cola.png", "refrigerante"},
		Image{"Fanta", "guarana/fanta.png", "refrigerante"},
		Image{"Fanta Guaraná", "guarana/fanta-guarana.png", "refrigerante"},
		Image{"Guaraná Antartica", "guarana/guarana-antartica.png", "refrigerante"},
		Image{"Pepsi", "guarana/pepsi.png", "refrigerante"},

		//geral
		Image{"Boticário", "marcas/boticario.png", "geral"},
		Image{"Piracanjuba", "marcas/piracanjuba.png", "geral"},
		Image{"Piraquê", "marcas/piraque.png", "geral"},
		Image{"Play Doh", "marcas/play-doh.png", "geral"},
		Image{"Ri Happy", "marcas/ri-happy.png", "geral"},

		//futebol
		Image{"Corinthians", "times/corinthians.png", "futebol"},
		Image{"Brasil", "times/brasil.png", "futebol"},

		//globo
		Image{"Bom dia Brasil", "globo/bom-dia-brasil.png", "globo"},
		Image{"Bom dia São Paulo", "globo/bom-dia-sp.png", "globo"},
		Image{"Caldeirão do Huck", "globo/caldeirao-do-huck.png", "globo"},
		Image{"Como Será", "globo/como-sera.png", "globo"},
		Image{"É de casa", "globo/e-de-casa.png", "globo"},
		Image{"Faustão", "globo/faustao.png", "globo"},
		Image{"Faustão Ding Dong", "globo/faustao-ding-dong.png", "globo"},
		Image{"Globo Esporte", "globo/globo-esporte.png", "globo"},
		Image{"Jornal Hoje", "globo/jornal-hoje.png", "globo"},
		Image{"Jornal Nacional", "globo/jornal-nacional.png", "globo"},
		Image{"Mais Você", "globo/mais-voce.png", "globo"},
		Image{"Novo Mundo", "globo/novo-mundo.png", "globo"},
		Image{"Os dias eram assim", "globo/os-dias-eram-assim.png", "globo"},
		Image{"Pega-Pega", "globo/pega-pega.png", "globo"},
		Image{"Pequenas empresas grandes negócios", "globo/pequenas-empresas-grandes-negocios.png", "globo"},
		Image{"SPTV 1", "globo/sptv-1.png", "globo"},
		Image{"SPTV 2", "globo/sptv-2.png", "globo"},
		Image{"Video Show", "globo/video-show.png", "globo"},
	}

	for key, image := range images {
		slide := models.Slide{Name: image.Name, Type: "image", Value: "/images/" + image.Filename, Group: image.Group, Order: key}
		slides = append(slides, slide)
	}

	db.Connect()
	defer db.Session.Close()



	db := db.Session.DB(db.Mongo.Database)

	cSlide := db.C("slide")
	cLogin := db.C("login")

	for _, slide := range slides {
		result := models.Slide{}
		cSlide.Find(bson.M{"name": slide.Name, "group": slide.Group}).One(&result)

		if (models.Slide{}) != result {
			continue
		}

		fmt.Println("Criando slide", slide.Name, "grupo", slide.Group)
		cSlide.Insert(&slide)
	}

	for _, login := range logins {
			result := models.Login{}
			cLogin.Find(bson.M{"email": login.Email}).One(&result)

			if (models.Login{}) != result {
				continue
			}

		hasher := md5.New()
		hasher.Write([]byte(login.Email + login.Password))
		login.Token = hex.EncodeToString(hasher.Sum(nil))

		fmt.Println("Criando login", login.Email)
		cLogin.Insert(&login)
	}
}


