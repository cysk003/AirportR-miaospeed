package macros

import (
	"github.com/airportr/miaospeed/interfaces"
	"github.com/airportr/miaospeed/utils/structs"

	"github.com/airportr/miaospeed/service/macros/geo"
	"github.com/airportr/miaospeed/service/macros/invalid"
	"github.com/airportr/miaospeed/service/macros/ping"
	"github.com/airportr/miaospeed/service/macros/script"
	"github.com/airportr/miaospeed/service/macros/sleep"
	"github.com/airportr/miaospeed/service/macros/speed"
	"github.com/airportr/miaospeed/service/macros/udp"
)

var registeredList = map[interfaces.SlaveRequestMacroType]func() interfaces.SlaveRequestMacro{
	interfaces.MacroSpeed: func() interfaces.SlaveRequestMacro {
		return &speed.Speed{}
	},
	interfaces.MacroPing: func() interfaces.SlaveRequestMacro {
		return &ping.Ping{}
	},
	interfaces.MacroUDP: func() interfaces.SlaveRequestMacro {
		return &udp.Udp{}
	},
	interfaces.MacroGeo: func() interfaces.SlaveRequestMacro {
		return &geo.Geo{}
	},
	interfaces.MacroScript: func() interfaces.SlaveRequestMacro { return &script.Script{} },
	interfaces.MacroSleep:  func() interfaces.SlaveRequestMacro { return &sleep.Sleep{} },
}

func Find(macroType interfaces.SlaveRequestMacroType) interfaces.SlaveRequestMacro {
	if newFn, ok := registeredList[macroType]; ok && newFn != nil {
		return newFn()
	}

	return &invalid.Invalid{}
}

func FindBatch(macroTypes []interfaces.SlaveRequestMacroType) []interfaces.SlaveRequestMacro {
	return structs.Map(macroTypes, func(m interfaces.SlaveRequestMacroType) interfaces.SlaveRequestMacro {
		return Find(m)
	})
}
