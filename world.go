package elf

type Tile struct {
    X, Y, Z    int
    Units []Elf
    Icon rune
}

type Formation struct {
}

type Tree struct {
    Formation
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
                w.Tiles[i][j][k] = Tile{i, j, k, make([]Elf, 1), ' '}
            }
        }
    }
}

func(w *World) AddElf(x int, y int, z int) () {
    slice := w.Tiles[x][y][z].Units
    slice = append(slice, Elf{})
}

func(w *World) AddTree(x int, y int, z int) () {
    units := w.Tiles[x][y][z].Units
    units = append(units, Elf{})
}
