# DeepGo Dungeon

A **Procedural Dungeon Generator** built in Go using **Ebitengine** game framework. This project uses Binary Space Partitioning (BSP) to carve out randomized, retro-styled layouts.

##  The Vision

I couldn't find any straightforward implementations of BSP algorithms, nonetheless in Golang. Therefore,
the goal was to build a "game-adjacent" tool that handles the complex math of dungeon generation while maintaining
a chunky pixel aesthetic.

## How it Works: The BSP Split

The generator follows a recursive process to ensure no two dungeons are the same:

1. **Split:** The world is treated as one large rectangle. It is split into two smaller "containers" based on a random percentage (30%–70%).
2. **Ratio Control:** If a container becomes too wide or too tall, the algorithm forces a split in the opposite direction to prevent "sliver" rooms.
3. **Carve:** Once the tree reaches its maximum depth, a room is carved inside each leaf container with randomized padding.
4. **Connect:** Corridors are drawn to connect the centers of sibling nodes, ensuring every room is reachable.

## Getting Started

### Installation

1. Clone the repo:
```bash
git clone https://github.com/owenHochwald/deepgo-dungeon.git
cd deepgo-dungeon
go mod tidy # install deps
go run main.go # build and run
```
