package main

import (
    "machine"
    "time"
)

func main() {

    // Utilisation de la LED intégrée en surface de la carte, broche D13
    var led machine.Pin = machine.Pin(13)

    // Configuration de la broche en sortie
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})

    for {
        led.High() // sortie au niveau logique haut, LED allumée
        time.Sleep(time.Millisecond * 500) // temporisation 500 millisecondes
        led.Low() // sortie au niveau logique bas, LED éteinte
        time.Sleep(time.Millisecond * 500) // temporisation 500 millisecondes
    }
}
