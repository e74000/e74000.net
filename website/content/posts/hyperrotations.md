---
title: "Hyper-Rotations"
description: "[maths] [visualisation]"
date: "1970-01-04"
program: true
programName: "cube.wasm"
thumbnail: "cube.png"
---

## Hypercubes

Most shapes can be characterised by vertices with edges that connect them. This still applies in higher dimensions. When you want to perform a linear transformation on a shape, all you need to do is multiply each vertices' position vector by the transformation matrix representing your transformation. One way that you can represent this information in the computer program is with a list of vectors. However, you can make the program run faster by concatenating them into one vertex matrix. The edges can then be represented with an adjacency matrix. Since each column in the vertices matrix represents a vertex, you can use rows and columns of the adjacency matrix to describe each column. Every row corresponds to one of the vertices, and each $1$ corresponds to an edge. Here is an example of what these matrices would look like in the case of a standard 3D cube:

$$Vertex = \begin{bmatrix} -1 & +1 & -1 & +1 & -1 & +1 & -1 & +1 \\\\  -1 & -1 & +1 & +1 & -1 & -1 & +1 & +1 \\\\ -1 & -1 & -1 & -1 & +1 & +1 & +1 & +1 \end{bmatrix} \\\\ ~~ \\\\ Adjacency= \begin{bmatrix}  0 & 1 & 1 & 0 & 1 & 0 & 0 & 0 \\\\  1 & 0 & 0 & 1 & 0 & 1 & 0 & 0 \\\\  1 & 0 & 0 & 1 & 0 & 0 & 1 & 0 \\\\  0 & 1 & 1 & 0 & 0 & 0 & 0 & 1 \\\\  1 & 0 & 0 & 0 & 0 & 1 & 1 & 0 \\\\  0 & 1 & 0 & 0 & 1 & 0 & 0 & 1 \\\\  0 & 0 & 1 & 0 & 1 & 0 & 0 & 1 \\\\  0 & 0 & 0 & 1 & 0 & 1 & 1 & 0 \end{bmatrix}$$

Notice anything?

The vertices comprise every possible combination of three +1s and -1s, arranged like this, and it looks a bit like counting in binary.

The adjacency matrix is slightly less obvious, however. One thing that might stick out is that each row/column only contains three 1s. What really is happening is that there is an edge between every pair of vertices with a distance of one between them. This can be worked out computationally by counting the ones after applying a bitwise XOR to the row/column index.

The other interesting fact of this adjacency matrix is that it makes a self-similar fractal. This is made more evident when you increase the number of dimensions. Pictured below is the adjacency matrix for a 9D hypercube.

![A fractal created from the edge distances of a 9d hypercube.](/images/XORFractal.png)

You can use this method to generate hypercubes in arbitrarily high dimensions.

## Projections
When displaying 3d graphics on a computer screen, a projection is used to map points in 3d space to 2d points on the screen. The same applies to high-dimensional spaces.

One of the simplest methods is to ignore higher dimensions. This is an orthographic projection. This has some benefits, as it is very performant and easy to implement.
In order to project a high dimensional vector $V_n$ to 2d space, you simply create a new 2d vector $P_2$ where $P_2 = (\begin{smallmatrix} V_x \\\\ V_y \end{smallmatrix})$ . Very simple!

Another commonly used type of projection is an isometric projection. In 3d, you map each axis to a set of 2d unit vectors, each spaced $120^\circ$ from each other. This doesn't appear to generalise to higher dimensions very well, but you can (sort of) extrapolate it such that each axis in your n-dimensional space maps to a set of unit vectors, each spaced $\frac{360^\circ}{n}$ from each other.
This doesn't make much sense, as you can end up with 2 axes that point in opposite directions, but it appears to work fine.

When you perceive the 3d world around you, you don't use orthographic or isometric projections. Instead, you use what is known as a perspective projection. You can approximate a 3d perspective projection by dividing a vector's $x$ and $y$ components by its depth (Called a "weak perspective projection").

For example if you had 3d vector $V_3,$ you could create a 2d vector $P_2$ and where $P_2 = (\begin{smallmatrix} V_x /  V_z \\\\ V_y / V_z \end{smallmatrix})$

This is good, but how can you upgrade this to work with higher dimensions?

Fortunately, there are two approaches to this! One of them is effectively like an orthographic projection, but instead, you remove all dimensions other than $x$ , $y$ and $z$ and then apply the weak perspective projection. This method is called "Perspective-Trim" in the program.


If you want to vaguely conserve the information from the higher dimensions, an alternate option is to average excess dimensions into 3 before applying the weak perspective projection. This option is called "Perspective - Avg" in the program above.

## Rotations

There are a few ways that you can rotate vectors. The method that this project uses is called Givens rotations.

For an $n$-dimensional rotation there is a $n \times n$ matrix representing the givens rotation:

$$G(\theta) = \begin{bmatrix} 1 & \cdots & 0 & \cdots & 0 & \cdots & 0 \\\\ \vdots & \ddots & \vdots && \vdots && \vdots \\\\ 0 & \cdots & \cos{\theta} &\cdots & -\sin{\theta} & \cdots & 0 \\\\  \vdots && \vdots & \ddots & \vdots && \vdots \\\\ 0 & \cdots & \sin{\theta} &\cdots & \cos{\theta} & \cdots & 0 \\\\ \vdots && \vdots && \vdots & \ddots & \vdots \\\\ 0 & \cdots & 0 & \cdots & 0 & \cdots & 1 \end{bmatrix}$$

To generate each rotation matrix, you pick two integers $i$ , $j$ where $i>j$ , and an identity matrix $G$ . You can then set $G_{i,i} = G_{j_j} = \cos \theta$ , $G_{i, j} = \sin \theta$ and $G_{j,i}= - \sin \theta$

A givens rotation matrix exists for every possible combination of $i$ and $j$ , leading to the number of rotation axes for a dimension $n$ to be $^nC_2$.

An interesting property of givens rotations is that they only act in 2 dimensions at a time, leaving all other dimensions unaffected. This can prove difficult when projecting shapes in dimensions higher than four down to three as if you were just to ignore the extra dimensions, there would be some rotation axes that do not cause any visual difference in the projection.

To apply the rotation, you multiply the vector you want to rotate by each rotation matrix in a fixed order. The order of the rotation must remain the same because Euler rotations are not communicative. You can visualise this intuitively by rotating a die in three axes but in different orders.

## Other Stuff

In 3D space, when you rotate something, say that you are rotating around an axis such as $X$ , $Y$ or $Z$ , where that axis is a line unaffected by the rotation. In 4D space, there are always 2 unaffected axes, so you are rotating around a hyperplane instead. This means you could rotate something around $XY$ , $XZ$ , $XW$ , $YZ$ , $YW$ or $ZW$ . In fact, in even higher dimensional spaces, you can rotate around solids and even hypersolids!