package files

import( 
	"github.com/fmartintbx/godesde0/ejercicios"
	"os"
	//"bufio"
	//"ioutil"
	"fmt"
)

func GrabaTabla(){ 
	var texto string = ejercicios.TabladeMultiplicar()
	archivo, err := os.Create("./files/txt/tabla.txt")
	if err != nil { 
		fmt.Println("Error al crear el archivo "+err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla(){ 
	var texto string = ejercicios.TabladeMultiplicar()
}