# CSC360 Final — Go RPG Arena

A planned turn-based terminal RPG arena in Go for CSC360. The player will fight
small battles, choose attack styles, and equip upgraded weapons. The final
project will demonstrate **three design patterns**, one from each category.

## Current status

Workspace setup and project planning only. All three patterns and gameplay are
planned work.

## Planned design patterns

| Category | Pattern | Purpose in the arena | Planned demonstration |
| --- | --- | --- | --- |
| Creational | Factory Method | Create enemies through interchangeable encounter creators | Use two creators to produce different enemy types through the same contract |
| Structural | Decorator | Add equipment bonuses by wrapping a basic weapon | Show a basic weapon, one upgrade, and stacked upgrades |
| Behavioral | Strategy | Let a fighter change attack algorithms | Switch the same fighter between melee and magic against an armored target |

Together, these patterns will support a small battle: an encounter creator
produces an enemy, a decorated weapon supplies attack power, and the selected
attack strategy determines how that power interacts with armor.

Go interfaces and structs will represent the required pattern roles. The
[project plan](docs/project-plan.md) maps every essential component and outlines
a five-minute live demo.

## Workspace structure

```text
CSC360-Final/
├── cmd/
│   └── arena/
│       └── main.go          # Empty application entry point
├── internal/
│   ├── combat/             # Future fighters and attack strategies
│   ├── enemies/            # Future enemy products and encounter creators
│   ├── equipment/          # Future weapons and decorators
│   ├── game/               # Future battle loop and encounters
│   └── ui/                 # Future terminal menus and input
├── docs/
│   └── project-plan.md     # Pattern roles, milestones, and demo outline
├── .gitignore
├── go.mod
└── README.md
```

Requires Go 1.22 or newer. The empty entry point can be run with:

```sh
go run ./cmd/arena
```

It currently exits without output. Future Go tests will live beside the code
they test in files ending in `_test.go`.

## Next steps

Begin with Strategy for the initial pattern progress assignment, then add Factory
Method and Decorator for the final project. Each pattern will need a functional
example, comments explaining its roles, and tests before the final demo.
