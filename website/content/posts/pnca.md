---
title: "Probabilistic Neural Cellular Automata"
description: "[CA]"
date: "1970-01-06"
program: true
programName: "pnca.wasm"
thumbnail: "pnca.png"
---

Cellular automata provide a fascinating framework for studying complex systems with distinct critical thresholds. In this post, I plan to show a cellular automaton that serves as a remarkably simple model of neurons.

Each cell in the grid represents a neuron characterized by an activation level ranging from 0 to 1. By examining this model, we can observe the emergence of complex behaviour and gain insights into criticality and phase transitions within neuronal systems.

In this cellular automaton, neurons interact with their neighbouring neurons based on their activation levels. If a neuron's activation surpasses a threshold value of 0.5, it has a probability of activating one of its neighbouring neurons. The parameter σ represents the expected number of neighbours each neuron could activate. By adjusting σ, we can explore the effects of different activation propensities on the system.

Additionally, on each timestep, a certain amount of noise is introduced, simulating the inherent stochasticity and variability present in neuronal systems. Moreover, activations are proportionally reduced by a damping factor, representing the tendency of activations to diminish over time.

The most intriguing property of this cellular automaton is its ability to exhibit complex behaviour when σ = 1. When σ is slightly above or below 1, the simulation gradually tends to either full activation or complete inactivation. This behaviour signifies that a phase transition occurs as σ approaches a value greater than 1. Notably, complex behaviour can only emerge when σ lies precisely at the midpoint of this phase transition. This delicate balance highlights the critical nature of neuronal systems and their propensity for exhibiting intricate dynamics.

Interestingly, this cellular automaton offers insights into the behaviour of the brain itself. There is a theory that suggests the brain operates near a similar phase transition, indicating that criticality plays a crucial role in its functioning. However, it is important to note that this theory is still a subject of debate and ongoing research, and consensus has yet to be reached regarding its validity. Despite this, studying simple models like this cellular automaton can provide valuable perspectives and spark further exploration into the behaviour of complex biological systems.