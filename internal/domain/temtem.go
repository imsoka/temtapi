package domain

type Temtem struct {
    Id string;
    Name string;
    TempediaNumber uint;
    Stats map[string]uint8;
    TempediaEntry string;
}
