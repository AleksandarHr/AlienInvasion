# Alien Invasion Simulator
A simple CLI for simulating an alien invasion of a made-up world.

---

* [Specifications](#specifications)
* [Assumptions](#assumptions)
* [Install](#install)
* [Run](#run)
* [Parameters](#parameters)

---

## Specifications

Mad aliens are about to invade the earth. You are given the following **map**:
* Every line contains information about *one city*.
* The *city name* is first.
* The city name is followed by *1-4 directions (north, south, east, or west)*. Each direction represents a road to another city in the specified direction.

In total **N** aliens will invade our world, where the number is specified as a command-line argument. The aliens behave as follows:
* Initially, every alien is spawned at a *random location*.
* Once all aliens are spawned, they start moving around. At each iteration, a given alien *can travel in any direction leading out of the city it currently occupies*.

In the case where two aliens are *in the same city*, the aliens **fight** and, as a consequence they both **die** and the city (along with all of its connection roads) **gets destroyed**.

The stopping conditions of the simulations are one of the following:
* If the iteration runs for the **maximum number of iterations** (specified as a command-line argument, defaulted at 10,000) the simulations stops.
* If, at any point before the maximum number of iterations are simulated, **all aliens have died**, the simulation stops.
* If, at any point before the maximum number of iterations are simulated, **all aliens are trapped** (e.g. all remaining aliens in the world are in cities which are isolated, have no road connections to any other moves, meaning those aliens do not have valid moved), the simulation stops. 
---

## Assumptions

This implementation of an alien invasino simulator works under the following assumptions:
* The map file contains a line for **each city** in the world.
* The information for a given city is **full**. In other words, no information about the topology of the world needs to be infered.
* Every connection between given two cities is considered a two-way connection. For example, if *Foo* has a connection to *Bar* (e.g. an alien can move from *Foo* to *Bar*), then *Bar* has a connection to *Foo* as well (e.g. an alien can move from *Bar* to *Foo*).
* City names contain only english letters and dashes (e.g. no other special characters). The dashes can be used to separate multi-part city names (e.g. Qu-ux). Alternatively, *Qu-ux* can be written as *Quux*.
* Aliens **do not move in parallel**, but in a random order every iteration. In other words, at any step of a given iteration, only one alien makes a move. The world state is updated to reflect this move (e.g. the location of the alien is updated, any fights are taken into account) before the next alien its move for the iteration.
* There are two explicit stopping conditions for the simulation - either **all aliens dead**, or each remaining alien has moved **at least max iterations** number of times. A third implicit stopping condition has been added - the simulation will stop if all remaining aliens are **trapped**, even if they have not all moved at least max iterations number of times, since trapped aliens have no valid moves.
* The input number of aliens and input number of maximum iterations fit in an *int*.

An example of an **invalid map file** (e.g. either the links are considered only as one-way, or there are lines missing for some of the cities):
> Foo north=Bar west=Baz south=Qu-ux
>
> Bar south=Foo west=Bee

Instead, the **correct map file** version of the above topology would look as follows:
> Foo north=Bar west=Baz south=Qu-ux
>
> Bar south=Foo west=Bee
>
> **Baz east=Foo**
>
> **Bee east=Bar**
>
> **Qu-ux north=Foo**

---

## Install

**Install Golang**

* Follow the [Go installation instructions](https://go.dev/doc/install) for the relevant OS.

**Download source code from GitHub**

```bash
git clone https://github.com/AleksandarHr/AlienInvasion.git
cd AlienInvasion
```

## Run
**Build and Run**

```bash
go build -o bin/AlienInvasion main.go
./bin/AlienInvasion
```

An example output:
```bash
12:25:44 INF > Alien 0 spawned in Foo.
12:25:44 INF > Alien 1 spawned in Bee.
Bee has been destroyed by alien 1 and alien 2
12:25:44 INF > Alien 2 tried to spawn in Bee where an alien already exists. Bee was destroyed.
12:25:44 INF > Alien 3 spawned in Qu-ux.
12:25:44 INF > Alien 4 spawned in Baz.
12:25:44 INF > Alien 0 moved to Bar.
12:25:44 INF > Alien 3 moved to Foo.
Foo has been destroyed by alien 3 and alien 4
12:25:44 INF > Alien 4 tried to move where an alien already exists. Foo was destroyed.
12:25:44 INF > All aliens are trapped in isolated cities. Exitting simulation.
```

The time-stamped lines are logs providing more detailed information about each step taken by an alien. Further detailed debug logs for each simulation execution are recorded in a *./logs* subfolder.

## Parameters
There are three available parameters for the simulation:
* Use **alienCount** (or **N**) to specify the number of aliens to spawn in the world. The default number is *5*.
* Use **mapFileName** (**m**) to specify the file containing the map information. The default value is *map.txt*
* Use **iterations** (or **i**) to specify the maximum number of iterations. The default value is *10,000*.

Detailed usage information:
```bash
Usage:
  AlienInvasion [flags]

Flags:
  -N, --alienCount int       Specify number of aliens. (default 5)
  -h, --help                 help for AlienInvasion
  -i, --iterations int       Specify number of maximum iterations. (default 10000)
  -m, --mapFileName string   Specify map file name. (default "map.txt")
```