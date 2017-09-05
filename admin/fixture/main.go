package main

import (
	"fmt"
	"github.com/dsaouda/fiap-que-isso/admin/fixture/util"
	"github.com/dsaouda/fiap-que-isso/admin/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"crypto/md5"
	"encoding/hex"
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
		models.Slide{"Letra A", "text", "A", "alfabeto"},
		models.Slide{"Letra B", "text", "B", "alfabeto"},
		models.Slide{"Letra C", "text", "C", "alfabeto"},
		models.Slide{"Letra D", "text", "D", "alfabeto"},
		models.Slide{"Letra E", "text", "E", "alfabeto"},
		models.Slide{"Letra F", "text", "F", "alfabeto"},
		models.Slide{"Letra G", "text", "G", "alfabeto"},
		models.Slide{"Letra H", "text", "H", "alfabeto"},
		models.Slide{"Letra I", "text", "I", "alfabeto"},
		models.Slide{"Letra J", "text", "J", "alfabeto"},
		models.Slide{"Letra K", "text", "K", "alfabeto"},
		models.Slide{"Letra L", "text", "L", "alfabeto"},
		models.Slide{"Letra M", "text", "M", "alfabeto"},
		models.Slide{"Letra N", "text", "N", "alfabeto"},
		models.Slide{"Letra O", "text", "O", "alfabeto"},
		models.Slide{"Letra P", "text", "P", "alfabeto"},
		models.Slide{"Letra Q", "text", "Q", "alfabeto"},
		models.Slide{"Letra R", "text", "R", "alfabeto"},
		models.Slide{"Letra S", "text", "S", "alfabeto"},
		models.Slide{"Letra T", "text", "T", "alfabeto"},
		models.Slide{"Letra U", "text", "U", "alfabeto"},
		models.Slide{"Letra V", "text", "V", "alfabeto"},
		models.Slide{"Letra W", "text", "W", "alfabeto"},
		models.Slide{"Letra X", "text", "X", "alfabeto"},
		models.Slide{"Letra Y", "text", "Y", "alfabeto"},
		models.Slide{"Letra Z", "text", "Z", "alfabeto"},

		//números
		models.Slide{"Número 0", "text", "0", "numero"},
		models.Slide{"Número 1", "text", "1", "numero"},
		models.Slide{"Número 2", "text", "2", "numero"},
		models.Slide{"Número 3", "text", "3", "numero"},
		models.Slide{"Número 4", "text", "4", "numero"},
		models.Slide{"Número 5", "text", "5", "numero"},
		models.Slide{"Número 6", "text", "6", "numero"},
		models.Slide{"Número 7", "text", "7", "numero"},
		models.Slide{"Número 8", "text", "8", "numero"},
		models.Slide{"Número 9", "text", "9", "numero"},
		models.Slide{"Número 10", "text", "10", "numero"},
		models.Slide{"Número 11", "text", "11", "numero"},
		models.Slide{"Número 12", "text", "12", "numero"},
		models.Slide{"Número 13", "text", "13", "numero"},
		models.Slide{"Número 14", "text", "14", "numero"},
		models.Slide{"Número 15", "text", "15", "numero"},
		models.Slide{"Número 16", "text", "16", "numero"},
		models.Slide{"Número 17", "text", "17", "numero"},
		models.Slide{"Número 18", "text", "18", "numero"},
		models.Slide{"Número 19", "text", "19", "numero"},
		models.Slide{"Número 20", "text", "20", "numero"},
		models.Slide{"Número 21", "text", "21", "numero"},
		models.Slide{"Número 22", "text", "22", "numero"},
		models.Slide{"Número 23", "text", "23", "numero"},
		models.Slide{"Número 24", "text", "24", "numero"},
		models.Slide{"Número 25", "text", "25", "numero"},
		models.Slide{"Número 26", "text", "26", "numero"},
		models.Slide{"Número 27", "text", "27", "numero"},
		models.Slide{"Número 28", "text", "28", "numero"},
		models.Slide{"Número 29", "text", "29", "numero"},
		models.Slide{"Número 30", "text", "30", "numero"},
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

		slide := models.Slide{image.Name, "image", base64, image.Group}
		slides = append(slides, slide)
	}

	session, _ := mgo.Dial("192.168.33.10")
	defer session.Close()

	db := session.DB("fiap_que_isso")

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


