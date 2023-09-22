---
title: "The Ising Model"
description: "[physics] [CA]"
date: "1970-01-01"
program: true
programName: "ising"
thumbnail: "ising.png"
---

The Ising model is a powerful statistical model used to represent the "spins" of atoms in a ferromagnetic material. In this article, I will attempt to explain the Ising model and show how to perform a Monte Carlo simulation to understand its behaviour. Additionally, I will try to uncover some intriguing features, including the impact of temperature and the emergence of phase transitions.

The Ising model represents particles arranged in a 2D lattice, where each particle is assigned a spin value of either +1 or -1. These spins interact with neighbouring spins on the lattice, reflecting the interactions between atoms in a ferromagnetic material. Notably, the model assumes that interactions between non-adjacent spins are negligible to simplify computations.

Incorporating the effect of temperature, the Ising model introduces heat that disturbs the spin interactions. The temperature determines the level of noise in the system, affecting the overall behaviour of the model. As the temperature increases, the noise amplifies, potentially leading to a loss of order in the spins.

To simulate the Ising model using a Monte Carlo approach, the model must repeatedly select a location on the lattice and calculate the probability of that cell changing its spin state. The probability is determined using the following equation, which considers the spins of neighbouring cells:

`$$G_{n} = G_{i+1,j} + G_{i-1,j} + G_{i,j+1} + G_{i,j-1} \\ p = e^{-2 \beta (G_{i,j} + G_{n})}$$`

In this equation, $G_{i,j}$ represents the spin at cell $(i, j)$, $G_{n}$ is the sum of spins in neighbouring cells, and $\beta$ is the reciprocal of the temperature.

The term - (G_{i,j} + G_{n}), known as the Hamiltonian, corresponds to the energy of the spin configuration. As a result, the probability of a state change is proportional to the energy of the given spin. Intuitively, this implies that the system tends to minimize its energy, leading to a decrease in overall energy as the simulation progresses.

At the critical temperature of the Ising model, an extraordinary phenomenon emerges—self-similarity. When observing the model at different scales, from small local neighbourhoods to larger regions of the lattice, striking similarities can be observed. For example, patterns that occur at a small scale are replicated and echoed at larger scales, revealing a fractal-like nature within the system. This self-similarity suggests that the underlying dynamics driving the Ising model at criticality exhibit a form of universality, where the behaviour remains consistent regardless of the observation scale.

One of the fascinating aspects of the Ising model is its ability to exhibit phase transitions. Even though it is one of the simplest statistical models, it showcases a shift in behaviour at critical temperatures. You can explore this phenomenon yourself by running the provided program, which simulates the Ising model. The code for this project can be found on my [GitHub](https://github.com/e74000/Simple-Ising-Model-Simulation/).
