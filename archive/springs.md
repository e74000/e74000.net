---
title: "Spring Wave Model"
description: "[physics/CA]"
date: "1970-01-05"
program: true
programName: "springs.wasm"
thumbnail: "springs.png"
---

> You can click on a point in the program window to add a velocity to a specific cell... 

This is a simple simulation of how waves would travel on a square 2x2 lattice of springs.

The simulation is designed as a cellular automata, where each pixel represents the
displacement at a given point on a grid of springs.

Each cell in the grid of cells is connected to its neighbors with springs, as well as a
spring that connects the cell to zero vertical displacement.

This means you can get the force exerted on the cell at any given point by the difference
in displacement to a neighboring grid cell times the neighbor spring constant, plus the
displacement of the cell times the global spring constant. These parameters can be adjusted
in the program.

If you take `$a_{i,j,t}$`, `$v_{i,j,t}$` and `$x_{i,j,t}$` as the
acceleration, velocity and displacement at `$i,j$` at time `$t$`, and `$k_n$` is the neighbor spring constant `$k_g$`
is the global spring constant, as well as `$\Delta t$` and `$m$` is the mass of the cells.

This means that you can express the update function as follows:

`$$a_{i,j,t+1} = \frac{1}{m}(k_{n} (x_{i-1,j,t}\; + x_{i+1,j,t}\; + x_{i,j-1,t}\; + x_{i,j+1,t}\; - 4x_{i,j,t}) - k_g x_{i,j,t}) \\ v_{i,j,t+1} \; = v_{i,j,t} \;  + a_{i,j,t+1}\; \times \Delta t \\ x_{i,j,t+1} \; = x_{i,j,t+1} \; + v_{i,j,t}\; \times \Delta t$$`

You may notice that some of this looks very similar to the wave equation solution from my [other post](../website/content/posts/wave.md).
This is because they are pretty much two different ways of approaching the same problem.

I find this model particularly neat since it allows you to think about waves much less abstractly than you would
otherwise. It also gives you some extra parameters that you wouldn't otherwise have.

One of the assumptions that you have to make with this model is that the cells cannot move laterally. If you were
to allow this to happen though, and `$k_g \; = 0$`, then this would make a pretty neat model for how cloth deforms.