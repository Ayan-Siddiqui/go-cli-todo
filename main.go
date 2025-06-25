package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	cf := NewCommandFlags()
	storage.Load(&todos)
	cf.Execute(&todos)
	storage.Save(&todos)
}
