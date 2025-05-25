# Elevator System Design

A simple elevator system implementation in Go that simulates multiple elevators serving multiple floors.

## Features

- Multiple elevator support
- Smart elevator selection based on proximity and direction
- Handles both pickup and dropoff requests
- Tracks number of people in each elevator
- Supports up and down movements
- Maximum capacity per elevator

## System Components

1. **Lift (Elevator)**
   - Tracks current floor, direction, and stops
   - Manages pickup and dropoff points
   - Handles movement logic
   - Maximum capacity: 10 people

2. **LiftManager**
   - Manages multiple elevators
   - Handles new lift requests
   - Selects the most appropriate elevator for each request
   - Coordinates elevator movements

## Usage Example

```go
// Initialize system with 6 floors and 2 elevators
manager := NewLiftManager()
manager.Init(6, 2)

// Request elevator from floor 0 to floor 3
elevatorId := manager.RequestNewLift(0, 3)

// Move elevators one tick
manager.Tick()

// Get current state of all elevators
states := manager.GetLiftStates()
```

## API

- `RequestNewLift(start, dest int) int`: Request an elevator from start floor to destination floor
- `Tick()`: Move all elevators one step
- `GetLiftStates() []string`: Get current state of all elevators
- `GetNumberOfPeopleOnLift(liftId int) int`: Get number of people in specified elevator
- `GetLiftsStoppingOnFloor(floor int, dir rune) []int`: Get elevators stopping at specified floor

## Building and Running

```bash
go build
./elevator-design
``` 