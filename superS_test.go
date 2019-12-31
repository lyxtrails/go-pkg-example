package example

import (
  "testing"
)

var testSuperS SuperS

func setupSuperSTest() {}

func subtestNewSuperS1(t *testing.T) {
  var err error
  testSuperS, err = NewSuperS(SuperSConfig{
    Attr: "attribute",
  })
  if err != nil {
    t.Errorf("Test subtestNewSuperS failed: Receive error %v", err)
  }
}

func subtestNewSuperS2(t *testing.T) {
  tempSuperS, err := NewSuperS(SuperSConfig{})
  testSuperS = tempSuperS.(*superS)
  if err != nil {
    t.Errorf("Test subtestNewSuperS2 failed: Receive error %v", err)
  }
}

func subtestSuperSAttr1(t *testing.T) {
  testNewVal := "NewAttribute"
  var err error
  err = testSuperS.SetAttr(testNewVal)
  if err != nil {
    t.Errorf("Test subtestSuperSAttr1 failed: Setting Attr received error %v", err)
  }
  var newVal string
  newVal, err = testSuperS.GetAttr()
  if err != nil {
    t.Errorf("Test subtestSuperSAttr1 failed: Getting Attr received error %v", err)
  }
  if newVal != testNewVal {
    t.Errorf("Test subtestSuperSAttr1 failed: Getting Attr expected value %v, actual %v", testNewVal, newVal)
  }
}

func cleanupSuperSTest() {}
