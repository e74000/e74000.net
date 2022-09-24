---
title: "Wave Equation Simulation"
description: "[physics/CA]"
date: "1970-01-03"
program: true
programName: "wave.wasm"
thumbnail: "wave.png"
---

# Wave Equation simulation

This is a very simplified model of the wave equation.
It uses a very simple update method that uses a finite
difference method to perform a numerical simulation
of the wave equation.

This browser simulation has been made fairly low
resolution so that it runs smoothly on all devices
however this simulation can be made arbitrarily 
precise.

The update method uses the following equation:

`$G_{i,j,t+1} = 2 G_{i,j,t} - G_{i,j,t-1} + C(G_{i-1,j,t} + G_{i+1,j,t} + G_{i,j-1,t} + G_{i,j+1,t} - 4 G_{i,j,t})$`

Where `$G_{i,j,t}$` is the displacement of point `$(i,j)$` at time `$t$` and `$C$` is a constant `$\le 1$`.
