package observer

import "fmt"

type Item struct {
	ObserverList []Observer
	Name         string
	InStock      bool
}

func (i *Item) Register(observer Observer) {
	i.ObserverList = append(i.ObserverList, observer)
}

func (i *Item) DeRegister(observer Observer) {
	i.ObserverList = removeFromSlice(i.ObserverList, observer)
}

func (i *Item) NotifyAll() {
	for _, observer := range i.ObserverList {
		observer.Update(i.Name)
	}
}

func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s in stock \n", i.Name)
	i.InStock = true
	i.NotifyAll()
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	lenList := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetId() == observer.GetId() {
			observerList[lenList-1], observerList[i] = observerList[i], observerList[lenList-1]
			return observerList[:lenList-1]
		}
	}
	return observerList
}
