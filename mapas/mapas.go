package mapas

import "fmt"

func MostrarMapas(){ 
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{ 
		"Barcelona": 39,
		"Real Madrid": 38,
		"Chivas": 37,
		"River Plate": 30,
	}
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato{ 
		fmt.Printf("Equipo %s, tiene un punttaje de %d ", equipo, puntaje)
	}
	delete(campeonato, "Real Madrid")

}