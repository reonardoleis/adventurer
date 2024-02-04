package entities

type World struct {
	Name        string
	Description string
}

func (w *World) JSON() []byte {
	return []byte(`{"name":"` + w.Name + `","description":"` + w.Description + `"}`)
}

func (w World) IsValid() bool {
	return w.Name != "" && w.Description != ""
}
