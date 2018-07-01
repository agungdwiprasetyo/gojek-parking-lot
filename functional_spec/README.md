# GO-JEK Parking Lot Problem

**author: Agung Dwi Prasetyo**

Solve parking lot problem written in Golang.


## Component

Component list in this application:

#### Parking
```parking``` is main component, for initialize a parking lot. In each parking lot there is information of slot capacity and a set of slots that will be occupied by the car.

```parking``` object has fields:
1. ```Capacity (uint)```, for total capacity a parking lot
2. ```Slots ([]Slot)```, a set of slots that will be occupied by the car

Then, ```parking``` has methods:
1. ```New()```, for constructor
2. ```Get()```, for get saved ```parking``` data from saved storage in each lifecycle.
3. ```parking.Save()```, for save ```parking``` state in a lifecycle.
4. ```parking.FindNearestSlot()```, for find nearest free slot.
5. ```parking.AddCar()```, for add new car in selected slot number.
6. ```parking.GetFilledSlots()```, for get all free slots.
7. ```parking.GetSlotsByCarColor()```, for get all slots in a selected car by color.
8. ```parking.GetSlotByCarNumber()```, for get slot in a selected car by registration number.
9. ```parking.RemoveCar()```, for remove car in specific slot in parking lot by car identity.
10. ```parking.RemoveCarBySlot()```, for remove car in specific slot in parking lot by slot number.

package source: ```functional_spec/parking```

#### Slot
```slot``` is a cleared area that is intended for parking car, with identity serial index number.

```slot``` has fields:
1. ```Index (uint)```, for serial number sign in a slot.
2. ```Car (Car)```, for car object identity.

Then, ```slot``` has methods:
1. ```New()```, for constructor.
2. ```slot.Allocate()```, for allocate new car that will park in this slot.
3. ```slot.GetCarNumber()```, for get car registration number info who was parked in this slot.
4. ```slot.GetCarColor()```, for get car color info who was parked in this slot.
5. ```slot.Free()```, for make slot empty again from the car.
6. ```slot.IsFree()```, for check this slot is empty or not.

package source: ```functional_spec/slot```

#### Car
```car``` is basic object.  ```car```fill a ```slot``` and reduce parking lot slot capacity.

```car``` has fields:
1. ```Number (string)```, show registration number from a car
2. ```Color (string)```, show color from a car


Then, ```car``` has methods:
1. ```New()```, for constructor.
2. ```car.IsEqual(car2)```, for check with another car is equal or not (check registration number and color between two objects car).

package source: ```functional_spec/car```

#### Command
```cmdmanager``` is a module that contains a set of commands. which serves to control the *lifecycle* of the above 3 components.

package source: ```functional_spec/cmdmanager```

## Unit testing

* ```parking```. Contain in directory ```functional_spec/parking```, in file ```parking_test.go```
```sh
$ cd functional_spec/parking
$ go test -v
```

* ```slot```. Contain in directory ```functional_spec/slot```, in file ```slot_test.go```
```sh
$ cd functional_spec/slot
$ go test -v
```

* ```car```. Contain in directory ```functional_spec/car```, in file ```car_test.go```
```sh
$ cd functional_spec/car
$ go test -v
``` 

* ```cmdmanager```. Contain in directory ```functional_spec/cmdmanager```, in file ```*_test.go```
```sh
$ cd functional_spec/cmdmanager
$ go test -v
``` 

## Setup application

Run bash ```setup``` in ```bin``` directory
```sh
parking_lot $ bin/setup
```
That command has automatically run unit testing, unit testing commands contained in the file  ```bin/run_functional_tests```

This application fully controlled by command. Run bash ```parking_lot``` in bin directory with 2 options:

* The inputs commands are expected and taken from the file specified
```sh
parking_lot $ bin/parking_lot [input_filepath]
```
* Or start the program in interactive mode.
```sh
parking_lot $ bin/parking_lot
```
**Command list**

* ```>> create_parking_lot [capacity]```
Initialization of parking lot with parameters of slot capacity. This command must be run first to initialize the parking lot.

* ```>> park [car_registration_number] [car_color]```
The parking car in available slot with identity of registration number and color.
If success, program will print ```Allocated slot number: [nearest_slot_number]```. If failed,
(parking lot is full, slot already filled) will print ```Sorry, parking lot is full```

* ```>> leave [slot_number]```
The slot is available again after the car leaves the parking lot (give the entrance ticket) so that the slot can be occupied by the another car will park.

* ```>> status```
For print parking area status in table format.
```Slot No.    Registration No     Colour```

* ```>> registration_numbers_for_cars_with_colour [car_color]```
For print all registration car with specific color who was parking.

* ```>> slot_numbers_for_cars_with_colour [car_color]```
For print all slot number with specific car color who was parking.

* ```>> slot_number_for_registration_number [car_registration_number]```
For print slot number with specific car registration color who was parking.
