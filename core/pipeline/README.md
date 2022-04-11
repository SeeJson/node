[![GoDoc](https://godoc.org/github.com/SeeJson/node/core/pipeline?status.svg)](https://godoc.org/github.com/SeeJson/node/core/pipeline)
[![Build Status](https://travis-ci.org/SeeJson/node/core/pipeline.svg?branch=master)](https://travis-ci.org/SeeJson/node/core/pipeline)
[![cover.run](https://cover.run/go/github.com/SeeJson/node/core/pipeline.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fhyfather%2Fpipeline)
[![Go Report Card](https://goreportcard.com/badge/github.com/SeeJson/node/core/pipeline)](https://goreportcard.com/report/github.com/SeeJson/node/core/pipeline)

# pipeline

This package provides a simplistic implementation of Go pipelines
as outlined in [Go Concurrency Patterns: Pipelines and cancellation.](https://blog.golang.org/pipelines)

# Docs
GoDoc available [here.](https://godoc.org/github.com/SeeJson/node/core/pipeline)

# Example Usage

```
import "github.com/SeeJson/node/core/pipeline"

p := pipeline.New()
p.AddStageWithFanOut(myStage, 10)
p.AddStageWithFanOut(anotherStage, 100)
doneChan := p.Run(inChan)

<- doneChan
```

More comprehensive examples can be found [here.](./examples)
