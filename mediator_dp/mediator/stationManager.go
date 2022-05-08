package mediator

import "sync"

type StationManager struct {
	isPlatformFree bool
	lock *sync.Mutex
	TrainQueue []Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		lock: &sync.Mutex{},
	}
}

func (s *StationManager) CanLand(train Train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}

	s.TrainQueue = append(s.TrainQueue,train)
	return false
}

func (s *StationManager) NotifyFree()  {
	s.lock.Lock()
	defer s.lock.Unlock()

	if !s.isPlatformFree {
		s.isPlatformFree = true
	}

	if len(s.TrainQueue) > 0 {
		firstTrainInQueue := s.TrainQueue[0]
		s.TrainQueue = s.TrainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}

