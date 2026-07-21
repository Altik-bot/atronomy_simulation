# atronomy_simulation
# Astronomy Simulation Project

## Overview

This project simulates gravitational interaction between celestial bodies using Newtonian physics. The current focus is a two body system, with planned extension to three body and N body systems.

The same simulation will be implemented in three languages:

* Python
* C++
* Go

The goal is to compare:

* execution speed
* computational efficiency
* memory and disk usage

Each implementation will solve the same physics problem with identical initial conditions.

---

## Features

* Two body gravitational simulation
* Step based numerical integration
* Trajectory tracking
* Multi language implementation
* Performance comparison framework

---

## Physics Model

### Gravitational Force

F = G * (m1 * m2) / r²

### Vector Form

F⃗ = G * m1 * m2 / |r|³ * r⃗

Where:

* r⃗ = (x2 - x1, y2 - y1)
* |r| is distance

### Acceleration

a⃗ = F⃗ / m

### Velocity Update

v⃗ = v⃗ + a⃗ * dt

### Position Update

x⃗ = x⃗ + v⃗ * dt

---

## Algorithm

Each simulation step:

1. Compute distance between bodies
2. Compute force vector
3. Convert force to acceleration
4. Update velocity
5. Update position
6. Store trajectory

---

## Code Structure

### Body Class

Each body stores:

* mass
* position (x, y)
* velocity (vx, vy)
* force (fx, fy)

### Trajectory Storage

* path1_x, path1_y
* path2_x, path2_y

---

## Visualization

The simulation will include visualization to analyze motion and stability.

Planned tools:

* Python: matplotlib for plotting trajectories
* C++: simple output with optional integration with SFML or OpenGL
* Go: basic plotting or export data for visualization

Visualization features:

* trajectory curves for each body
* orbit shape analysis
* comparison between runs

---

## Multi Language Benchmark

Each language will run the same simulation with identical parameters.

Metrics to compare:

* execution time per simulation
* total runtime for fixed number of steps
* memory usage during execution
* compiled binary size or script size

### Why these languages

* Python: fast development, slower execution
* C++: high performance, low level control
* Go: balance between speed and simplicity

---

## Numerical Considerations

### Time Step (dt)

* small dt increases accuracy
* large dt causes instability

### Distance Handling

* avoid division by zero
* use small epsilon for minimal distance

### Stability

* energy drift exists in simple integration
* better integrators planned

---

## Planned Improvements

### Physics

* three body simulation
* general N body system
* energy conservation tracking

### Algorithms

* Verlet integration
* Runge Kutta methods

### Visualization

* real time animation
* interactive parameter control

### Performance

* profiling each language
* optimization of loops and memory

---

## Example Use Cases

* binary star system
* planet orbit simulation
* chaotic three body motion

---

## How to Run

1. Clone repository
2. Choose implementation (Python, C++, or Go)
3. Run simulation file
4. View output or generated plots

---

## Motivation

This project builds:

* understanding of classical mechanics
* numerical simulation skills
* performance comparison between programming languages

It also prepares for more advanced simulations and research level problems.

---
