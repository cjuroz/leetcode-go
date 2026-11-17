package main

import (
	"errors"
	"fmt"
)

// --- Domain Types & Enums ---

type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
	Van
)

func (v VehicleType) String() string {
	switch v {
	case Motorcycle:
		return "Motorcycle"
	case Car:
		return "Car"
	case Van:
		return "Van"
	default:
		return "Unknown"
	}
}

type SpotType int

const (
	MotorcycleSpot SpotType = iota
	CompactSpot
	RegularSpot // Also serves as Large spot for vans (3 contiguous required)
)

func (s SpotType) String() string {
	switch s {
	case MotorcycleSpot:
		return "Motorcycle"
	case CompactSpot:
		return "Compact"
	case RegularSpot:
		return "Regular"
	default:
		return "Unknwon"
	}
}

// --- Vehicle Abstraction (OOP Interface) ---

type Vehicle interface {
	LicensePlate() string
	Type() VehicleType
	SpotsNeeded() int
	CanFitInSpot(spotType SpotType) bool
}

type BaseVehicle struct {
	plate       string
	vehicleType VehicleType
	spotsNeeded int
}

func (b BaseVehicle) LicensePlate() string { return b.plate }
func (b BaseVehicle) Type() VehicleType    { return b.vehicleType }
func (b BaseVehicle) SpotsNeeded() int     { return b.spotsNeeded }

type MotorcycleVehicle struct{ BaseVehicle }

func NewMotorcycle(plate string) *MotorcycleVehicle {
	return &MotorcycleVehicle{BaseVehicle{plate: plate, vehicleType: Motorcycle, spotsNeeded: 1}}
}

func (m *MotorcycleVehicle) CanFitInSpot(spotType SpotType) bool {
	// A motorcycle can park in any spot
	return true
}

type CarVehicle struct{ BaseVehicle }

func NewCar(plate string) *CarVehicle {
	return &CarVehicle{BaseVehicle{plate: plate, vehicleType: Car, spotsNeeded: 1}}
}

func (c *CarVehicle) CanFitInSpot(spotType SpotType) bool {
	// A car can park in a single compact spot or a regular spot
	return spotType == CompactSpot || spotType == RegularSpot
}

type VanVehicle struct{ BaseVehicle }

func NewVan(plate string) *VanVehicle {
	return &VanVehicle{BaseVehicle{plate: plate, vehicleType: Van, spotsNeeded: 3}}
}

func (v *VanVehicle) CanFitInSpot(spotType SpotType) bool {
	return spotType == RegularSpot
}

// --- Parking Spot ---

type ParkingSpot struct {
	id       int
	spotType SpotType
	vehicle  Vehicle
}

func (s *ParkingSpot) IsAvailable() bool {
	return s.vehicle == nil
}

func (s *ParkingSpot) Park(v Vehicle) {
	s.vehicle = v
}

func (s *ParkingSpot) RemoveVehicle() {
	s.vehicle = nil
}

// --- Parking Lot ---

type ParkingLot struct {
	spots          []*ParkingSpot
	parkedVehicles map[string][]*ParkingSpot // maps license plate -> assigned spots
	occupiedSpots  int
	vanSpotsTaken  int
}

func NewParkingLot(motorcycleCount, compactCount, regularcount int) *ParkingLot {
	total := motorcycleCount + compactCount + regularcount
	spots := make([]*ParkingSpot, 0, total)

	id := 1
	for i := 0; i < motorcycleCount; i++ {
		spots = append(spots, &ParkingSpot{id: id, spotType: MotorcycleSpot})
		id++
	}

	for i := 0; i < compactCount; i++ {
		spots = append(spots, &ParkingSpot{id: id, spotType: CompactSpot})
		id++
	}

	for i := 0; i < regularcount; i++ {
		spots = append(spots, &ParkingSpot{id: id, spotType: RegularSpot})
		id++
	}

	return &ParkingLot{
		spots:          spots,
		parkedVehicles: make(map[string][]*ParkingSpot),
	}
}

