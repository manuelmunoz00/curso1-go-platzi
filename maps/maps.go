package maps

// GetMap
func GetMap() map[string] int  {
	//mapas necesitan ser inicializados
	mapTest := make(map[string] int)
	mapTest["key1"] = 1
	mapTest["key2"] = 2
	delete(mapTest,"key1")
	return mapTest
}

// GetMapaEdades
func GetMapaEdades() map[string] int  {
	mapaEdades := map[string] int{
		"Juan": 34,
		"Diego": 28,
	}
	return mapaEdades
}

// GetEdad
func GetEdad(nombre string) int  {
	mapaEdades := map[string] int{
		"Juan": 34,
		"Diego": 28,
	}
	return mapaEdades[nombre]
}