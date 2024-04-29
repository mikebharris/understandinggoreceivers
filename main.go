package understandingpointerreceivers

type RunnableThingContainer interface {
	Run()
}

type IntegerContainer struct {
	thing int
}

func (w *IntegerContainer) Run() {
	w.thing = 2
}

type RuneContainer struct {
	thing rune
}

func (r *RuneContainer) Run() {
	r.thing = 'z'
}

type BoxOfThingContainers struct {
	thingContainers []RunnableThingContainer
}

func (w BoxOfThingContainers) AddThingContainer(thingContainer RunnableThingContainer) BoxOfThingContainers {
	w.thingContainers = append(w.thingContainers, thingContainer)
	return w
}

func (w BoxOfThingContainers) RunThings() BoxOfThingContainers {
	for _, thingContainer := range w.thingContainers {
		thingContainer.Run()
	}
	return w
}

type BoxOfThingContainersUsingPointerReceivers struct {
	thingContainers []RunnableThingContainer
}

func (w *BoxOfThingContainersUsingPointerReceivers) AddThingContainer(thingContainer RunnableThingContainer) {
	w.thingContainers = append(w.thingContainers, thingContainer)
}

func (w *BoxOfThingContainersUsingPointerReceivers) RunThings() {
	for _, thingContainer := range w.thingContainers {
		thingContainer.Run()
	}
}
