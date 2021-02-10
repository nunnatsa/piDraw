# PiHATDraw
RasberyPi [Sense HAT](https://www.raspberrypi.org/products/sense-hat/) is a device that is connected to a raspberry Pi and includes several additional hardware and sensors. It includes a 8X8 LED display and a joistic.

This software should run on a raspberry pi with a hat. It uses the hat to draw: the display is a virtual window - a subset of the bigger canvas. To draw, move the cursor to the required pixel and press he joystick to set the pixel.

The software also starts a web application that serves a HTML page to display the full draw.

![Sense HAT](images/pi-hat.jpeg | width=250)

![webapp](images/weapp.png | width=300)

## How to Build and Run
I checked this on raspberry pi 2B with a [Sense HAT](https://www.raspberrypi.org/products/sense-hat/), with golang version 1.15.8

Just installing golang on rasperry pi will give you an old version of go. Better way is to download and install the latest version:

Got to [golang download page](https://golang.org/dl/) and get the file for armv6l. Currently it's [go1.15.8.linux-armv6l.tar.gz](https://golang.org/dl/go1.15.8.linux-armv6l.tar.gz)

You can use the instructions [here](https://pimylifeup.com/raspberry-pi-golang/) for the installation.

After that, clone this project in a directory in your go path, e.g. `~/go/src/github.com/nunnatsa`:
```shell
git clone https://github.com/nunnatsa/piDraw.git
```

Now build the software:
```shell
cd piDraw
go build .
```

If everything went fine, the new piDraw executable file is created. Go on and run it:
```shell
./piDraw
```

You can now draw with the HAT joystic on the HAT display.

When on the same network as your raspberry pi, open a web browser and go to your raspberry pi hostname in port 8080, the default is [http://raspberrypi:8080/](http://raspberrypi:8080/)

Now you can see the full drawing. Notice that the page is sync with the HAT. Open another tab or browser and noticed that both of them are updated simultaneity.
