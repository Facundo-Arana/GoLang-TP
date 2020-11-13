package main

// Voy a trabajar equipos de futbol

func main() {
	torneo := NewTournament()

	e1 := NewTeam("Boca Juniors", "Russo", 0)
	e1.add(1, "Andrada")
	e1.add(2, "Zambrano")
	e1.add(3, "Fabra")
	e1.add(4, "Buffarini")
	e1.add(5, "Campuzano")

	e2 := NewTeam("Riber", "Gallardo", 1)
	e2.add(1, "Armani")
	e2.add(2, "Pinola")
	e2.add(3, "Casco")
	e2.add(4, "Montiel")
	e2.add(5, "Ponzio")

	torneo.add(e1)
	torneo.add(e2)
	torneo.Print()
}
