# Understanding Go Recveivers and their constructs

I had confused myself about using Receivers and the difference between using pointer receivers and non-pointer receivers in different types of context.

This could have been easily solved by me re-reading pages 158-159 of _The Go Programming Language_ by Donovan and Kernighan, but I'd got myself into a programming pickle and had forgotten to go back to basics.

I finally did and with the help of @hoegrammer, we worked it out by making a simple coding example that tested our assumptions.

This reminded me that it's always the best thing to do: go back to basics and check one's assumptions.

In this example, we have a container for a _thing_.  The thing can be of various types, and therefore the container can be of various types of _thing container_.  Each type of thing container implements the _RunnableThingContainer_ interface, which means that the thing containers must implement the _Run()_ method.

We then have a _Box_ of these _thing containers_ and we call the _Run()_ method on each of the things in the containers in the box.

In one example the Box uses pointer receivers, meaning that the operations we do act on __the things pointed to__, which amount to the instances of the things declared in the test.  This leads us towards a more classic OO-style as implemented in languages like Java:

```go
boxOfThingContainersUsingPointerReceivers := BoxOfThingContainersUsingPointerReceivers{}
boxOfThingContainersUsingPointerReceivers.AddThingContainer(&integerContainer)
boxOfThingContainersUsingPointerReceivers.AddThingContainer(&runeContainer)
boxOfThingContainersUsingPointerReceivers.RunThings()
```

In the other example, the Box uses non-pointer receivers, meaning that the operations we do act on __copies__ of the things declared in the tests.  We therefore need to return copies of the things at the end of the methods, which leads us towards a more functional programming style as written in languages like JavaScript:

```go
BoxOfThingContainers{}.
    AddThingContainer(&integerContainer).
    AddThingContainer(&runeContainer).
    RunThings()
```