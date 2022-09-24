---
title: "The Classical XY Model"
description: "[physics/CA]"
date: "1970-01-02"
program: true
programName: "xy.wasm"
thumbnail: "xy.png"
---

# The Classical XY Model

The classical XY model is a 2D generalisation of the [Ising model](/posts/isingmodel).
This means that instead of the spins as being represented as either a `$+1$` or a `$-1$` they
are represented as an angle between `$0$` and `$2 \pi$`.

The simulation of the classical XY model uses a slightly different method to the Ising model
due to the fact that the spin is no longer discrete. However, the rules that govern the
model are equally simple.

Also, for this model I have added two more variables than the Ising model: Interaction
strength and external field. The interaction strength describes how much neighboring
lattice spins affect the spin of any given point. The external field biases the
interactions of the spins in a certain direction, basically if all the points in the
lattice had an extra neighbor interacting with them.
The temperature is functionally identical to that of the Ising model.

The rules are fairly simple:

`$G_{n,t} = \sin(G_{i,j,t}-G_{i-1,j,t}\;) + \sin(G_{i,j,t}-G_{i+1,j,t}\;) + \sin(G_{i,j,t}-G_{i,j-1,t}\;) + \sin(G_{i,j,t}-G_{i,j+1,t}\;)$`

`$G_{i,j,t+1} = I G_{n,t} + Tr + x \sin(G_{i,j,t})$`

Where `$G_{i,j,t}\;$` is the spin of a point `$(i,j)$` on the lattice at time `$t$`, `$I$` is the interaction strength,
`$T$` is the temperature, `$r$` is a random number in the range `$-1 \le r \le 1$` and `$x$` is the strength of the
external field.

In the model above, the angle of the spin is represented by the hue. One of the interesting properties of this
is how it forms vortices and antivortices in pairs. You can see these in the model as two circular rainbows, joined
by a colour. 

The other cool property of this model is that there is a temperature phase change as well as an interaction strength
one.

You can see the code for this model on [GitHub](https://github.com/e74000/Classical-XY-Model-Go/).