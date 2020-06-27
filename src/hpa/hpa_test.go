package main

import "testing"

func TestLooping(t *testing.T){
  resultado := Looping()
  if resultado != "Code.education Rocks!" {
    t.Errorf("O resultado esperado: %s, obtido: %s", "Code.education Rocks!", resultado)
  }
}