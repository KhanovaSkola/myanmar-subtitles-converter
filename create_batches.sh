#!/bin/bash

remaining=$(wc -l < zawgyi_ytids_remaining.dat)
ytids_per_batch=35

let nbatch=remaining/ytids_per_batch+1
ibatch=6

nlines=0
for ((i=ibatch;i<nbatch+ibatch;i++));do
  let nlines+=ytids_per_batch
  head -n $nlines zawgyi_ytids_remaining.dat | tail -n $ytids_per_batch > batch$i.dat
done

