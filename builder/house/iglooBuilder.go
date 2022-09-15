package main

type IglooBuilder struct {
    windowType string
    doorType   string
    floor      int
}

func newIglooBuilder() *IglooBuilder {
    return &IglooBuilder{}
}

////////////////////////////// 实现 iBuilder 接口中的方法

func (b *IglooBuilder) setWindowType() {
    b.windowType = "Snow Window"
}

func (b *IglooBuilder) setDoorType() {
    b.doorType = "Snow Door"
}

func (b *IglooBuilder) setNumFloor() {
    b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
    return House{
        doorType:   b.doorType,
        windowType: b.windowType,
        floor:      b.floor,
    }
}