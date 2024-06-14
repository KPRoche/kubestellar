# KubeStellar Quickstart Setup

This Quick Start is based on Scenario 1 of our [examples](examples.md).
In a nutshell, it will help you:

    1. Prepare your system (prerequisites)
    2. Create the Kubestellar core components on a cluster
    3. Commission a workload to a WEC

## Before You Begin


{%
    include-markdown "pre-reqs.md"
    end="## For Building"
    heading-offset=2
%}


## Create the KubeStellar Core components

First set up the main core and establish its initial state using our helm chart:

  {%
    include-markdown "core-chart.md"
    heading-offset=2
  %}

## Define, bind and commission a workload on a WEC

  {%
    include-markdown "example-wecs.md"
    heading-offset=2
  %}

