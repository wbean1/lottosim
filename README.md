# lottosim
---
Simulates various ways of selecting winning lottery tickets.  Each method should support both MegaMillions (-mm) & PowerBall (-pb) game modes.

## methods
0. simple-random: returns the first ticket created via rand()
1. double-pick: draws numbers from pool until the needed amount of numbers have been selected twice; returns this as winner
2. full-ticket-double: draws tickets (via simple-random) until the same ticket has been drawn twice; returns this as winner
3. first-ordered: draws tickets (via simple-random) until a ticket was drawn in numerical order; returns this as winner
4. highest-frequency: draws X balls and creates a ticket from the highest-frequency numbers drawn; returns this as winner