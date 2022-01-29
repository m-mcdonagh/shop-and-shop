package shell

type CMD struct {
	Help func(...string) string
	Run  func(...string) string
}
