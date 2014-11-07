Create Ubuntu Bootable USB 
--------------------------

```
hdiutil convert -format UDRW -o ~/path/to/target.img ~/path/to/ubuntu.iso
```
```
diskutil unmountDisk /dev/diskN
```

```
sudo dd if=/path/to/downloaded.img of=/dev/rdiskN bs=1m
```
```
diskutil eject /dev/diskN
```

##### References:

- http://www.ubuntu.com/download/desktop/create-a-usb-stick-on-mac-osx
