package flyweight

type Factory struct {
	SoldierMap map[string]InterfaceSoldier
}

func NewFactory() *Factory {
	return &Factory{
		SoldierMap: make(map[string]InterfaceSoldier),
	}
}

func (f *Factory) CreateSoldier(name string) InterfaceSoldier {
	if val, ok := f.SoldierMap[name]; ok {
		return val
	}
	f.SoldierMap[name] = NewSoldier(name)
	return f.SoldierMap[name]
}

func (f *Factory) GetSize() int {
	return len(f.SoldierMap)
}
