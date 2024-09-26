# Introduction
There is a long-existing [cimgui-go](https://github.com/AllenDang/cimgui-go) issue [#224](https://github.com/AllenDang/cimgui-go/issues/224).
It means to improve our code generator to support C callbacks.

# The Prolbem

Current way to do C->Go->C callbacks is to export go global function to C and use it as a C callback.
An example code is present in this repo [here](problem/main.go).

There is no way to pass a local go function or method as a C callback.

# Considered Paths of solution

If anyone has any idea, all cimgui-go community (and I hope not only) will be grateful.
You're welcome to open a discussion/issue with your idea!

## Path 1: Use some type of global registry

Add a global variable for the callback, store it and add global function that calls it.

Disadvantages:
- absolutely disqualified for our project - can use only 1 callback at a time.

## Path 2: Understand how Go functino pointers works and force-call them from C

This absolutely crazy idea. I have no idea how to do it and I'm not sure it's possible at all.

