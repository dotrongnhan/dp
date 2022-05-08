package command

type Command interface {
	execute()
}

type OnCommand struct {
	Device Device
}

type OffCommand struct {
	Device Device
}

func (o *OnCommand) execute() {
	o.Device.on()
}

func (o *OffCommand) execute() {
	o.Device.off()
}