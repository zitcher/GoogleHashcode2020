# GoogleHashcode2020

GoogleHashcode2020 competition

# Pizza Problem

    Recursive call:
        OptimalSlices(p[], m) = max(
            0 if OptimalSlices(p[:-1], m - p[-1]) + p[-1] > m else OptimalSlices(p[:-1], m - p[-1]) + p[-1]
            OptimalSlices(p[:-1], m)
        )