// ParkVehicle finds the best available spot(s) and parks the vehicle.
// Strategy: Prefer smallest valid spot first so larger spots stay open for vans/cars.
func (p *ParkingLot) ParkVehicle(v Vehicle) error {
	if _, exists := p.parkedVehicles[v.LicensePlate()]; exists {
		return errors.New("vehicle is already parked in the lot")
	}

	needed := v.SpotsNeeded()

	// Check spot types in order of smallest to largest (Motorcycle -> Compact -> Regular)
	preferredOrder := []SpotType{MotorcycleSpot, CompactSpot, RegularSpot}

	for _, targetType := range preferredOrder {
		if !v.CanFitInSpot(targetType) {
			continue
		}

		// Look for 'needed' contiguous available spots of targetType
		for i := 0; i <= len(p.spots)-needed; i++ {
			canFit := true
			for j := 0; j < needed; j++ {
				spot := p.spots[i+j]
				if !spot.IsAvailable() || spot.spotType != targetType {
					canFit = false
					break
				}
			}

			if canFit {
				assigned := make([]*ParkingSpot, 0, needed)
				for j := 0; j < needed; j++ {
					p.spots[i+j].Park(v)
					assigned = append(assigned, p.spots[i+j])
				}
				p.parkedVehicles[v.LicensePlate()] = assigned
				if v.Type() == Van {
					p.vanSpotsTaken += needed
				}
				p.occupiedSpots += needed

				return nil
			}
		}
	}

	return fmt.Errorf("no available spots for %s (%s)", v.Type(), v.LicensePlate())
}

// RemoveVehicle frees the spots occupied by a vehicle.
func (p *ParkingLot) RemoveVehicle(plate string) error {
	assigned, exists := p.parkedVehicles[plate]
	if !exists {
		return fmt.Errorf("vehicle %s not found in parking lot", plate)
	}

	if len(assigned) > 0 && assigned[0].vehicle.Type() == Van {
		p.vanSpotsTaken -= len(assigned)
	}

	for _, spot := range assigned {
		spot.RemoveVehicle()
	}
	delete(p.parkedVehicles, plate)
	p.occupiedSpots -= len(assigned)

	return nil
}

// --- Required Query Methods ---

// TotalSpots returns the total capacity of the parking lot.
func (p *ParkingLot) TotalSpots() int {
	return len(p.spots)
}

// RemainingSpots returns how many spots are currently free.
func (p *ParkingLot) RemainingSpots() int {
	return len(p.spots) - p.occupiedSpots
}

// IsFull returns true if zero spots are available.
func (p *ParkingLot) IsFull() bool {
	return len(p.spots) == p.occupiedSpots
}

// IsEmpty returns true if all spots are available.
func (p *ParkingLot) IsEmpty() bool {
	return len(p.parkedVehicles) == 0
}

// AreSpotsFull returns true if all spots of a specific SpotType are taken.
func (p *ParkingLot) AreSpotsFull(spotType SpotType) bool {
	for _, s := range p.spots {
		if s.spotType == spotType && s.IsAvailable() {
			return false
		}
	}

	return true
}

// VanSpotsCount returns how many spots vans are currently taking up.
func (p *ParkingLot) VanSpotsCount() int {
	return p.vanSpotsTaken
}

// --- Driver / Test Execution ---

func main() {
	// Initialize lot: 2 motorcycle spots, 3 compact spots, 6 regular spots (11 total)
	lot := NewParkingLot(2, 3, 6)

	fmt.Printf("Total spots: %d\n", lot.TotalSpots())
	fmt.Printf("Is empty? %v\n\n", lot.IsEmpty())

	// Park 2 motorcycles (fills motorcycle spots)
	lot.ParkVehicle(NewMotorcycle("MOTO-1"))
	lot.ParkVehicle(NewMotorcycle("MOTO-2"))
	fmt.Printf("Motorcycle spots full? %v\n", lot.AreSpotsFull(MotorcycleSpot))

	// Park a 3rd motorcycle (spills over into a compact spot)
	lot.ParkVehicle(NewMotorcycle("MOTO-3"))
	fmt.Printf("Remaining spots after 3 motorcycles: %d\n", lot.RemainingSpots())

	// Park 2 vans (takes 3 regular spots each = 6 spots total)
	lot.ParkVehicle(NewVan("VAN-1"))
	lot.ParkVehicle(NewVan("VAN-2"))
	fmt.Printf("Spots taken by vans: %d\n", lot.VanSpotsCount())
	fmt.Printf("Regular spots full? %v\n", lot.AreSpotsFull(RegularSpot))

	// Try to park a 3rd van (should fail)
	if err := lot.ParkVehicle(NewVan("VAN-3")); err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Fill remaining compact spots with cars
	lot.ParkVehicle(NewCar("CAR-1"))
	lot.ParkVehicle(NewCar("CAR-2"))

	fmt.Printf("\nRemaining spots: %d\n", lot.RemainingSpots())
	fmt.Printf("Is lot full? %v\n", lot.IsFull())

	// Remove a van and check updated counts
	lot.RemoveVehicle("VAN-1")
	fmt.Printf("\nAfter VAN-1 leaves -> Spots taken by vans: %d, Remaining spots: %d\n",
		lot.VanSpotsCount(), lot.RemainingSpots())
}
