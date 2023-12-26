package mod

type Movie struct {
	ID       int
	Title    string
	ISBN     string
	Director Director
}
type Director struct {
	Fname string
	Lname string
}
