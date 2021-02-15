package hat

import (
    "github.com/nunnatsa/piDraw/datatype"
    "log"
    "os"
)


func NewHatMock(e chan<- datatype.HatEvent, s <-chan *datatype.DisplayMessage) {
    go func() {
        for range s {
            log.Println("[HAT Mock] Got display event")
        }
    }()

    go func() {
        var c []byte = make([]byte, 1)
        for {
            _, err := os.Stdin.Read(c)
            if err != nil {
                log.Println("[HAT Mock] Error getting user input;", err)
                continue
            }
            switch c[0] {
            case 'w', 'W':
                e <- datatype.MoveUp
            case 'a', 'A':
                e <- datatype.MoveLeft
            case 'z', 'Z':
                e <- datatype.MoveDown
            case 's', 'S':
                e <- datatype.Pressed
            case 'd', 'D':
                e <- datatype.MoveRight
            }
        }
    }()
}
