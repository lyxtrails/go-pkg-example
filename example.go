package example

const (
  DEFAULTATTR = "defaultValue"
)

type S1 interface {
  SuperS
}

type SuperS interface {
  GetAttr() (result string, err error)
  SetAttr(input string) (err error)
}
