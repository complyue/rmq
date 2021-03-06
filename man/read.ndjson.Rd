% Generated by roxygen2: do not edit by hand
% Please edit documentation in R/rmq.R
\name{read.ndjson}
\alias{read.ndjson}
\title{create an R object from the head of a stream of newline-delimited JSON bytes}
\usage{
read.ndjson(x, byteOffset)
}
\arguments{
\item{x}{A raw byte vector of one or more newline-delimited-JSON objects.}

\item{byteOffset}{specifies which object to return, namely the one starting at x[byteOffset]}
}
\value{
A list with two elements. The first is the next byteOffset that should be used to read the next object (the one after the one returned from the current call); the second is the R object represented by the frame starting at the supplied byteOffset.
}
\description{
Given a vector of raw bytes written in newline-delimited-JSON format, \code{read.ndjson} converts these into an list of 1) the next byteOffset to be used for reading the next frame (just after the next newline); and 2) the correspending reified R object obtained from converting the line at byteOffset into R.
}
\seealso{
\url{https://github.com/ugorji/go}

\url{http://ndjson.org/}

Other rmq.functions: \code{\link{from.msgpack}},
  \code{\link{r2r.call}}, \code{\link{r2r.server}},
  \code{\link{read.msgpack.frame}},
  \code{\link{read.tmframe}}, \code{\link{rmq.call}},
  \code{\link{rmq.default.addr}}, \code{\link{rmq.server}},
  \code{\link{to.msgpack}}
}

