package example

type s1 struct {
  SuperS
}

func NewS1() (S1, error) {
  s, err := NewSuperS(SuperSConfig{})
  return &s1{s}, err
}

func (s *s1) GetAttr() (string, error) {
  superSAttr, err := s.SuperS.GetAttr()
  return "s1_" + superSAttr, err
}
