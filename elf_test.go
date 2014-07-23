package elf

import "testing"

func TestInit(t *testing.T) {
    var world World
    world = World{}
    world.Init(10, 10, 10)
    var tile = &world.Tiles[9][9][9]
    if tile == nil {
        t.Errorf("Shouldn't be blank %v", tile)
    }
}

func TestAddElf(t *testing.T) {
    var world World
    world = World{}
    world.Init(10, 10, 10)
    t.Logf("hey")
    world.AddElf(1,1,0)
    var tile = world.Tiles[1][1][0]
    t.Logf("hey %v", tile)
    if tile.Units == nil {
        t.Errorf("Shouldn't be blank %v", tile.Units)
    }
}

func TestAddTree(t *testing.T) {
    var world World
    world = World{}
    world.Init(10, 10, 10)
    t.Logf("hey")
    world.AddTree(1,1,0)
}
