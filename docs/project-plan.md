# Project plan

## Goal and scope

Build a small terminal RPG arena in Go demonstrating one creational, one
structural, and one behavioral design pattern. Keep the game small enough to
explain all three patterns and show their behavior in a five-minute live demo.

Start with one arena, two enemy types, two attack styles, and a basic weapon
with two possible upgrades. Use standard-library Go and text output. Multiplayer,
graphics, databases, and a large world are outside the initial scope.

Only workspace setup is complete. All pattern implementations, gameplay, and
tests are future work.

## Creational: Factory Method

**Purpose:** let encounter creators decide which enemy to construct while game
code works with shared creator and enemy contracts.

| Required element | Planned role |
| --- | --- |
| Creator interface | Encounter creator contract defining an enemy factory method |
| Concrete creator | Goblin encounter creator and skeleton encounter creator, each implementing the factory method |
| Product interface | Shared enemy behavior and statistics contract |
| Concrete product | Goblin and skeleton implementations of the enemy contract |

Plan a shared encounter setup workflow that receives a creator, calls its factory
method, and uses the returned enemy through the product interface. Concrete
creators choose their products. In Go, use interface-based composition for this
collaboration rather than relying on class inheritance. Keep product creation in
the concrete creators rather than a single switch over enemy names.

**Live example:** use the same encounter setup with two creators and display the
different enemies produced. Explain that adding another enemy/creator should not
require changing the encounter setup workflow.

## Structural: Decorator

**Purpose:** add and combine weapon bonuses without defining a separate weapon
type for every combination of upgrades.

| Required element | Planned role |
| --- | --- |
| Component interface | Weapon contract exposing a description and attack power |
| Concrete component | Basic sword implementing that contract |
| Decorator interface | Weapon-compatible decorator contract; a shared wrapper holds another weapon and delegates to it |
| Concrete decorator | Sharpened and enchanted upgrades that wrap a weapon and extend its description and power |

Each decorated weapon must remain usable through the same weapon interface.
Plan for the shared wrapper to retain the wrapped component and for concrete
upgrades to extend delegated behavior. Upgrades should be composable at runtime.

**Live example:** display the same sword before upgrades, after one upgrade, and
after stacking two upgrades. Explain how wrapping avoids a separate class or
struct for every upgrade combination.

## Behavioral: Strategy

**Purpose:** change how a fighter calculates attack damage without placing all
attack algorithms inside the fighter.

| Required element | Planned role |
| --- | --- |
| Strategy interface | Shared attack calculation contract |
| Concrete strategy | Melee and magic implementations with different armor interactions |
| Context | Fighter that stores a selected strategy and delegates attack calculations to it |

Game code will choose the strategy. Weapon power can come from the equipment
component; the strategy handles the attack algorithm, while decorators handle
weapon upgrades.

**Live example:** attack the same armored enemy with one fighter, switch the
fighter's strategy, and show the changed result.

## Milestones

1. **Workspace setup (complete):** Go module, empty entry point, organized
   folders, README, planning notes, and connected GitHub repository.
2. **Strategy progress assignment (planned):** design the attack contract,
   implement two styles and a context, demonstrate switching, explain the roles,
   and commit the work.
3. **Factory Method (planned):** define enemy and creator contracts, implement
   two concrete creators/products, and demonstrate the shared encounter workflow.
4. **Decorator (planned):** define the weapon component and decorator contracts,
   implement a basic weapon and two wrappers, and demonstrate stacking upgrades.
5. **Small playable battle (planned):** connect enemy creation, equipment, and
   attack selection; add health, turn order, input handling, and win/loss outcomes.
6. **Demo preparation (planned):** test each pattern and battle outcomes, add
   concise comments identifying the essential roles, and rehearse the walkthrough.

## Organization

- `cmd/arena/`: application entry point and wiring.
- `internal/combat/`: future fighters and attack strategies.
- `internal/enemies/`: future enemy products and encounter creators.
- `internal/equipment/`: future weapon component and decorators.
- `internal/game/`: future turn loop and encounter setup workflow.
- `internal/ui/`: future terminal menus and input handling.
- `docs/`: planning notes.
- Future tests will live beside their Go packages in `_test.go` files.

## Five-minute demo outline

- **0:00–0:30:** introduce the arena and name each pattern and category.
- **0:30–1:40:** show Factory Method roles and create two enemy types.
- **1:40–2:50:** show Decorator roles and stack weapon upgrades.
- **2:50–4:00:** show Strategy roles and switch attack styles on one fighter.
- **4:00–5:00:** demonstrate a short integrated battle and explain the benefits
  and how the patterns work together.

Use predictable scenarios so differences are visible without waiting for random
outcomes. Keep the files for each pattern ready for the code walkthrough.

## Decisions to make while implementing

- What minimum information belongs in each interface?
- How should enemy types and attack styles differ?
- How should equipment upgrades combine and stay balanced?
- What is the smallest battle that demonstrates all three patterns clearly?
