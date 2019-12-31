package example

import (
  "os"
  "testing"
)

func setup() {}

func teardown() {}

func TestRun(t *testing.T) {
  setupS1Test()
  t.Run("s1/NewS1/1", subtestNewS11)
  t.Run("s1/Attr/1", subtestS1Attr1)
  cleanupS1Test()

  setupSuperSTest()
  t.Run("superS/NewSuperS/1", subtestNewSuperS1)
  t.Run("superS/NewSuperS/2", subtestNewSuperS2)
  t.Run("superS/Attr/1", subtestSuperSAttr1)
  cleanupSuperSTest()
}

func TestMain(m *testing.M) {
  setup()
  tests := m.Run()
  teardown()
  os.Exit(tests)
}
