The config below has been edited to use natural scrolling with trackpad:

```
Section "InputClass"
        Identifier "libinput touchpad catchall"
        MatchIsTouchpad "on"
        MatchDevicePath "/dev/input/event*"
        Driver "libinput"
	      Option "NaturalScrolling" "True"
EndSection
```
