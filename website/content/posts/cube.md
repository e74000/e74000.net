---
title: "nD Rotations"
description: "[maths]"
date: "1970-01-04"
program: true
programName: "cube.wasm"
thumbnail: "cube.png"
---

# Rotating Shapes in N Dimensions

> Apologies in advance for the poorly ordered parameter menu - there is a limitation
in my website that means that I cannot control the order that they appear in.
I probably could fix it, but I would rather spend my time making more projects :)

In the program above you can see and interact with a 5D cube. The actual program 
that is running can support arbitrarily large numbers of dimensions however, since
the number of rotation axis grows with `$^nC_2$` of the number of dimensions, I
figured that five dimensions was the largest number of sliders I could reasonably
put under the program.

The method I chose to rotate the cube with is by making a set of Givens rotation
matrices when the cube is created for the set number of dimensions. You can see
how I made the matrices on [wikipedia](https://en.wikipedia.org/wiki/Givens_rotation)
and my [GitHub](https://github.com/e74000/nD-Cubes/).

One of the facts that I find particularly interesting is how you project n-dimensional
shapes into 3 dimensions and then 2. For converting 3 dimensions to 2 I used the simple
computer graphics equation of `$x_{screen} = x / z$` and `$y_{screen} = y / z$`.

To convert higher dimensions to 3 dimensions you effectively need to "compress" the
dimensions into `$x$`, `$y$` or `$z$`. My strategy for doing this was to create a 
matrix of dimensions 3 by n and then multiplied that by the vector to be projected.