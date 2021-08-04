package common

import (
	"fmt"
)

type Services interface {
	SetName(name string)
	GetName() string
	SetCategary(ca string)
	GetCategary() string
	SetSLA(sla float32)
	GetSLA() float32
}

type Service struct {
	Name     string
	SLA      float32
	Categary string
}

func NewServiceManager(name string, sla float32, catagary string) Services {
	return &Service{
		Name:     name,
		SLA:      sla,
		Categary: catagary,
	}

}

func (s *Service) SetName(name string) {
	s.Name = name
}

func (s *Service) GetName() string {
	return s.Name
}

func (s *Service) SetCategary(ca string) {
	s.Categary = ca
}

func (s *Service) GetCategary() string {
	return s.Categary
}

func (s *Service) SetSLA(sla float32) {
	s.SLA = sla
}

func (s *Service) GetSLA() float32 {
	return s.SLA
}

func (s Service) HealthCheck(y, mon int, dt float32) float32 {
	bm := []int{1, 3, 5, 7, 8, 10, 12}
	sm := []int{4, 6, 9, 11}
	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		if mon == 2 {
			s.SLA = 1 - float32(dt/(29*24*60))
		} else if mon > 0 && mon < 13 {
			for _, a1 := range bm {
				if mon == a1 {
					s.SLA = 1 - float32(dt/(31*24*60))
					fmt.Println(s.SLA)
				}
			}
			for _, a2 := range sm {
				if mon == a2 {
					s.SLA = 1 - float32(dt/(30*24*60))
					fmt.Println(s.SLA)
				}
			}
		} else {
			fmt.Println("You input Month isn't right,please check...")
		}
	} else {
		if mon == 2 {
			s.SLA = 1 - float32(dt/(28*24*60))
		} else if mon > 0 && mon < 13 {
			for _, a1 := range bm {
				if mon == a1 {
					s.SLA = 1 - float32(dt/(31*24*60))
					// fmt.Println(s.SLA)
				}
			}
			for _, a2 := range sm {
				if mon == a2 {
					s.SLA = 1 - float32(dt/(30*24*60))
					// fmt.Println(s.SLA)
				}
			}
		} else {
			fmt.Println("You input Month isn't right,please check...")
		}
	}
	return s.SLA
}
