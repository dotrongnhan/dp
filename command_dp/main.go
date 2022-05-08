package main

import "command/command"

func main()  {
	tv := &command.TV{}

	onCommand := &command.OnCommand{
		Device: tv,
	}

	//offCommand := &command.OffCommand{
	//	Device: tv,
	//}

	onButton := &command.Button{
		Command: onCommand,
	}

	onButton.Press()
}
