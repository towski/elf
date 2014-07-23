package elf

type Tile struct {
    X, Y    int
    Icon rune
}

type World struct {
    Tiles [][][]Tile
}

func(w *World) Init(x int, y int, z int) () {
    w.Tiles = make([][][]Tile, x)
    for i := 0; i < x; i++ {
        w.Tiles[i] = make([][]Tile, y)
        for j := 0; j < y; j++ {
            w.Tiles[i][j] = make([]Tile, z)
            for k := 0; k < z; k++ {
                w.Tiles[i][j][k] = Tile{}
            }
        }
    }
}
