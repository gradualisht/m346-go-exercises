package main

import "fmt"

func main() {
	// map[modulnummer]bezeichnung
	modules := map[int]string{
		104: "Einführung in die Informatik",
		117: "Netzwerke einfach erklären",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// löschen
	delete(modules, 117)

	// hinzufügen
	modules[254] = "Geschäftsprozesse modellieren"

	// ersetzen
	modules[104] = "Grundlagen Informatik (aktualisiert)"

	fmt.Println(modules)
}
