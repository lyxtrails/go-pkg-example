package example

import (
  "testing"
)

var testS1 S1

func setupS1Test() {}

func subtestNewS11(t *testing.T) {
  var err error
  testS1, err = NewS1()
  if err != nil {
    t.Errorf("Test subtestNewS1 failed: Receive error %v", err)
  }
}

func subtestS1Attr1(t *testing.T) {
  // Test setter func
  testNewVal := "NewAttribute"
  var err error
  err = testS1.SetAttr(testNewVal)
  if err != nil {
    t.Errorf("Test subtestS1Attr1 failed: Setting Attr received error %v", err)
  }

  // Test overridden getter func in s1
  var newVal string
  newVal, err = testS1.GetAttr()
  if err != nil {
    t.Errorf("Test subtestS1Attr1 failed: Getting Attr received error %v", err)
  }
  if newVal != "s1_"+testNewVal {
    t.Errorf("Test subtestS1Attr1 failed: Getting Attr expected value %v, actual %v", testNewVal, newVal)
  }
}

func cleanupS1Test() {}
