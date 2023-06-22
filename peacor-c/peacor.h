
/*
 computes Pearson correlation matrix for a set of observations for multiple variates
 accepts tab or space separated files
 Note: spaces are NOT allowed within fields
*/
#ifndef PEACOR_H
#define PEACOR_H
#include <stdlib.h>
#include <stdio.h>
#include <math.h> // sqrt
#include <getopt.h>
#include <string.h> // strcpy
#include <gsl/gsl_statistics.h>
#include "util.h"
/*
#include <stdarg.h> // variable number of arguments
#include <stddef.h> // NULL (null pointer constant); size_t;
*/

const int STRMX = 64; // max length of variable's symbols
const int DBLMX = 16; // max length of variable's values
const float VERSION = 0.2;
#endif
