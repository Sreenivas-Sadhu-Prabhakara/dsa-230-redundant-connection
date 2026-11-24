# Redundant Connection

> NeetCode 150 · Graphs — feed edges to union-find in input order; the first union whose endpoints already share a root is the extra edge. Go, stdlib only, tested.

## Problem

A tree got one extra edge (n nodes, n edges). Return the edge to delete so a tree remains — specifically the one appearing *last* in the input among valid choices.

## Approach

This is union-find's signature use case: process edges in given order, merging endpoint sets; an edge whose endpoints already share a root would close a cycle — and because every earlier edge was tree-compatible, this is exactly the "last valid answer" the problem wants, straight out of the processing order with no extra logic. Path-halving `find` plus union-by-rank give effectively O(1) per edge. DFS could find the cycle too, but would then need tie-breaking on input order — the data structure encodes the requirement for free, which is the elegance being taught. (Kruskal, dsa-169, is this same skip-the-cycle-closer loop with weights.)

**Complexity:** O(n α(n)) ≈ O(n) time, O(n) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
