# hashimg

Written for Boston Golang "Chopped" meetup. Given the following packages (and
any of their subpackages), make anything:

* bytes
* crypto
* encoding
* image
* regexp

Additionally, the following packages from the standard library were available:

* flag
* fmt

I ended up using all of the "basket" packages, even where not strictly
necessary; e.g. the only real reason to base64 encode the resulting image is
just so that we can do something with the encoding package. Likewise, we don't
_really_ need a regular expression to validate the input, but why not throw it
in?

Given a "username", this program finds the SHA 256 checksum of the username and
uses it to generate a GitHub-style default avatar. The first three bytes are
used to create color A, the next three bytes are used to create color B, and the
next 25 bytes are used to form a 5x5 pixel grid, with the value being used to
determine whether the pixel should be colored with A or B.

Next, the 5x5 pixel image is tesselated into a 10x10 pixel image for some neat
mirroring, and then it is scaled up by a user-provided multiplier.

Finally, the image is base64 encoded and printed to standard output.

Demo:

    $ hashimg -username ben
    iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAIAAAAiOjnJAAACjklEQVR4nOzcsY0iQRBAUeZEShcTNkngTgJkdAHhXgrslv62ennPH3WP9FVWqa+3y+Oywt/z/u1v/93PJedOrLrz5NyJP0tO5dcTFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZG4Tj5etT/O+1bty5tYJIRFQlgkhEVCWCSERUJYJIRFQlgkhEVCWCSERUJYJIRFQlgkhEVCWCSERUJYJI5V77xP7Pju+Y53njCxSAiLhLBICIuEsEgIi4SwSAiLhLBICIuEsEgIi4SwSAiLhLBICIuEsEgIi4SwSBzP87Xk4B13zyc+7X9NLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLeef8hO955wsQiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiJxPM/XkoN3fG994tP+18QiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIicdwuj9V3+LLJHveOu+er7jxhYpEQFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZG4Tj7e8c10d37f5M4mFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFonRzvuOe9yfZtUb8SYWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFglhkRAWif8BAAD//5F8Qoy5p4odAAAAAElFTkSuQmCC

Result:

<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAADICAIAAAAiOjnJAAACjklEQVR4nOzcsY0iQRBAUeZEShcTNkngTgJkdAHhXgrslv62ennPH3WP9FVWqa+3y+Oywt/z/u1v/93PJedOrLrz5NyJP0tO5dcTFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZG4Tj5etT/O+1bty5tYJIRFQlgkhEVCWCSERUJYJIRFQlgkhEVCWCSERUJYJIRFQlgkhEVCWCSERUJYJI5V77xP7Pju+Y53njCxSAiLhLBICIuEsEgIi4SwSAiLhLBICIuEsEgIi4SwSAiLhLBICIuEsEgIi4SwSBzP87Xk4B13zyc+7X9NLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLeef8hO955wsQiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiJxPM/XkoN3fG994tP+18QiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIiISwSwiIhLBLCIiEsEsIicdwuj9V3+LLJHveOu+er7jxhYpEQFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZG4Tj7e8c10d37f5M4mFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFonRzvuOe9yfZtUb8SYWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFglhkRAWCWGREBYJYZEQFglhkRAWif8BAAD//5F8Qoy5p4odAAAAAElFTkSuQmCC">
