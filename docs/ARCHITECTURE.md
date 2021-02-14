# Software Architecture
## General Description
The architecture of this software is using the event-driven approach.

The HAT joystick produces 5 type of events (Move up, down, left or right, and pressing). The HAT sends these events to a
dedicated channel. The HAT display is listening to another channel, with a display events. These events contain the required 
matrix to be displayed, and the relative location of the cursor.

The Canvas is the in memory storage for the picture.

The webapp present the simple HTML page. When the client open this page, a javascript script opens a websocket with the 
web application. This connection stays open for continuous messaging from the webapp.

When opening the websocket, the webapp registers the connection as a channel in the notifier.

The controller is responsible to coordinate everything. It receives the joystick events and modifies the canvas accordingly, 
then triggers an event with the floating window content and the location of the cursor, by sending this information to the 
HAT display channel. The HAT display reads these message from the channel and update the HAT physical display accordingly.

In addition, the controller publishes an event using the notifier with full canvas data, the cluster location, and the floating window location. 
The notifier then sends the message to all the registered client - meaning the open websockets, or the active clients in other words.

The javascript in the HTML keep reading these messages and updates the web display.

The webapp also provides two APIs: `/api/canvas/reset` to reset the picture, and `/api/canvas/color`
to change the cursor (the pen) color. The javascript script uses these APIs when the user press buttons in the web page.
The webapp, when handling these APIs, triggers a client action event - which is a message in the client action channel.
Again, the controller reads these messages from this channel, modifies the canvas and produces the same events as done in 
joystick events

The software uses the https://github.com/nathany/bobblehat packages in order to perform the HAT hardware related
functionality.
