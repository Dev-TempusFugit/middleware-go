package main

func execute(name string, f func(string)) {
	f(name)
}

func main() {
	name := "Maxim"
	execute(name, funciones.Saludar())
}
