# A small CLI tool to generate preview images

This is a tiny command line application that generates a base64 encoded preview of an image to be used with progressive image loading. I built this mainly for myself but decided to open source it in case other people might find it useful.

I use this to generate placeholders for static images that don't get handled by the backend (otherwise I would just generate the placeholder there).

## Usage

```bash
$ preview ./my-image.jpg

/9j/2wCEAAEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAQEBAf/AABEIABgAIAMBIgACEQEDEQH/xAGiAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgsQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+gEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoLEQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2gAMAwEAAhEDEQA/APzu8VeGcRSnZ1B7ev8Ak186ajo/kXZO3oxPT349Of5/XNfdvi61hW3l4H3Tnvj69+enNfJfiGJDeOAAPmwPfk59enY/415uZKLozUtnFtt7bP8A4H9M/QnOStvpb8bdradfS+rO9+FzGGeBc8blGPxGAP8APtnAzX6a/DBXmht+DyF4x2IH5dR7Hk9q/OX4W6VJc3tuFUkFlHH1H1449e/0r9avg34Omlt7UmM4Kpn5Seccduv5f0P8keKtHDQVV3Slq916Lp9y8z9K4Nr1nUitbaetv6v+tj83fGd9/o8oBPQ5zxwM/lnp17etfLWoOZ9Qxkcvj9f849+vNfSXjP8A1M/+6f5mvmmf/kJD/roP5mv6nzupOOFm07aP56L/ADPz+lCMqiuv6X/D6n2R8BtAW6urTcoOWU9Ae44H4/pnkdK/bP4M+GIY7K0Plj7i549Rx29P/rHvX47fs8f8fNj/ALy/zFft38H/APjytf8Aci/9Br+KfFDFVp1qsZSuk5L8UtvT+tT9Z4Xw1OEYOK1ev467d7L8z//Z
```

Or pipe it directly to your clipboard (OSX):

```bash
$ preview ./my-image.jpg | pbcopy
```

Now you can use it as a placeholder while your real image loads.

```html
<img src="data:image/svg+xml;base64,/9j/2wCEAA/xAGiAAABBQEBAQE..." />
```

How to actually _implement_ progressive image loading is out of scope for this readme. There are multiple ways to go about it, however. See the links below for inspiration.

## Further Reading

- [Taking A Look At The State Of Progressive Images And User Perception](https://www.smashingmagazine.com/2018/02/progressive-image-loading-user-perceived-performance/) (2018)
- [How Medium does progressive image loading](https://jmperezperez.com/medium-image-progressive-loading-placeholder/) (2015)