package example

type superS struct {
  attr string
}

type SuperSConfig struct {
  Attr string
}

func NewSuperS(cfg SuperSConfig) (SuperS, error) {
  if cfg.Attr == "" {
    cfg.Attr = DEFAULTATTR
  }
  return &superS{cfg.Attr}, nil
}

func (s *superS) GetAttr() (string, error) {
  return s.attr, nil
}

func (s *superS) SetAttr(input string) error {
  s.attr = input
  return nil
}
