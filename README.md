# Self Organizing List in Go

[![Run Status](https://api.shippable.com/projects/59dba4fb183eb3070005b779/badge?branch=master)](https://app.shippable.com/github/Riturajcse/Self-Organizing-List)
[![Go Report Card](https://goreportcard.com/badge/github.com/Riturajcse/Self-Organizing-List)](https://goreportcard.com/report/github.com/Riturajcse/Self-Organizing-List)
[![github license](https://img.shields.io/github/license/Riturajcse/Self-Organizing-List.svg)](https://github.com/Riturajcse/Self-Organizing-List)

This is a library implementing Self Organizing List for the Go programming language (http://golang.org/).

A self-organizing list is a list that reorders its elements based on some self-organizing heuristic to improve average access time.The aim of a self-organizing list is to improve efficiency of linear search by moving more frequently accessed items towards the head of the list. A self-organizing list achieves near constant time for element access in the best case. A self-organizing list uses a reorganizing algorithm to adapt to various query distributions at runtime.

## Techniques for Rearranging Nodes
While ordering the elements in the list, the access probabilities of the elements are not generally known in advance. This has led to the development of various heuristics to approximate optimal behavior. The basic heuristics used to reorder the elements in the list are:

### 1. Move to Front Method (MTF):

This technique moves the element which is accessed to the head of the list.
```
    At the t-th item selection:
         if item i is selected:
                 move item i to head of the list
           
```

### 2. Count Method:

In this technique, the number of times each node was searched for is counted i.e. every node keeps a separate counter variable which is incremented every time it is called. 
```
    init: count(i) = 0 for each item i
    At t-th item selection:
       if item i is searched:
           count(i) = count(i) + 1
           rearrange items based on count
           
```

### 3. Transpose Method:

This technique involves swapping an accessed node with its predecessor. 
```
    At the t-th item selection:
         if item i is selected:
             if i is not the head of list:
                     swap item i with item (i - 1)
           
```
