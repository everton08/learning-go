package main

type Car struct {
    Id              int       `json:"id"`
    Fuel            string    `json:"fuel"`
    Brand           string    `json:"brand"`
    Model           string    `json:"model"`
    LicensePlate    string    `json:"licensePlate"`
}

type Cars []Car
