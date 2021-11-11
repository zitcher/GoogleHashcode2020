# GoogleHashcode2020

GoogleHashcode2020 competition

# Standing
Team: UrsusArctos <br>
Points: 26,883,109 <br>
Global Rank: 370/10,724 <br>
US Rank: 22/619

# Library Problem
We recognized this problem to be a variation of the NP hard set cover problem. We used a version of the greedy appoximation algorithm for set cover to generate our solution and finished top 5%.

At the time of the competition, I wasn't familiar with metaheuristics and was more familair with approximation algorithms for NP hard problems. With my current knowledge, I know now that metaheurstics would provide a far more fruitful approach to this type of problem.

# Pizza Problem

We used dynamic programming to tackle this problem. Only 2 rows of the table are necessary for this DP problem (this saves a lot of memory).

        OptimalSlices(p[], m) = max(
            0 if OptimalSlices(p[:-1], m - p[-1]) + p[-1] > m else OptimalSlices(p[:-1], m - p[-1]) + p[-1]
            OptimalSlices(p[:-1], m)
        )
