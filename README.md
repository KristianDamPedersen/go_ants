# GO ants
A small terminal based Ant Colony Optimization (ACO) program. This repo consists of two elements, the rendering (built on top of Termui) and the actual mathematics.
Big shoutout to the awesome work behind Termui, please check out: https://github.com/gizak/termui.

## 🚀 Getting started
To get started with this project, start off by cloning it to your local machine and make sure you have Go installed (https://go.dev/). After that, cd into the project directory and run:
```
go get && go install && go build

```
Which will pull the dependencies, install them and compile the program. Now simply run: 
```
./go_ants

```
To execute the program. Currently this will only display a rendering of connected, randomly generated dots in the terminal.

## 📌 Current features
* Can generate list of random points
* Can draw a list of points and lines between them in the terminal

## 🚧 Known issues
1. Can't render vertical lines between points

## Roadmap
* ✅ Basic Terminal-based rendering
* Single-ant traversal
