Refactoring_071525.md
I want to stop this refactoring to a non-ps2ns embodiment and pursue something different. I request that your responses be concise and that you make no suggestions unless I specifically request them.

The following will be discussed with respect to the current version of rudd_Large_070925.go located at ~/Golang/Large. Do not alter the existing '070925' file. Instead, create a new file using so much of '070925' as you find useful.

The platform used in the '070925' file has a user interface that allows a user to find a test pattern sequence for a 'fault_A'. One representation of the found test pattern sequence is the sequence of S/I's the user enters throughout his search.

We are going to separate the 'search phase' (ps2ns, ns2fp, seek) from the 'simulation phase' (setUP, one_set_BOOL). They are pretty much separated functionally, and we will make a complete separation by creating a replacement for the existing foo().

The user interface will be that user enters a 'fault_A' and leaves 'fault_C' blank to begin a 'search phase. During the 'search phase' the user reads the displayed seek() created list and select's a displayed S/I as at present. In this refactor the sequence of selected S/I's is accumulated, in order, and will define a test pattern sequence for 'fault_A'. So gathering, in order of entry, the selected S/I's is necessary. The accumulated S/I's will be reset at a fresh start and following any '999' reset. No slimulation using setUP, one_set_BOOL will be done (that's a change to foo).

In addition to '999' and at the same point the user will be offered a second option, an 'xxx' break which will not reset the accumulation of S/I's. 'xxx' is some three digit code different from '999'. 'xxx' will begin a 'simulation phase' during which the accumulated S/I's will be applied to a 'fault_C'. The initial 'fault_C' will be blank and the outputs and next-state of the simulation will be saved as 'fault_C' is blank, and 'fault_C' equals 'fault_A'.

Now, apply the accumulated S/I's from the 'search phase' to (1) a 'fault_C' not blank and not equal to 'fault_A', or (2) one-at-a-time to each 'fault_C' in a user provided list of 'fault_C's, or (3) to all possible 'fault_C's in the circuit.

Having found a sequence, the user can manually copy the sequence to paper and then use that sequence to test it's effect on a simulated 'fault_C' machine. He does that by selecting the 'fault_A' and a 'fault_C' at the start, or following a '999' reset.

I would like an easily user setable switch to switch between a 'search' phase, and a second phase that will automate some of what will take place while using the '070925' user interface. The switch I referred to will be selectable at a '999' point. In other words, when reaching a point at which '999' is a user option, the user will be presented with an 'xxx' option ('xxx' being some arbitrary three digit number other than '999') that will select this alternate mode of operation. If the user chooses to continue or restart his search for a test sequence the operation will remain as it is at present (continue or '999' resent). However if the user selects the 'xxx' option at that point, the new mode of operation will begin and will use the S/I sequence accumulated from a fresh start of from a '999' reset. So there will be a need to accumulate successive S/I's for use in this 'xxx' alternate phase.

During a 'search phase' (1) accumulate the sequence of S/I's since start or since '999' reset, and (2) reset the accumulation of S/I's at a fresh restart or '999' reset.

When the user selects the 'xxx' switch, he is signallng that he has found a test sequence for 'fault_A'. and is ready to begin a 'simulation phase'. Simulation has been ongoing during the 'search phase' as it is presently in the '070925' file. But during the 'xxx' selected 'simulation phase' we will be using 'fault_A' and the accumulated S/I's