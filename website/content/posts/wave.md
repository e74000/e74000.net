---
title: "Wave Simulation"
description: "[physics] [CA]"
date: "1970-01-03"
program: true
programName: "wave.wasm"
thumbnail: "wave.png"
---

The wave equation is a fundamental linear differential equation capable of representing a wide range of wave phenomena. This simple equation provides a versatile framework for understanding and simulating wave behaviour from electromagnetic waves to sound waves and ripples in water. In this post, I hope to show the wave equation, its compact representation, and the discretization process necessary for numerical simulations.

The wave equation can be succinctly expressed as:

`$$\ddot{u} = c^2 \nabla^2u$$`

Here, $u$ represents the displacement, and $\nabla^2$ signifies the Laplacian operator, a mathematical tool that quantifies changes in a function between different points in space.

In order to simulate the wave equation, spatial discretization is essential. The first step involves discretizing the Laplacian operator to operate on a 2D lattice. This can be achieved by transforming it into a 2D convolution using the following kernel:

`$$K = \begin{bmatrix} 0 & 1 & 0 \\ 1 & -4 & 1 \\ 0 & 1 & 0 \\ \end{bmatrix}$$`

Consequently, the wave equation can be expressed in discrete form as:

`$$\ddot{u}_{i,j} = c^2 (u_{i-1,j} + u_{i+1,j} + u_{i,j-1} + u_{i,j+1} - 4 u_{i,j})$$`

Once the equation is discretized spatially, the next step involves discretizing it in time. To achieve this, [Verlet](https://en.wikipedia.org/wiki/Verlet_integration) integration, a numerical method used for integrating the motion of objects, can be employed. Verlet integration breaks the motion into steps of $\Delta t$, and the following relationship holds:

`$$x_{t+1} = 2x{t} - x{t-1} + \ddot{x}_{t} \Delta t^2$$`

Applying Verlet integration to our wave equation yields the following expression:

`$$\ddot{u}_{i,j,t+1} = 2u_{i,j,t} - u_{i,j,t-1} + c^2 (u_{i-1,j,t} + u_{i+1,j,t} + u_{i,j-1,t} + u_{i,j+1,t} - 4 u_{i,j,t}) \Delta t^2$$`

By employing the discretized spatial and temporal forms of the wave equation, it becomes feasible to run numerical simulations of the wave equation on a computer.

Above, you can see a simple demonstration of this model.