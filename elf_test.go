package elf

import "testing"

func TestSqrt(t *testing.T) {
    var world World
    world = World{}
    world.Init(10, 10, 10)
    var tile = &world.Tiles[9][9][9]
    if tile == nil {
        t.Errorf("Shouldn't be blank %v", tile)
    }
}
