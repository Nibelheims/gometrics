# WARNING
This small project is still under construction.

# gometrics
Gather some system metrics like CPU &amp; RAM usage and send them over HID RAW
reports. This is meant to display resource usage on keyboards with embedded
displays. Tested with a [Lily58](https://github.com/kata0510/Lily58)
running [QMK](https://docs.qmk.fm/#/) with the
[HID RAW](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_rawhid.md)
feature enabled.

# Building
This requires CGO. You can check it is enabled using `go env -w CGO_ENABLED=1`.
