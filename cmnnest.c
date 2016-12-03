#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <time.h>
#include <float.h>
#include <math.h>
#include "figure.h"
#include "nest_structs.h"
#include "crosscheck.h"
#include "cmnfuncs.h"
#include "nestdefs.h"

int mutate(struct Individ *src, struct Individ *mutant, int setsize);
int gensequal(struct Individ *indiv1, struct Individ *indiv2);
int gensequal2(struct Individ *indiv1, struct Individ *indiv2, struct Figure *figset);
int crossover(struct Individ *par1, struct Individ *par2, struct Individ *child, int setsize);

int crossover(struct Individ *par1, struct Individ *par2, struct Individ *child, int setsize)
{
	int i, j;
	int g1, g2;

	if (par1->gensize < 3)
		return SMALL_INDIVID;

	if (par1->gensize != par2->gensize)
		return DIFFERENT_SIZE;

	srand(time(NULL));

	g1 = rand() % (par1->gensize);
	g2 = rand() % (par2->gensize);

	while (g1 == g2) 
		g2 = rand() % (par1->gensize);

	child->par1 = par1->genom;
	child->par2 = par2->genom;
	child->gensize1 = par1->gensize;
	child->gensize2 = par2->gensize;
	child->g1 = g1;
	child->g2 = g2;
	
	child->genom = (int*)xmalloc(sizeof(int) * setsize);
	child->gensize = par1->gensize;
	
	child->genom[g1] = par1->genom[g1];
	child->genom[g2] = par1->genom[g2];

	for (i = 0, j = 0; i < par2->gensize && j < par2->gensize; i++, j++) {
		if (j == g1 || j == g2) {
			i--;
			continue;
		}

		if (par2->genom[i] == child->genom[g1] || par2->genom[i] == child->genom[g2]) {
			j--;
			continue;
		}

		child->genom[j] = par2->genom[i];
	}
	
	return SUCCESSFUL;	
}

int gensequal2(struct Individ *indiv1, struct Individ *indiv2, struct Figure *figset) 
{
	int i;

	if (indiv1->gensize != indiv2->gensize)
		return 0;

	for (i = 0; i < indiv1->gensize; i++) {
		int a, b;

		a = indiv1->genom[i];
		b = indiv2->genom[i];

		if (figset[a].id != figset[b].id)
			return INDIVIDS_UNEQUAL;
	}

	return INDIVIDS_EQUAL;
}

int gensequal(struct Individ *indiv1, struct Individ *indiv2) 
{
	int i;

	if (indiv1->gensize != indiv2->gensize)
		return 0;

	for (i = 0; i < indiv1->gensize; i++)
		if (indiv1->genom[i] != indiv2->genom[i])
			return INDIVIDS_UNEQUAL;

	return INDIVIDS_EQUAL;
}

int mutate(struct Individ *src, struct Individ *mutant, int setsize)
{
 	int n1, n2, tmp, i;

	if (src->gensize == 1)
		return SMALL_INDIVID;
	
	mutant->gensize = src->gensize;
	mutant->genom = (int*)xmalloc(sizeof(int*) * setsize);
    
	srand(time(NULL));

	n1 = rand() % (src->gensize - 1);
	n2 = rand() % (src->gensize);

	while (n1 == n2) 
		n2 = rand() % (src->gensize);
 
	for (i = 0; i < src->gensize; i++)
		mutant->genom[i] = src->genom[i];
		
	tmp = mutant->genom[n1];
	mutant->genom[n1] = mutant->genom[n2];
	mutant->genom[n2] = tmp;

	return SUCCESSFUL;
}