---
title: "The Classical XY Model"
description: "[physics] [CA]"
date: "1970-01-02"
program: true
programName: "xy.wasm"
thumbnail: "xy.png"
---

The classical XY model serves as a 2D generalization of the [Ising model](/posts/ising). While the Ising model represents spins as discrete values (+1 or -1), the classical XY model introduces a continuous representation using angles ranging from 0 to 2π. In this article, we will delve into the simulation of the classical XY model, highlighting the key differences from the Ising model and exploring the fascinating properties that emerge from this extension.

To simulate the classical XY model, a slightly modified approach is employed due to the continuous nature of the spins. Despite this, the fundamental rules governing the model remain remarkably simple, akin to the Ising model. These rules govern the interactions and dynamics of the spins within the lattice, revealing the underlying behaviour of the system.

In contrast to my Ising model Simulation, this model introduces two additional variables: the interaction strength and the external field. The interaction strength quantifies the influence exerted by neighbouring lattice spins on a given point, shaping the system's overall behaviour. On the other hand, the external field introduces a bias in spin interactions, effectively simulating the impact of an extra neighbour interacting with each lattice point. The temperature, similar to the Ising model, plays a vital role in controlling the dynamics and fluctuations within the system.

The rules governing the classical XY model are expressed through the following equation:

`$$G_{n,t} = \sin(G_{i,j,t}-G_{i-1,j,t}\;) + \sin(G_{i,j,t}-G_{i+1,j,t}\;) + \sin(G_{i,j,t}-G_{i,j-1,t}\;) + \sin(G_{i,j,t}-G_{i,j+1,t}\;) \\ G_{i,j,t+1} = I G_{n,t} + Tr + x \sin(G_{i,j,t})$$`

In this equation, $G_{i,j,t}$ represents the spin at lattice point $(i,j)$ at time $t$. $I$ denotes the interaction strength, $T$ represents the temperature, $r$ is a random number between -1 and 1, and $x$ signifies the strength of the external field.

One of the intriguing characteristics of the classical XY model is the formation of vortices and antivortices in pairs. These vortices manifest as circular rainbows joined by a continuous spectrum of colours, reflecting the angles of the spins. This unique property arises from the continuous nature of the spins and adds a visually captivating aspect to the model's behaviour.

Another fascinating feature of the classical XY model is the existence of temperature and interaction strength phase transitions. At critical points, significant changes occur in the system's behaviour, leading to the emergence of distinct phases. Exploring these phase transitions provides valuable insights into the complex dynamics and collective behaviour of the spins within the lattice.

The code for this model is available on [GitHub](https://github.com/e74000/Classical-XY-Model-Go/).