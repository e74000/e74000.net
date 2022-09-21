---
title: "The Ising Model"
description: "[physics/CA]"
date: "2022-09-07"
program: true
programName: "ising.wasm"
---

# The Ising Model

The Ising model is a statistical model to
represent the "spins" of atoms in a 
ferromagnetic material.

The particles in this model are arranged in
a 2d lattice and given a spin that can be
either +1 or -1. Each spin will then
interact with the spins in other positions
on the lattice.

To make the computation more efficient it is 
assumed that the interaction between 
non-adjacent spins is negligible.

Heat is also added that disturbs the
interactions and adds a certain amount of
noise proportional to the temperature.

To perform a monte-carlo simulation of this
all you need to do is repeatedly pick a
location on the lattice and then work out
the probability of that cell changing state
using the following equation:

`$G_{n} = G_{i+1,j} + G_{i-1,j} + G_{i,j+1} + G_{i,j-1}$`

`$p = e^{-2 \beta (G_{i,j} + G_{n})}$`

Where `$G_{i,j}$` is the spin in a cell at `$(i, j)$`
on the lattice, `$G_{n}$` is the sum of the spins
of all neighboring cells, and `$\beta$` is the
reciprocal of the temperature.

In this equation `$-(G_{i,j} + G{n})$` is known as the
hamiltonian, which effectively is the energy of
this configuration of spins. Seeing as the
probability of a state change is then proportional
to the energy of the given spin - it would make
sense intuitively that this acts to reduce energy.

Another interesting detail is that the ising model
is among the simplest statistical models that shows
a phase transition - try to see it yourself with the
program above.

You can see the code for this on [GitHub](https://github.com/e74000/Simple-Ising-Model-Simulation/)
