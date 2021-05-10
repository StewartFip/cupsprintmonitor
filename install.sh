#!/bin/bash

env GOOS=linux GOARCH=arm GOARM=6 go install && scp ~/go/bin/linux_arm/cupsprintmonitor pi@raspberrypi.local:/home/pi/bin && rm ~/go/bin/linux_arm/cupsprintmonitor